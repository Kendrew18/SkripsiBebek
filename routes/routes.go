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

	//NDL
	UM.POST("/sign-up-postman", user_management.CreatePostman)

	return e
}
