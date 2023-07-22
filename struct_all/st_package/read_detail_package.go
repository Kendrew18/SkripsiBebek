package st_package

type Detail_Resident struct {
	Date        string `json:"date"`
	Status_Name string `json:"status_name"`
}

type Detail_Postman struct {
	PackageID     string `json:"package_id"`
	Noresi        string `json:"noresi"`
	Name          string `json:"name"`
	Street_Name   string `json:"street_name"`
	Building_Name string `json:"building_name"`
	Room_Number   string `json:"room_number"`
}