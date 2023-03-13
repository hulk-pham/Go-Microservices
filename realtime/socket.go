package realtime

import (
	"encoding/json"
	"fmt"
	"hulk/go-webservice/core/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	OpenRoom   string = "open-room"
	CloseRoom         = "close-room"
	DeleteRoom        = "delete-room"
	Messages          = "messages"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WShandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic("Failed to set websocket upgrade")
		}

		for {
			t, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			var wsRequest dto.WSRequest
			json.Unmarshal([]byte(msg), &wsRequest)

			switch msgType := wsRequest.Type; msgType {
			case OpenRoom:
				openRoomHandler(conn, wsRequest, t)
			case CloseRoom:
			case DeleteRoom:
			case Messages:
				sendMessagesHandler(conn, wsRequest, t)
			default:
				fmt.Printf("Unsupport type: %s.\n", msgType)
			}
			writeContent, _ := json.Marshal(wsRequest)
			conn.WriteMessage(t, writeContent)
		}
	}

}

func openRoomHandler(conn *websocket.Conn, wsRequest dto.WSRequest, t int) {
	roomid := wsRequest.RoomId
	listener := RoomManager.OpenListener(roomid)
	defer RoomManager.CloseListener(roomid, listener)
}

func sendMessagesHandler(conn *websocket.Conn, wsRequest dto.WSRequest, t int) {
	RoomManager.Submit(wsRequest.UserId, wsRequest.RoomId, wsRequest.Text)
}
