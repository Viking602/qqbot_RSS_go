package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"qqbot-RSS-go/bot"
	"qqbot-RSS-go/bot/handlers"
	"qqbot-RSS-go/modles/query"
	"strings"
)

func GroupMsg(message string, groupId int, botUid int64, userId int, ws *websocket.Conn, mt int) {
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
		msgData := "帮助:\n" +
			"rss-all\t查询本群订阅信息\n" +
			"rss-about\t关于\n" +
			"rss-status\t运行状态\n" +
			"rss-add RSS格式URL\t添加RSS订阅\n" +
			"rss-live 房间号(暂时仅支持B站)\t添加直播间订阅\n" +
			"rss-del 订阅名称\t删除RSS订阅\n" +
			"rss-live-del 房间号\t删除直播订阅\n" +
			"开始搜图 图片\t搜索图片\n" +
			"点歌 歌曲名称\t目前仅支持网易云音乐"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-about":
		msgData := "about:\n" +
			"项目地址:https://github.com/Viking602/qqbot_RSS_go"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "确实":
		msgData := "[CQ:image,file=f3faacc3e754f0aa1261d0760f21ab1f.image,l=https://gchat.qpic.cn/gchatpic_new/1900097700/725315770-2601464787-F3FAACC3E754F0AA1261D0760F21AB1F/0?term=2,subType=1]"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-add":
		uri := strings.Replace(message, "rss-add ", "", 1)
		data := handlers.CommandAddRss(uri, botUid, groupId, userId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws)
	case "rss-live":
		roomCode := strings.Replace(message, "rss-live ", "", 1)
		data := handlers.CommandAddLive(roomCode, botUid, groupId, userId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws)
	case "rss-del":
		urlName := strings.Replace(message, "rss-del ", "", 1)
		data := handlers.CommandDelRss(botUid, groupId, urlName, userId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws)
	case "rss-live-del":
		roomCode := strings.Replace(message, "rss-live-del ", "", 1)
		data := handlers.CommandDelLive(botUid, groupId, roomCode, userId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws)
	case "搜图":
		log.Infof("收到群%v,用户%v发来的搜图请求,接收BOT%v", groupId, userId, botUid)
		imgUrl := strings.Replace(message, "搜图", "", 1)
		if imgUrl != "" && imgUrl != " " {
			if strings.Contains(imgUrl, "]") {
				bot.SendGroupMessageSocket(groupId, "正在搜索请稍后...", mt, ws)
				uri := strings.Split(imgUrl, ",")[2]
				result := handlers.CommandNAO(strings.Replace(uri, "url=", "", 1))
				data := strings.Join(result, "")
				bot.SendGroupForwardMsgSocket(groupId, data, mt, ws)
			}
		}
	case "点歌":
		log.Infof("收到群%v,用户%v发来的点歌请求,接收BOT%v", groupId, userId, botUid)
		musicName := strings.Replace(message, "点歌", "", 1)
		fmt.Println(musicName)
		if musicName != "" {
			data := handlers.CommandSearchMusic(musicName)
			bot.SendGroupMessageSocket(groupId, data, mt, ws)
		} else {
			bot.SendGroupMessageSocket(groupId, "使用方法:点歌 歌曲名称", mt, ws)
		}
	}
}
