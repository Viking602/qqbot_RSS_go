package bot

import (
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

func SendGroupMessageSocket(groupId int, text string, mt int, ws *websocket.Conn) {
	msg := []byte(`{"action":"send_group_msg","params":{"group_id":` + strconv.Itoa(groupId) + `,"message":"` + text + `"}}`)
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		log.Printf("群聊消息发送失败%v", err.Error())
	}
}

func GetMsgSocket(messageId int, mt int, ws *websocket.Conn) {
	msg := []byte(`{"action":"get_msg","params:{"message_id:"` + strconv.Itoa(messageId) + `}}"`)
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		log.Printf("获取消息失败%v", err.Error())
	}
}

func SendPrivateMessageSocket(userId int, text string, mt int, ws *websocket.Conn) {
	msg := []byte(`{"action":"send_group_msg","params":{"private_id":` + strconv.Itoa(userId) + `,"message":"` + text + `"}}`)
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		log.Printf("私聊消息发送失败%v", err.Error())
	}
}
