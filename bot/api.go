package bot

import (
	"github.com/gorilla/websocket"
	"strconv"
)

func SendGroupMessageSocket(groupId int, text string, mt int, ws *websocket.Conn) {
	msg := []byte(`{"action":"send_group_msg","params":{"group_id":` + strconv.Itoa(groupId) + `,"message":"` + text + `"}}`)
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		return
	}
}

func GetMsgSocket(messageId int, mt int, ws *websocket.Conn) {
	msg := []byte(`{"action":"get_msg","params:{"message_id:"` + strconv.Itoa(messageId) + `}}"`)
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		return
	}
}
