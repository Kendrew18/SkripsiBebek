package user_management

type Read_All_Resident struct {
	ResidentID string `json:"resident_id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Room_No    string `json:"room_no"`
	Email      string `json:"email"`
}
