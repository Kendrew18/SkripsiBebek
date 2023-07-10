package routes

import (
	"SkripsiBebek/contollers/user_management"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-NDL")
	})

	UM := e.Group("/UM")

	//User Management
	//Login
	UM.GET("/login", user_management.Login)
	//Sign-Up-Postman
	UM.POST("/sign-up-postman", user_management.CreatePostman)
	//Sign-Up-Admin-Bd
	UM.POST("/sign-up-admin-bd", user_management.CreateAdminAndBuilding)

	return e
}
