package _package

import (
	_package "SkripsiBebek/models/package"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input_Package
func InputPackage(c echo.Context) error {
	PostmanID := c.FormValue("postman_id")
	No_resi := c.FormValue("no_resi")
	Name := c.FormValue("name")
	Street_Name := c.FormValue("street_name")
	Building_Name := c.FormValue("building_name")
	Room_Number := c.FormValue("room_number")

	result, err := _package.Input_Package(PostmanID, No_resi, Name,
		Street_Name, Building_Name, Room_Number)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Package
func ReadPackage(c echo.Context) error {
	ID := c.FormValue("id")
	Status := c.FormValue("status")

	st, _ := strconv.Atoi(Status)

	result, err := _package.Read_Package(ID, st)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Package_History
func ReadPackageHistory(c echo.Context) error {
	ID := c.FormValue("id")
	Status := c.FormValue("status")

	st, _ := strconv.Atoi(Status)

	result, err := _package.Read_Package_History(ID, st)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Detail_Package_Resident
func ReadDetailPackageResident(c echo.Context) error {
	PackageID := c.FormValue("package_id")

	result, err := _package.Read_Detail_Package_Resident(PackageID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Detail_Package
func ReadDetailPackage(c echo.Context) error {
	PackageID := c.FormValue("package_id")

	result, err := _package.Read_Detail_Package(PackageID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Status_Package (Return Postman)
func UpdateStatusPackage(c echo.Context) error {
	PackageID := c.FormValue("package_id")

	result, err := _package.Update_Status_Package(PackageID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update-Status-Package (Admin)
func UpdateStatusPackageAdmin(c echo.Context) error {
	AdminID := c.FormValue("adminid")
	No_resi := c.FormValue("no_resi")
	Name := c.FormValue("name")
	Street_Name := c.FormValue("street_name")
	Building_Name := c.FormValue("building_name")
	Room_Number := c.FormValue("room_number")

	result, err := _package.Update_Status_Package_Admin(AdminID, No_resi,
		Name, Street_Name, Building_Name, Room_Number)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Status_Package_Resident
func UpdateStatusPackageResident(c echo.Context) error {
	PackageID := c.FormValue("package_id")

	result, err := _package.Update_Status_Package_Resident(PackageID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Data_Package
func UpdateDataPackage(c echo.Context) error {
	PackageID := c.FormValue("package_id")
	No_resi := c.FormValue("no_resi")
	Name := c.FormValue("name")
	Street_Name := c.FormValue("street_name")
	Building_Name := c.FormValue("building_name")
	Room_Number := c.FormValue("room_number")

	result, err := _package.Update_Data_Package(PackageID, No_resi, Name,
		Street_Name, Building_Name, Room_Number)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
