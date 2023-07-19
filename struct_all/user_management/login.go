package user_management

type Login struct {
	ID         string `json:"id"`
	Status     int    `json:"status"`
	BuildingID string `json:"building_id"`
}
