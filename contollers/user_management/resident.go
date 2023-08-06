package user_management

import (
	"SkripsiBebek/models/user_management"
	"github.com/labstack/echo/v4"
	"net/http"
)

//Read_CSV_And_Sign_Up_Resident
func ReadCSV(c echo.Context) error {
	BuildingID := c.FormValue("buildingid")

	result, err := user_management.Read_CSV(c.Response(), c.Request(), BuildingID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//See_All_Resident
func SeeAllResident(c echo.Context) error {
	BuildingID := c.FormValue("buildingid")

	result, err := user_management.See_All_Resident(BuildingID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Delete Resident
func DeleteResident(c echo.Context) error {
	ResidentID := c.FormValue("residentid")

	result, err := user_management.Delete_Resident(ResidentID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Profile_Resident
func UpdateProfileResident(c echo.Context) error {
	ResidentID := c.FormValue("residentid")
	name := c.FormValue("name")
	surname := c.FormValue("surname")
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := user_management.Update_Profile_Resident(ResidentID, name,
		surname, email, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SeeProfileResident(c echo.Context) error {
	ResidentID := c.FormValue("residentid")

	result, err := user_management.See_Profile_Resident(ResidentID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
