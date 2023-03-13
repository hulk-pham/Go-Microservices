package model

import (
	"hulk/go-webservice/common"
)

type Message struct {
	common.Model
	UserId string `json:"user_id"`
	RoomId string `json:"room_id"`
	Text   string `json:"text"`
	Media  string `json:"media"`
}
