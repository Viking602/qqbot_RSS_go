package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"qqbot-RSS-go/bot"
	"qqbot-RSS-go/bot/handlers"
	"qqbot-RSS-go/modles/msg"
	"qqbot-RSS-go/modles/query"
	"qqbot-RSS-go/services/other"
	"strconv"
	"strings"
	"time"
)

func GroupMsg(message string, groupId int, botUid int64, userId int, role string, ws *websocket.Conn, mt int) {
	startTime := time.Now()
	msgInfo := strings.Split(message, " ")[0]
	switch msgInfo {
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
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws, false)
	case "#ping":
		msgData := fmt.Sprintf("%v", time.Since(startTime))
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws, false)
	case "rss-help":
		msgData := "帮助:\n" +
			"rss-all\t查询本群订阅信息\n" +
			"rss-about\t关于\n" +
			"rss-status\t运行状态\n" +
			"添加订阅 RSS格式URL\t添加RSS订阅\n" +
			"添加直播订阅 房间号(暂时仅支持B站)\t添加直播间订阅\n" +
			"删除订阅 订阅名称\t删除RSS订阅\n" +
			"删除直播订阅 房间号\t删除直播订阅\n" +
			"搜图 图片\t搜索图片(暂时仅支持saucenao)\n" +
			"点歌 歌曲名称\t目前仅支持网易云音乐\n" +
			"查询消息 消息ID\t查询撤回消息，参数为撤回后机器人提醒的消息ID，考虑隐私问题目前仅管理员及群主使用"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws, false)
	case "rss-about":
		msgData := "about:\n" +
			"项目地址:https://github.com/Viking602/qqbot_RSS_go"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws, false)
	case "确实":
		msgData := "[CQ:image,file=https://gchat.qpic.cn/gchatpic_new/1/0-0-68324A5E6ADD67F2F20A549ABD768D21/0?term=2,subType=1]"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws, false)
	case "添加订阅":
		uri := strings.Replace(message, "添加订阅 ", "", 1)
		newUri := strings.Replace(uri, "rsshub.app", "rss.vark.fun", 1)
		data := handlers.CommandAddRss(newUri, botUid, groupId, userId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws, false)
	case "添加直播订阅":
		roomCode := strings.Replace(message, "添加直播订阅 ", "", 1)
		data := handlers.CommandAddLive(roomCode, botUid, groupId, userId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws, false)
	case "删除订阅":
		urlName := strings.Replace(message, "删除订阅 ", "", 1)
		data := handlers.CommandDelRss(botUid, groupId, urlName, userId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws, false)
	case "删除直播订阅":
		roomCode := strings.Replace(message, "删除直播订阅 ", "", 1)
		data := handlers.CommandDelLive(botUid, groupId, roomCode, userId)
		bot.SendGroupMessageSocket(groupId, data, mt, ws, false)
	case "搜图":
		log.Infof("收到群%v,用户%v发来的搜图请求,接收BOT%v", groupId, userId, botUid)
		imgUrl := strings.Replace(message, "搜图", "", 1)
		if imgUrl != "" && imgUrl != " " {
			if strings.Contains(imgUrl, "]") {
				bot.SendGroupMessageSocket(groupId, "正在搜索请稍后...", mt, ws, false)
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
			bot.SendGroupMessageSocket(groupId, data, mt, ws, false)
		} else {
			bot.SendGroupMessageSocket(groupId, "使用方法:点歌 歌曲名称", mt, ws, false)
		}
	case "查询消息":
		log.Infof("收到群%v,用户%v发来的查询消息请求,接收BOT%v, 角色:%v", groupId, userId, botUid, role)
		if role == "owner" || role == "admin" {
			messageId := strings.Replace(message, "查询消息 ", "", 1)
			i, err := strconv.Atoi(messageId)
			if err != nil {
				log.Errorf(err.Error())
			}
			bot.SendGetMsg(i, mt, ws)
		} else {
			bot.SendGroupMessageSocket(groupId, "角色权限不足", mt, ws, false)
		}
	case "历史上的今天":
		msgData := handlers.CommandToday()
		t := time.Now()
		date := fmt.Sprintf("%d年%d月%d日", t.Year(), t.Month(), t.Day())
		bot.SendGroupMessageSocket(groupId, "今天是:"+date+msgData, mt, ws, false)
	case "吉田直树的回答":
		message := make([]string, 0)
		message = append(message,
			"[CQ:image,file=https://gchat.qpic.cn/gchatpic_new/1/0-0-4FB894AA4BF1507CA47420A0A229D64E/0?term=2,subType=0]",
			"[CQ:image,file=https://gchat.qpic.cn/gchatpic_new/1/0-0-E0A15CA4DBEDBD5D358B42007BAD3EF4/0?term=2,subType=0]")
		msgData := message[rand.Intn(len(message))]
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws, false)
	}
}

func MoyuMsg(mt int, ws *websocket.Conn) {
	groupId, _ := strconv.Atoi(os.Getenv("TMP_GROUPID"))
	data := other.FishMan()
	var fishmanmsg msg.FishermanMsg
	fisherr := json.Unmarshal(data, &fishmanmsg)
	msgData := fmt.Sprintf("[CQ:image,file=%s]", fishmanmsg.Data.MoyuUrl)
	bot.SendGroupMessageSocket(groupId, msgData, mt, ws, false)
	if fisherr != nil {
		log.Errorf("发生异常:%v", fisherr.Error())
		log.Infof("参数返回:%v", data)
		return
	}
}
