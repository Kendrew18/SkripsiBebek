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

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT count(BuildingID) FROM building ORDER BY co ASC "

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

	sqlStatement = "SELECT count(AdminID) FROM admin ORDER BY co ASC "

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

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}
