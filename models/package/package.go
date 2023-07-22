package _package

import (
	"SkripsiBebek/db"
	"SkripsiBebek/struct_all/st_package"
	"SkripsiBebek/tools"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//Input-st_package
func Input_Package(PostmanID string, NoResi string, Name string,
	Street_Name string, Building_Name string, Room_Number string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT count(PackageID) FROM package ORDER BY co ASC "

	err := con.QueryRow(sqlStatement).Scan(&nm)

	nm = nm + 1

	temp := strconv.FormatInt(nm, 10)

	PackageID := "PA-" + temp

	sqlStatement = "INSERT INTO package(co,PackageID, NoResi, Name, Street_Name,Building_Name,Room_Number,PostmanID) values(?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm, PackageID, NoResi, Name, Street_Name,
		Building_Name, Room_Number, PostmanID)

	nm = int64(0)

	sqlStatement = "SELECT count(IDDetailStatus) FROM detail_status ORDER BY co ASC "

	err = con.QueryRow(sqlStatement).Scan(&nm)

	nm = nm + 1

	temp = strconv.FormatInt(nm, 10)

	DST := "DS-" + temp

	var time1 = time.Now()
	date_sql := time1.Format("2006-01-02")

	sqlStatement = "INSERT INTO detail_status(co,iddetailstatus, idpacakage, idstatus, date) values(?,?,?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm, DST, PackageID, "STAT-1", date_sql)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read_package
func Read_Package(ResidentID string, AdminID string, PostmanID string) (tools.Response, error) {
	var res tools.Response
	var Package st_package.Read_Package
	var arr_Package []st_package.Read_Package

	con := db.CreateCon()

	if ResidentID != "" {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,StatusName FROM package join detail_status ds on package.IDDetail = ds.IDDetailStatus  join status ON ds.IDStatus=status.StatusID WHERE ResidentID=? ORDER BY co ASC"

		rows, err := con.Query(sqlStatement, ResidentID)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&Package.Id_Package, &Package.No_Resi,
				&Package.Street_Name, &Package.Status)
			if err != nil {
				return res, err
			}
			arr_Package = append(arr_Package, Package)
		}
	} else if AdminID != "" {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,IDDetail FROM package join detail_status ds on package.IDDetail = ds.IDDetailStatus  join status ON ds.IDStatus=status.StatusID WHERE AdminID=? && ds.IDStatus=? ORDER BY co ASC"

		rows, err := con.Query(sqlStatement, AdminID, "STAT-2")

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&Package.Id_Package, &Package.No_Resi,
				&Package.Street_Name, &Package.Status)
			if err != nil {
				return res, err
			}
			arr_Package = append(arr_Package, Package)
		}
	} else if PostmanID != "" {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,IDDetail FROM package join detail_status ds on package.IDDetail = ds.IDDetailStatus  join status ON ds.IDStatus=status.StatusID WHERE PostmanID=? && ds.IDStatus=? ORDER BY co ASC"

		rows, err := con.Query(sqlStatement, PostmanID, "STAT-1")

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&Package.Id_Package, &Package.No_Resi,
				&Package.Street_Name, &Package.Status)
			if err != nil {
				return res, err
			}
			arr_Package = append(arr_Package, Package)
		}
	}

	if arr_Package == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Package
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Package
	}

	return res, nil
}

//Read-Detail-Package (Resident)
func Read_Detail_Package_Resident(PackageID string) (tools.Response, error) {
	var res tools.Response
	var Package st_package.Detail_Resident
	var arr_Package []st_package.Detail_Resident

	con := db.CreateCon()

	sqlStatement := "SELECT date,StatusName FROM detail_status join status s ON s.StatusID=detail_status.IDStatus where IDPacakage=?"

	rows, err := con.Query(sqlStatement, PackageID)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&Package.Date, &Package.Status_Name)
		if err != nil {
			return res, err
		}
		arr_Package = append(arr_Package, Package)
	}

	if arr_Package == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Package
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Package
	}

	return res, nil
}

//Read-Detail-Package (Postman)
func Read_Detail_Package_Postman(PackageID string) (tools.Response, error) {
	var res tools.Response
	var Package st_package.Detail_Postman
	var arr_Package []st_package.Detail_Postman

	con := db.CreateCon()

	sqlStatement := "SELECT PackageID, NoResi, Name, Street_Name, Building_Name, Room_Number FROM package WHERE ResidentID=? ORDER BY co ASC"

	rows, err := con.Query(sqlStatement, PackageID)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&Package.PackageID, &Package.Noresi,
			&Package.Name, &Package.Street_Name,
			&Package.Building_Name, &Package.Room_Number)
		if err != nil {
			return res, err
		}
		arr_Package = append(arr_Package, Package)
	}

	if arr_Package == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Package
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Package
	}

	return res, nil
}

//Update-Status-Package (Return Postman)
func Update_Status_Package(packageID string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT count(IDDetailStatus) FROM detail_status ORDER BY co ASC "

	err := con.QueryRow(sqlStatement).Scan(&nm)

	nm = nm + 1

	temp := strconv.FormatInt(nm, 10)

	DST := "DS-" + temp

	var time1 = time.Now()
	date_sql := time1.Format("2006-01-02")

	sqlStatement = "INSERT INTO detail_status(co,iddetailstatus, idpacakage, idstatus, date) values(?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm, DST, packageID, "STAT-4", date_sql)

	sqlStatement = "UPDATE package SET IDDetail=? WHERE PackageID=?"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(DST, packageID)

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

	return res, nil
}

//Update-Status-Package (Admin)
func Update_Status_Package_Admin(AdminID string, NoResi string, Name string,
	Street_Name string, Building_Name string, Room_Number string) (tools.Response, error) {
	var res tools.Response
	con := db.CreateCon()

	temp := ""

	sqlStatement := "SELECT ResidentID FROM resident join building b on b.BuildingID = resident.BuildingID WHERE room_no=? && BuildingName=? && Address=? && CONCAT(name,' ',surname) LIKE ?"

	TN := "'%" + Name + "%'"

	fmt.Println(TN)
	_ = con.QueryRow(sqlStatement, Room_Number, Building_Name, Street_Name, TN).Scan(&temp)

	if temp != "" {

		//Create detail status baru
		packageID := ""

		sqlStatement := "SELECT PackageID FROM package WHERE NoResi=? && Room_Number=? && Building_Name=? && Street_Name=?"

		_ = con.QueryRow(sqlStatement, NoResi, Room_Number, Building_Name, Street_Name).Scan(&packageID)

		nm := int64(0)

		sqlStatement = "SELECT count(IDDetailStatus) FROM detail_status ORDER BY co ASC "

		err := con.QueryRow(sqlStatement).Scan(&nm)

		nm = nm + 1

		temp := strconv.FormatInt(nm, 10)

		DST := "DS-" + temp

		var time1 = time.Now()
		date_sql := time1.Format("2006-01-02")

		sqlStatement = "INSERT INTO detail_status(co,iddetailstatus, idpacakage, idstatus, date) values(?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm, DST, packageID, "STAT-2", date_sql)

		//Update Status Package
		sqlStatement = "UPDATE package SET IDDetail=?, AdminID=? WHERE PackageID=?"

		stmt, err = con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(DST, AdminID, packageID)

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
		res.Message = "Not Found"
	}

	return res, nil
}

//Update-Status-Package (Resident)

//Delete-package
