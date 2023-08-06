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

//See_Profile_Admin
func See_Profile_Admin(Admin_id string) (tools.Response, error) {
	var res tools.Response
	var Admin user_management.Admin_Profile

	con := db.CreateCon()

	sqlStatement := "SELECT AdminID,b.BuildingID, admin.Name,admin.Email,admin.Password,b.BuildingName,b.Address,b.Biography FROM admin Join building join building b on admin.BuildingID = b.BuildingID WHERE AdminID=?"

	err := con.QueryRow(sqlStatement, Admin_id).Scan(&Admin.Admin_id, &Admin.Building_id,
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

	return res, nil
}
