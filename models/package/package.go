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

//Input-Package (V)
func Input_Package(PostmanID string, NoResi string, Name string,
	Street_Name string, Building_Name string, Room_Number string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT co FROM package ORDER BY co DESC LIMIT 1"

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

	sqlStatement = "SELECT co FROM detail_status ORDER BY co DESC LIMIT 1"

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

	sqlStatement = "UPDATE package SET IDDetail=? WHERE PackageID=?"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(DST, PackageID)

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read-Package (V)
func Read_Package(ID string, status int) (tools.Response, error) {
	var res tools.Response
	var Package st_package.Read_Package
	var arr_Package []st_package.Read_Package

	con := db.CreateCon()

	if status == 1 {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,StatusName FROM package join detail_status ds on package.IDDetail = ds.IDDetailStatus  join status_pack ON ds.IDStatus=status_pack.StatusID WHERE ResidentID=? && ds.IDStatus=? ORDER BY package.co ASC"

		rows, err := con.Query(sqlStatement, ID, "STAT-2")

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
	} else if status == 2 {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,StatusName FROM package join detail_status ds on package.IDDetail = ds.IDDetailStatus  join status_pack ON ds.IDStatus=status_pack.StatusID WHERE AdminID=? && ds.IDStatus=? ORDER BY package.co ASC"

		rows, err := con.Query(sqlStatement, ID, "STAT-2")

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
	} else if status == 3 {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,StatusName FROM package JOIN detail_status ds ON package.IDDetail = ds.IDDetailStatus JOIN status_pack ON status_pack.StatusID = ds.IDStatus WHERE PostmanID=? && StatusID=? ORDER BY package.co ASC"

		rows, err := con.Query(sqlStatement, ID, "STAT-1")

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

//Read-Package-History
func Read_Package_History(ID string, status int) (tools.Response, error) {
	var res tools.Response
	var Package st_package.Read_Package
	var arr_Package []st_package.Read_Package

	con := db.CreateCon()

	if status == 1 {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,StatusName FROM package join detail_status ds on package.IDDetail = ds.IDDetailStatus  join status_pack ON ds.IDStatus=status_pack.StatusID WHERE ResidentID=? && ds.IDStatus=? ORDER BY package.co ASC"

		rows, err := con.Query(sqlStatement, ID, "STAT-3")

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
	} else if status == 2 {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,StatusName FROM package join detail_status ds on package.IDDetail = ds.IDDetailStatus  join status_pack ON ds.IDStatus=status_pack.StatusID WHERE AdminID=? && ds.IDStatus!=? ORDER BY package.co ASC"

		rows, err := con.Query(sqlStatement, ID, "STAT-2")

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
	} else if status == 3 {

		sqlStatement := "SELECT PackageID,NoResi,Street_Name,StatusName FROM package join detail_status ds on package.IDDetail = ds.IDDetailStatus  join status_pack ON ds.IDStatus=status_pack.StatusID WHERE PostmanID=? && ds.IDStatus!=? ORDER BY  package.co ASC"

		rows, err := con.Query(sqlStatement, ID, "STAT-1")

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

//Read-Detail-Package (Resident) (tidak jadi digunakan)
func Read_Detail_Package_Resident(PackageID string) (tools.Response, error) {
	var res tools.Response
	var Package st_package.Detail_Resident
	var arr_Package []st_package.Detail_Resident

	con := db.CreateCon()

	sqlStatement := "SELECT date,StatusName FROM detail_status join status_pack s ON s.StatusID=detail_status.IDStatus where IDPacakage=?"

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

//Read-Detail-Package (V)
func Read_Detail_Package(PackageID string) (tools.Response, error) {
	var res tools.Response
	var Package st_package.Detail_Postman
	var arr_Package []st_package.Detail_Postman
	var St_Package st_package.Detail_Resident
	var ST_arr_Package []st_package.Detail_Resident

	con := db.CreateCon()

	sqlStatement := "SELECT PackageID, NoResi, package.Name, Street_Name, Building_Name,Street_Name,StatusName,p.Name,Room_Number FROM package JOIN postman p on package.PostmanID = p.PostmanID JOIN detail_status ds ON package.IDDetail = ds.IDDetailStatus JOIN status_pack sp on ds.IDStatus = sp.StatusID WHERE package.PackageID=?"

	_ = con.QueryRow(sqlStatement, PackageID).Scan(&Package.PackageID,
		&Package.Noresi, &Package.Name, &Package.Street_Name,
		&Package.Building_Name, &Package.Street_Name, &Package.Postman_Name,
		&Package.Current_Status, Package.Room_Number)

	fmt.Println(Package)

	sqlStatement = "SELECT date,StatusName FROM detail_status join status_pack s ON s.StatusID=detail_status.IDStatus where IDPacakage=?"

	rows, err := con.Query(sqlStatement, PackageID)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&St_Package.Date, &St_Package.Status_Name)
		if err != nil {
			return res, err
		}
		ST_arr_Package = append(ST_arr_Package, St_Package)
	}

	fmt.Println(ST_arr_Package)

	Package.Detail_Status = ST_arr_Package

	arr_Package = append(arr_Package, Package)

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

	sqlStatement := "SELECT co FROM detail_status ORDER BY co DESC LIMIT 1"

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

//Update-Status-Package (Admin) (V)
func Update_Status_Package_Admin(AdminID string, NoResi string, Name string,
	Street_Name string, Building_Name string, Room_Number string) (tools.Response, error) {
	var res tools.Response
	con := db.CreateCon()

	ResID := ""

	sqlStatement := "SELECT ResidentID FROM resident join building b on b.BuildingID = resident.BuildingID WHERE room_no=? && BuildingName=? && Address=? && CONCAT(name,' ',surname) LIKE ?"

	TN := "%" + Name + "%"

	fmt.Println(TN)
	_ = con.QueryRow(sqlStatement, Room_Number, Building_Name, Street_Name, TN).Scan(&ResID)

	fmt.Println(ResID)

	if ResID != "" {

		//Create detail status baru
		packageID := ""

		sqlStatement := "SELECT PackageID FROM package WHERE NoResi=? && Room_Number=? && Building_Name=? && Street_Name=?"

		_ = con.QueryRow(sqlStatement, NoResi, Room_Number, Building_Name, Street_Name).Scan(&packageID)

		nm := int64(0)

		sqlStatement = "SELECT co FROM detail_status ORDER BY co DESC LIMIT 1"

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
		sqlStatement = "UPDATE package SET IDDetail=?, AdminID=?, ResidentID=? WHERE PackageID=?"

		stmt, err = con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(DST, AdminID, ResID, packageID)

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

//Update-Status-Package (Resident) (V)
func Update_Status_Package_Resident(packageID string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := int64(0)

	sqlStatement := "SELECT co FROM detail_status ORDER BY co DESC LIMIT 1"

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

	_, err = stmt.Exec(nm, DST, packageID, "STAT-3", date_sql)

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

//Update-Data-Package
func Update_Data_Package(packageID string, NoResi string, Name string,
	Street_Name string, Building_Name string, Room_Number string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := ""

	sqlStatement := "SELECT StatusID FROM package JOIN detail_status ds on package.IDDetail = ds.IDDetailStatus JOIN status_pack sp on ds.IDStatus = sp.StatusID WHERE package.PackageID=? "

	err := con.QueryRow(sqlStatement).Scan(&nm)

	if err != nil {
		return res, err
	}

	if nm == "STAT-1" {

		sqlStatement = "UPDATE package SET NoResi=?,Name=?,Street_Name=?,Building_Name=?,Room_Number=? WHERE PackageID=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(NoResi, Name, Street_Name, Building_Name, Room_Number, packageID)

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
		res.Message = "Can't be Updated"
	}

	return res, nil
}
