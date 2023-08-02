package user_management

import (
	"SkripsiBebek/db"
	"SkripsiBebek/struct_all/user_management"
	"SkripsiBebek/tools"
	"net/http"
	"strconv"
)

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
