package user_management

import (
	"SkripsiBebek/db"
	"SkripsiBebek/struct_all/user_management"
	"SkripsiBebek/tools"
	"net/http"
	"strconv"
)

//Sign_Up_Postman
func Create_Postman(Name string, Email string, Password string) (tools.Response, error) {

	var res tools.Response
	var data user_management.Login

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT co FROM postman ORDER BY co DESC LIMIT 1"

	err := con.QueryRow(sqlStatement).Scan(&nm)

	nm = nm + 1

	temp := strconv.FormatInt(nm, 10)

	PostmanID := "P-" + temp

	sqlStatement = "INSERT INTO postman(co,PostmanID, Email, Password, Name) values(?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm, PostmanID, Email, Password, Name)

	stmt.Close()

	data.ID = PostmanID
	data.Status = 1
	data.BuildingID = ""

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = data

	return res, nil
}

//Update_Profile_Postman
func Update_Profile_Postman(postman_id string, name string, email string, password string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := ""

	sqlStatement := "SELECT PostmanID FROM postman WHERE postman.email=? && postman.PostmanID!=?"

	_ = con.QueryRow(sqlStatement, email, postman_id).Scan(&nm)

	if nm == "" {

		sqlStatement = "UPDATE postman SET name=?,email=?,password=? WHERE PostmanID=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(name, email, password, postman_id)

		if err != nil {
			return res, err
		}

		rowschanged, err := result.RowsAffected()

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
