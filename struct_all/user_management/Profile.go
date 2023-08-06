package user_management

type Admin_Profile struct {
	Admin_id      string `json:"admin_id"`
	Building_id   string `json:"building_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Building_Name string `json:"building_name"`
	Address       string `json:"address"`
	Biography     string `json:"biography"`
}

type Postman_Profile struct {
	Postman_id string `json:"postman_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Resident_Profile struct {
	Resident_id string `json:"resident_id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
