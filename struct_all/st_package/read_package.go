package st_package

type Read_Package struct {
	Id_Package  string `json:"id_package"`
	No_Resi     string `json:"no_resi"`
	Street_Name string `json:"street_name"`
	Status      string `json:"status"`
}

type Read_Id_Package struct {
	Id_Package string `json:"id_package"`
}

type Read_Package_Notif struct {
	Id_Package  string `json:"id_package"`
	No_Resi     string `json:"no_resi"`
	Street_Name string `json:"street_name"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}
