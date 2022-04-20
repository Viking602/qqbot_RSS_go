package bot

import (
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func SendGroupMessageSocket(groupId int, text string, mt int, ws *websocket.Conn, autoEscape bool) {
	msg := []byte(fmt.Sprintf(`{"action":"send_group_msg","params":{"group_id":%v,"message":"%v","auto_escape":"%v"}}`, groupId, text, autoEscape))
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		log.Warnf("群聊消息发送失败%v", err.Error())
	}
}

func SendGroupForwardMsgSocket(groupId int, text string, mt int, ws *websocket.Conn) {
	msg := []byte(`{"action":"send_group_forward_msg","params":{"group_id":` + strconv.Itoa(groupId) + `,"messages":[{"type":"node","data":{"name":"搜图小助手","uin":"10086","content":"` + text + `"}}]}}`)
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		log.Warnf("群聊消息发送失败%v", err.Error())
	}
}

func SendGetMsg(messageId int, mt int, ws *websocket.Conn) {
	msg := []byte(fmt.Sprintf(`{"action":"get_msg","params":{"message_id":%v}}`, messageId))
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		log.Warnf("群聊消息发送失败%v", err.Error())
	}
}

func SendMsgSocket(groupId int, text string, mt int, ws *websocket.Conn) {
	msg := []byte(fmt.Sprintf(`{"action":"send_msg","params":{"group_id":%v,"message":"%v", "auto_escape": true}}`, groupId, text))
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		log.Warnf("群聊消息发送失败%v", err.Error())
	}
}

func SendPrivateMsgSocket(userId int64, message string, mt int, ws *websocket.Conn) {
	msg := []byte(fmt.Sprintf(`{"action":"send_private_msg","params":{"user_id":%v,"message":%v}}`, userId, message))
	fmt.Println(string(msg))
	err := ws.WriteMessage(mt, msg)
	if err != nil {
		log.Warnf("私聊消息发送失败%v", err.Error())
	}
}
