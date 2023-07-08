package user_management

import (
	"SkripsiBebek/db"
	"SkripsiBebek/struct_all/user_management"
	"SkripsiBebek/tools"
	"net/http"
)

func Login(Email string, Password string, Status int) (tools.Response, error) {

	var res tools.Response
	var us user_management.Login

	con := db.CreateCon()

	if Status == 3 {

		sqlStatement := "SELECT PostmanID FROM postman where Email=? && Password=? "

		err := con.QueryRow(sqlStatement, Email, Password).Scan(&us.ID)

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
