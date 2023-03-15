package entities

type Message struct {
	Model
	UserId string `json:"user_id"`
	RoomId string `json:"room_id"`
	Text   string `json:"text"`
	Media  string `json:"media"`
}
