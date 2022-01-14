package bot

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func SendGroupMessageSocket(groupId int, text string, mt int, ws *websocket.Conn) {
	msg := []byte(`{"action":"send_group_msg","params":{"group_id":` + strconv.Itoa(groupId) + `,"message":"` + text + `"}}`)
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
