package user_management

import (
	"SkripsiBebek/db"
	"SkripsiBebek/tools"
	"net/http"
	"strconv"
)

func Create_Postman(Name string, Email string, Password string) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT count(PostmanID) FROM postman ORDER BY co ASC "

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

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}
