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
	//See_Profile_Resident
	UM.GET("/see-profile-pos", user_management.SeeProfilePostman)
	//Update-Profile-Resident
	UM.PUT("/update-profile-pos", user_management.UpdateProfilePostman)

	//Admin
	//Sign-Up-Admin-Bd
	UM.POST("/sign-up-admin-bd", user_management.CreateAdminAndBuilding)
	//See_Profile_Resident
	UM.GET("/see-profile-adm", user_management.SeeProfileAdmin)
	//Update-Profile-Resident
	UM.PUT("/update-profile-adm", user_management.UpdateProfileAdminBuilding)

	//Resident
	//Read-CSV
	UM.POST("/read-CSV", user_management.ReadCSV)
	//See_All_Resident
	UM.GET("/see-all-resident", user_management.SeeAllResident)
	//Delete-Resident
	UM.DELETE("/delete-resident", user_management.DeleteResident)
	//See_Profile_Resident
	UM.GET("/see-profile-res", user_management.SeeProfileResident)
	//Update-Profile-Resident
	UM.PUT("/update-profile-res", user_management.UpdateProfileResident)

	//Package
	//Input-Package
	PCK.POST("/input-package", _package.InputPackage)
	//Read-Package
	PCK.GET("/read-package", _package.ReadPackage)
	//Read-Package-History
	PCK.GET("/read-package-his", _package.ReadPackageHistory)
	//Read-Detail-Package-Resident
	PCK.GET("/read-det-pack-res", _package.ReadDetailPackageResident)
	//Read-Detail-Package-Postman
	PCK.GET("/read-det-pack", _package.ReadDetailPackage)

	//Update-Status-Package (Return Postman)
	PCK.POST("/update-return", _package.UpdateStatusPackage)
	//Update-Status-Package (Admin)
	PCK.POST("/update-stat-admin", _package.UpdateStatusPackageAdmin)
	//Update-Status-Package (Resident)
	PCK.POST("/update-stat-res", _package.UpdateStatusPackageResident)
	//Update-Data-Package
	PCK.PUT("/update-data-package", _package.UpdateDataPackage)

	return e
}
