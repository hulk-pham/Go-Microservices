package dto

type WSRequest struct {
	Type   string `json:"type"`
	UserId string `json:"user_id"`
	RoomId string `json:"room_id"`

	Text string `json:"text"`
}
