package user_management

import (
	"SkripsiBebek/db"
	"SkripsiBebek/struct_all/user_management"
	"SkripsiBebek/tools"
	"net/http"
	"strconv"
)

//Login
func Login(Email string, Password string, Status int) (tools.Response, error) {

	var res tools.Response
	var us user_management.Login

	con := db.CreateCon()

	if Status == 3 {

		sqlStatement := "SELECT PostmanID FROM postman where Email=? && Password=? "

		err := con.QueryRow(sqlStatement, Email, Password).Scan(&us.ID)

		us.Status = Status
		us.BuildingID = ""

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			us.ID = ""
			res.Data = us
			return res, nil
		}
	} else if Status == 2 {

		sqlStatement := "SELECT AdminID,BuildingID FROM admin where Email=? && Password=? "

		err := con.QueryRow(sqlStatement, Email, Password).Scan(&us.ID, &us.BuildingID)

		us.Status = Status

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			us.ID = ""
			res.Data = us
			return res, nil
		}

	} else if Status == 1 {

		sqlStatement := "SELECT ResidentID,BuildingID FROM resident where email=? && password=? "

		err := con.QueryRow(sqlStatement, Email, Password).Scan(&us.ID, &us.BuildingID)

		us.Status = Status

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			us.ID = ""
			res.Data = us
			return res, nil
		}
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = us

	return res, nil
}

//Create Admin
func Create_Admin_And_Building(Email string, Password string, Name string,
	BuildingName string, Address string, Biography string) (tools.Response, error) {
	var res tools.Response
	var data user_management.Login

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT co FROM building ORDER BY co DESC LIMIT 1"

	err := con.QueryRow(sqlStatement).Scan(&nm)

	nm = nm + 1

	temp := strconv.FormatInt(nm, 10)

	BuildingID := "B-" + temp

	sqlStatement = "INSERT INTO building(co,BuildingID, BuildingName, Address, Biography) values(?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm, BuildingID, BuildingName, Address, Biography)

	//Admin
	nm = int64(0)

	sqlStatement = "SELECT co FROM admin ORDER BY co DESC LIMIT 1"

	err = con.QueryRow(sqlStatement).Scan(&nm)

	nm = nm + 1

	temp = strconv.FormatInt(nm, 10)

	AdminID := "AD-" + temp

	sqlStatement = "INSERT INTO admin(co,AdminID, Email, Password, Name, BuildingID) values(?,?,?,?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm, AdminID, Email, Password, Name, BuildingID)

	stmt.Close()

	data.ID = AdminID
	data.Status = 2
	data.BuildingID = BuildingID

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = data

	return res, nil
}

//Update_Profile_Admin_Building
func Update_Profile_Admin_Building(Admin_id string, Building_Id string, Email string, Password string, Name string,
	BuildingName string, Address string, Biography string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := ""

	sqlStatement := "SELECT AdminID FROM admin WHERE admin.Email=? && admin.AdminID!=?"

	err := con.QueryRow(sqlStatement, Email, Admin_id).Scan(&nm)

	if err != nil {
		return res, err
	}

	if nm == "" {

		sqlStatement = "UPDATE admin SET Name=?,Email=?,Password=? WHERE AdminID=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(Name, Email, Password, Admin_id)

		if err != nil {
			return res, err
		}

		rowschanged, err := result.RowsAffected()

		if err != nil {
			return res, err
		}

		//Building
		sqlStatement = "UPDATE building SET BuildingName=?,Address=?,Biography=? WHERE BuildingID=?"

		stmt, err = con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err = stmt.Exec(BuildingName, Address, Biography, Building_Id)

		if err != nil {
			return res, err
		}

		stmt.Close()

		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": rowschanged,
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "e-mail already exists"
	}

	return res, nil
}

//See_Profile
func See_Profile(id string, status int) (tools.Response, error) {
	var res tools.Response
	con := db.CreateCon()
	if status == 1 {
		//res
		var resi user_management.Resident_Profile

		sqlStatement := "SELECT ResidentID, name, surname,email,password FROM resident WHERE ResidentID=?"

		err := con.QueryRow(sqlStatement, id).Scan(&resi.Resident_id,
			&resi.Name, &resi.Surname, &resi.Email, &resi.Password)

		if err != nil {
			return res, err
		}

		if resi.Resident_id == "" {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = resi
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = resi
		}

	} else if status == 2 {
		//admin
		var Admin user_management.Admin_Profile

		sqlStatement := "SELECT AdminID,b.BuildingID, admin.Name,admin.Email,admin.Password,b.BuildingName,b.Address,b.Biography FROM admin Join building join building b on admin.BuildingID = b.BuildingID WHERE AdminID=?"

		err := con.QueryRow(sqlStatement, id).Scan(&Admin.Admin_id, &Admin.Building_id,
			&Admin.Name, &Admin.Email, &Admin.Password, &Admin.Building_Name,
			&Admin.Address, &Admin.Biography)

		if err != nil {
			return res, err
		}

		if Admin.Admin_id == "" {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = Admin
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = Admin
		}

	} else if status == 3 {
		//postman
		var pos user_management.Postman_Profile

		sqlStatement := "SELECT PostmanID,Name,Email,Password FROM postman WHERE PostmanID=?"

		err := con.QueryRow(sqlStatement, id).Scan(&pos.Postman_id,
			&pos.Name, &pos.Email, &pos.Password)

		if err != nil {
			return res, err
		}

		if pos.Postman_id == "" {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = pos
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = pos
		}
	}

	return res, nil
}
