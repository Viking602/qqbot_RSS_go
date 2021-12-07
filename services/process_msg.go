package services

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"qqbot-RSS-go/bot"
	"qqbot-RSS-go/bot/handlers"
	"qqbot-RSS-go/modles/query"
	"strings"
)

func GroupMsg(message string, groupId int, botUid int64, ws *websocket.Conn, mt int) {
	msg := strings.Split(message, " ")[0]
	switch msg {
	case "rss-all":
		groupData := query.Group(groupId, botUid)
		liveGroupData := query.LiveGroup(groupId, botUid)
		var urlInfo query.GroupUrl
		var result []string
		for _, url := range groupData {
			_ = json.Unmarshal([]byte(url), &urlInfo)
			result = append(result, urlInfo.UrlName)
		}
		var roomName query.RoomName
		for _, liveInfo := range liveGroupData {
			_ = json.Unmarshal([]byte(liveInfo), &roomName)
			result = append(result, roomName.RoomName+"直播间")
		}
		msgData := "当前订阅:\n" +
			strings.Join(result, "\n")
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-status":
		msgData := "正在运行"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-help":
		msgData := "help:\n" +
			"rss-all\t查询本群订阅信息\n" +
			"rss-about\t关于\n" +
			"rss-status\t运行状态\n" +
			"rss-add {RssUrl}\t添加RSS订阅\n" +
			"rss-live {roomId}\t添加直播间订阅"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-about":
		msgData := "about:\n" +
			"当前版本:DEV20211207\n" +
			"有问题请联系QQ:1900097700"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "确实":
		msgData := "[CQ:image,file=48008a017562dce6bb7e77cceb7af475.image,url=https://gchat.qpic.cn/gchatpic_new/1900097700/725315770-3053488658-48008A017562DCE6BB7E77CCEB7AF475/0?term=3,subType=0]"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-add":
		uri := strings.Replace(message, "rss-add ", "", 1)
		data := handlers.AddRss(uri, botUid, groupId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws)
	case "rss-live":
		roomCode := strings.Replace(message, "rss-live ", "", 1)
		data := handlers.AddLive(roomCode, botUid, groupId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws)
	}
}
