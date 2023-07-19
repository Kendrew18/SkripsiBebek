package routes

import (
	_package "SkripsiBebek/contollers/package"
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
	PCK := e.Group("/PCK")

	//User Management
	//Login
	UM.GET("/login", user_management.Login)

	//Postman
	//Sign-Up-Postman
	UM.POST("/sign-up-postman", user_management.CreatePostman)

	//Admin
	//Sign-Up-Admin-Bd
	UM.POST("/sign-up-admin-bd", user_management.CreateAdminAndBuilding)

	//Resident
	//Read-CSV
	UM.POST("/read-CSV", user_management.ReadCSV)
	//See_All_Resident
	UM.GET("/see-all-resident", user_management.SeeAllResident)

	//Package
	//Input-Package
	PCK.POST("/input-package", _package.InputPackage)
	//Read-Package
	PCK.GET("/read-package", _package.ReadPackage)
	//Read-Detail-Package-Resident
	PCK.GET("/read-det-pack-res", _package.ReadDetailPackageResident)
	//Read-Detail-Package-Postman
	PCK.GET("/read-det-pack-post", _package.ReadDetailPackagePostman)

	return e
}
