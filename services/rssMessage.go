package services

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/mmcdole/gofeed"
	"log"
	"qqbot-RSS-go/bot"
	"qqbot-RSS-go/db"
	"qqbot-RSS-go/modles/query"
	"qqbot-RSS-go/utils"
	"strings"
	"time"
)

func Sell(botId int64, mt int, ws *websocket.Conn) {
	urlData := query.Url(1, botId)
	var rssData query.GroupUrl
	for _, data := range urlData {
		err := json.Unmarshal([]byte(data), &rssData)
		if err != nil {
			log.Printf("解析错误:%v", err.Error())
		}
		fp := gofeed.NewParser()
		rspCode := utils.CheckCode(rssData.Url)
		if rspCode == 200 {
			feed, rssErr := fp.ParseURL(rssData.Url)
			if rssErr != nil {
				log.Printf("地址%v，连接错误:%v", rssData.Url, rssErr)
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
							log.Printf("机器人ID:%v 群ID:%v %v消息已通知 发布时间%v", botId, rssData.GroupCode, feed.Title, programTime)
						} else {
							log.Printf("机器人ID:%v 群ID:%v 开始检查订阅消息，检测到%v发布了一条新消息，发布时间%v触发通知", botId, rssData.GroupCode, feed.Title, programTime)
							db.InsertMsgId(message, rssData.Url, rssData.GroupCode, botId)
							bot.SendGroupMessageSocket(rssData.GroupCode, message, mt, ws)
						}
					} else {
						log.Printf("机器人ID:%v 群ID:%v 开始检查%v的订阅消息，未检测到新消息，上一条消息发布时间%v", botId, rssData.GroupCode, feed.Title, programTime)
					}
				}
			}
		}
	}
}

func BiliLive(botId int64, mt int, ws *websocket.Conn) {
	urlData := query.Url(2, botId)
	var liveData query.GroupUrl
	for _, data := range urlData {
		err := json.Unmarshal([]byte(data), &liveData)
		if err != nil {
			log.Fatalf("解析错误:%v", err.Error())
		}
		fp := gofeed.NewParser()
		resCode := utils.CheckCode(liveData.Url)
		if resCode == 200 {
			feed, rssErr := fp.ParseURL(liveData.Url)
			if rssErr != nil {
				log.Printf("连接错误:%v", rssErr)
			}
			if len(feed.Items) != 0 {
				for _, liveInfo := range feed.Items {
					log.Printf("群ID:%v 检查%v，已开播!开播时间:%v", liveData.GroupCode, feed.Title, utils.ReTime(liveInfo.Published))
					liveTime := utils.ReTime(liveInfo.Published)
					nowTime := time.Unix(time.Now().Unix()-600, 0).Format("2006-01-02 15:04:05")
					msgData := strings.Split(feed.Title, " ")[0] + `开播啦!\n` +
						"标题:" + strings.Split(liveInfo.Title, " ")[0] + `\n` +
						"链接:" + liveInfo.Link + `\n` +
						"开播时间:" + liveTime
					if liveTime > nowTime {
						liveMsg := query.SendInfo(liveData.Url, liveData.GroupCode, botId)
						if liveMsg == msgData {
							log.Printf("机器人ID:%v 群ID:%v 开播消息已通知", botId, liveData.GroupCode)
						} else {
							log.Printf("机器人ID:%v 群ID:%v 推送%v开播消息 开播时间:%v", botId, liveData.GroupCode, feed.Title, liveTime)
							db.InsertMsgId(msgData, liveData.Url, liveData.GroupCode, botId)
							bot.SendGroupMessageSocket(liveData.GroupCode, msgData, mt, ws)
						}
					}
				}
			} else {
				log.Printf("机器人ID:%v 群ID:%v 检查%v，未开播", botId, liveData.GroupCode, feed.Title)
			}
		}
	}
}