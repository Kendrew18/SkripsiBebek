package user_management

import (
	"SkripsiBebek/db"
	"SkripsiBebek/struct_all/user_management"
	"SkripsiBebek/tools"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//Read_CSV
func Read_CSV(writer http.ResponseWriter, request *http.Request, BuidingID string) (tools.Response, error) {
	var res tools.Response

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("excel")
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name : ", handler.Filename)
	fmt.Println("File Size : ", handler.Size)
	fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

	var tempFile *os.File
	path := ""
	if strings.Contains(handler.Filename, "csv") {
		path = "excel/Read" + ".csv"
		tempFile, err = ioutil.TempFile("excel/", "Read"+"*.csv")
	}

	if err != nil {
		return res, err
	}

	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return res, err2
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return res, err
	}

	fmt.Println("Success!!")
	fmt.Println(tempFile.Name())
	tempFile.Close()

	err = os.Rename(tempFile.Name(), path)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fmt.Println(tempFile.Name())

	fd, error := os.Open("./excel/Read.csv")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Successfully opened the CSV file")

	// read CSV file
	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(records[0])

	fd.Close()

	_ = os.Remove("./excel/Read.csv")

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT co FROM resident ORDER BY co DESC LIMIT 1"

	err = con.QueryRow(sqlStatement).Scan(&nm)

	co := 0

	for i := 1; i < len(records); i++ {

		ID := ""

		sqlStatement := "SELECT ResidentID FROM resident WHERE BuildingID=? && email=? "

		_ = con.QueryRow(sqlStatement, BuidingID, records[i][3]).Scan(&ID)

		if ID == "" {

			co++

			nm = nm + int64(1)

			temp := strconv.FormatInt(nm, 10)

			ResidentID := "R-" + temp

			sqlStatement = "INSERT INTO resident(co, ResidentID, name, surname, room_no, email, alphanumeric, password, BuildingID) values(?,?,?,?,?,?,?,?,?)"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(nm, ResidentID, records[i][0], records[i][1], records[i][2], records[i][3], records[i][4], records[i][4], BuidingID)

			stmt.Close()
		}
	}

	if co == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Data Already Include"
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
	}

	return res, nil
}

//See_All_Resident
func See_All_Resident(BuildingID string) (tools.Response, error) {
	var res tools.Response
	var Resident user_management.Read_All_Resident
	var arr_Resident []user_management.Read_All_Resident

	con := db.CreateCon()

	sqlStatement := "SELECT ResidentID,name,surname,room_no,email FROM resident WHERE BuildingID=? ORDER BY co ASC"

	rows, err := con.Query(sqlStatement, BuildingID)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&Resident.ResidentID, &Resident.Name,
			&Resident.Surname, &Resident.Room_No, &Resident.Email)
		if err != nil {
			return res, err
		}
		arr_Resident = append(arr_Resident, Resident)
	}

	if arr_Resident == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Resident
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Resident
	}

	return res, nil
}

//Delete_Resident
func Delete_Resident(ResidentID string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "DELETE FROM resident WHERE ResidentID=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(ResidentID)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowsAffected,
	}

	return res, nil
}
