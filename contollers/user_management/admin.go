package user_management

import (
	"SkripsiBebek/models/user_management"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Login(c echo.Context) error {
	Status := c.FormValue("status")
	Email := c.FormValue("email")
	Password := c.FormValue("password")

	st, _ := strconv.Atoi(Status)

	result, err := user_management.Login(Email, Password, st)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func CreateAdminAndBuilding(c echo.Context) error {
	Name := c.FormValue("name")
	Email := c.FormValue("email")
	Password := c.FormValue("password")
	BuildingName := c.FormValue("BuildingName")
	Address := c.FormValue("Address")
	Biography := c.FormValue("Biography")

	result, err := user_management.Create_Admin_And_Building(Email, Password,
		Name, BuildingName, Address, Biography)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
