package user_management

import (
	"SkripsiBebek/models/user_management"
	"github.com/labstack/echo/v4"
	"net/http"
)

//Sign-up-postman
func CreatePostman(c echo.Context) error {
	Name := c.FormValue("name")
	Email := c.FormValue("email")
	Password := c.FormValue("password")

	result, err := user_management.Create_Postman(Name, Email, Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
