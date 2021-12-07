package services

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/mmcdole/gofeed"
	"log"
	"qqbot-RSS-go/bot"
	"qqbot-RSS-go/db"
	"qqbot-RSS-go/modles/msg"
	"qqbot-RSS-go/modles/query"
	"qqbot-RSS-go/services/bilibili"
	"qqbot-RSS-go/utils"
	"strconv"
	"time"
)

func Sell(botId int64, mt int, ws *websocket.Conn) {
	urlData := query.Url(1, botId)
	var rssData query.GroupUrl
	for _, data := range urlData {
		err := json.Unmarshal([]byte(data), &rssData)
		if err != nil {
			log.Printf("解析错误:%v", err.Error())
			return
		}
		fp := gofeed.NewParser()
		rspCode := utils.CheckCode(rssData.Url)
		if rspCode == 200 {
			feed, rssErr := fp.ParseURL(rssData.Url)
			if rssErr != nil {
				log.Printf("地址%v，连接错误:%v", rssData.Url, rssErr)
				return
			}
			for nm, rssInfo := range feed.Items {
				if nm == 0 {
					programTime := utils.ReTime(rssInfo.Published)
					message := feed.Title + `\n` +
						"标题:" + rssInfo.Title + `\n` +
						"链接:" + rssInfo.Link + `\n` +
						"日期:" + programTime
					tm := time.Now().Unix() - 600
					nowTime := time.Unix(tm, 0).Format("2006-01-02 15:04:05")
					if programTime > nowTime {
						msgData := query.SendInfo(rssData.Url, rssData.GroupCode, botId)
						if msgData == message {
							log.Printf("BOT:%v 群ID:%v %v消息已通知 发布时间%v", botId, rssData.GroupCode, feed.Title, programTime)
						} else {
							log.Printf("BOT:%v 群ID:%v 开始检查订阅消息，检测到%v发布了一条新消息，发布时间%v触发通知", botId, rssData.GroupCode, feed.Title, programTime)
							db.InsertMsgId(message, rssData.Url, rssData.GroupCode, botId)
							bot.SendGroupMessageSocket(rssData.GroupCode, message, mt, ws)
						}
					} else {
						log.Printf("BOT:%v 群ID:%v 开始检查%v的订阅消息，未检测到新消息，上一条消息发布时间%v", botId, rssData.GroupCode, feed.Title, programTime)
					}
				}
			}
		}
	}
}

func NewBilLive(botUid int64, ws *websocket.Conn, mt int) {
	data := query.GetRoomCode(botUid)
	var roomCode query.RoomInfo
	for _, roomData := range data {
		roomCodeErr := json.Unmarshal([]byte(roomData), &roomCode)
		if roomCodeErr != nil {
			log.Printf("发生异常:%v", roomCodeErr)
			return
		}
		roomInfo := bilibili.LiveInfo(roomCode.RoomCode)
		var room msg.BiliLiveInfo
		err := json.Unmarshal(roomInfo, &room)
		if err != nil {
			log.Printf("序列化JSON发生异常:%v", err.Error())
			log.Printf("参数返回:%v", roomData)
			return
		}
		if room.Data.LiveStatus == 1 {
			nowTime := time.Unix(time.Now().Unix()-600, 0).Format("2006-01-02 15:04:05")
			if room.Data.LiveTime > nowTime {
				upData := bilibili.GetUpInfo(strconv.Itoa(room.Data.Uid))
				var upInfo msg.UpInfo
				upJsonErr := json.Unmarshal(upData, &upInfo)
				if upJsonErr != nil {
					log.Printf("发生异常:%v", upJsonErr.Error())
					log.Printf("参数返回:%v", upData)
					return
				}
				liveMsg := query.SendInfo(upInfo.Data.LiveRoom.Url, roomCode.GroupCode, botUid)
				message := `我是本群开播小助手！\n` + upInfo.Data.Name + `开播啦！\n` +
					`标题:` + room.Data.Title + `\n` +
					`分区:` + room.Data.AreaName + `\n` +
					`链接` + upInfo.Data.LiveRoom.Url + `\n` +
					`开播时间:` + room.Data.LiveTime
				if liveMsg == message {
					log.Printf("BOT:%v 群ID:%v 直播间ID:%v开播消息已通知", botUid, roomCode.GroupCode, room.Data.RoomId)
				} else {
					db.InsertMsgId(message, upInfo.Data.LiveRoom.Url, roomCode.GroupCode, botUid)
					bot.SendGroupMessageSocket(roomCode.GroupCode, message+`\n[CQ:image,file=`+upInfo.Data.LiveRoom.Cover+`]`, mt, ws)
				}
			} else {
				log.Printf("BOT:%v 群ID:%v 直播间ID:%v已开播", botUid, roomCode.GroupCode, room.Data.RoomId)
			}
		} else {
			log.Printf("BOT:%v 群ID:%v 直播间ID:%v未开播", botUid, roomCode.GroupCode, room.Data.RoomId)
		}
	}
}
