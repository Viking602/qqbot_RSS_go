package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"qqbot-RSS-go/bot"
	"qqbot-RSS-go/db"
	"qqbot-RSS-go/modles/config"
	"qqbot-RSS-go/modles/msg"
	"qqbot-RSS-go/modles/query"
	"qqbot-RSS-go/services"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func socket(c *gin.Context) {
	var cronType = 0
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Errorf("连接异常:%v", err)
		return
	}
	defer func(ws *websocket.Conn) {
		closeErr := ws.Close()
		if err != nil {
			log.Errorf("关闭连接发生异常:%v", closeErr.Error())
		}
	}(ws)
	for {
		mt, message, msgErr := ws.ReadMessage()
		log.Infof("Socket消息:%v", string(message))
		fmt.Println(string(message))
		if msgErr != nil {
			log.Errorf("连接地址%v 已断开:%v", ws.RemoteAddr(), msgErr)
			cronType = 0
			services.CronInfo(mt, ws, cronType)
			break
		}
		// 初始化解析
		var data map[string]interface{}
		dataErr := json.Unmarshal(message, &data)
		if dataErr != nil {
			log.Errorf("序列化JSON异常:%v", dataErr)
			break
		}

		var ret map[string]interface{}
		retErr := json.Unmarshal(message, &ret)
		if retErr != nil {
			log.Errorf("序列化返回消息异常%v", retErr.Error())
		}
		if ret["status"] == "ok" {
			var retMsg msg.RetMsg
			retMsgErr := json.Unmarshal(message, &retMsg)
			if retMsgErr != nil {
				log.Errorf("解析消息失败%v", retMsgErr.Error())
			}
			msgData := fmt.Sprintf("%v撤回消息内容为:%v", retMsg.Data.Sender.Nickname, retMsg.Data.Message)
			bot.SendMsgSocket(retMsg.Data.GroupId, msgData, mt, ws)
		}

		postType := data["post_type"]
		switch postType {
		case "message":
			messageType := data["message_type"]
			switch messageType {
			case "group":
				var groupMsg msg.GroupMsg
				groupMsgErr := json.Unmarshal(message, &groupMsg)
				if groupMsgErr != nil {
					log.Errorf("解析群消息异常:%v", groupMsgErr)
				}
				go services.GroupMsg(groupMsg.Message, groupMsg.GroupId, groupMsg.SelfId, groupMsg.Sender.UserId, groupMsg.Sender.Role, ws, mt)
			}
		case "meta_event":
			metaEventType := data["meta_event_type"]
			switch metaEventType {
			case "heartbeat":
				var botInfo msg.LiveEvent
				UnErr := json.Unmarshal(message, &botInfo)
				log.Infoln("Test")
				if UnErr != nil {
					log.Errorf("Error:%v", UnErr.Error())
				}
				log.Infoln("Test2")
				go services.Rss(botInfo.SelfId, mt, ws)
				go services.BilLive(botInfo.SelfId, ws, mt)
			case "lifecycle":
				var lifecycle msg.Lifecycle
				lifecycleErr := json.Unmarshal(message, &lifecycle)
				if lifecycleErr != nil {
					log.Errorf("解析异常:%v", lifecycleErr.Error())
				}
				log.Infof("UID:%v 已建立连接", lifecycle.SelfId)
			}
		case "notice":
			noticeType := data["notice_type"]
			switch noticeType {
			case "group_recall":
				var groupRecallMsg msg.GroupRecall
				groupRecallMsgErr := json.Unmarshal(message, &groupRecallMsg)
				if groupRecallMsgErr != nil {
					log.Errorf("解析异常:%v", groupRecallMsgErr.Error())
				}
				msgData := fmt.Sprintf("用户%v 撤回了一条消息，消息ID:%v", groupRecallMsg.UserId, groupRecallMsg.MessageId)
				bot.SendGroupMessageSocket(groupRecallMsg.GroupId, msgData, mt, ws, true)
			case "group_decrease":
				var groupDecrease msg.GroupDecrease
				groupDecreaseErr := json.Unmarshal(message, &groupDecrease)
				if groupDecreaseErr != nil {
					log.Errorf("解析异常:%v", groupDecreaseErr.Error())
				}
				msgData := fmt.Sprintf("用户%v 离开了本群", groupDecrease.UserId)
				bot.SendGroupMessageSocket(groupDecrease.GroupId, msgData, mt, ws, false)
			}
		case "request":
			var requestData msg.RequestType
			requestDataErr := json.Unmarshal(message, &requestData)
			if requestDataErr != nil {
				log.Errorf("序列化JSON异常:%v", dataErr)
				break
			}
			requestType := requestData.RequestType
			switch requestType {
			case "friend":
				msgData := fmt.Sprintf("收到好友请求:%v", requestData.UserId)
				fmt.Println(msgData)
				ownerId := query.GetBotOwner(requestData.SelfId)
				bot.SendPrivateMsgSocket(ownerId, msgData, mt, ws)
			case "group":
				msgData := fmt.Sprintf("收到群邀请:%v", requestData.Flag)
				fmt.Println(msgData)
				ownerId := query.GetBotOwner(requestData.SelfId)
				bot.SendPrivateMsgSocket(ownerId, msgData, mt, ws)
			}
		}
		if cronType > 0 {
			continue
		} else {
			cronType = 1
			go services.CronInfo(mt, ws, cronType)
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s [%s] \"%s %s %d %s %s \"%s\" %s\"\n",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ClientIP,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	config.InitConfig("config.yaml")
	db.DB = db.InitDB()

	data := config.GetConfig()
	marshal, err2 := json.Marshal(data)
	if err2 != nil {
		log.Fatalf("解析配置文件失败:%v", err2.Error())
	}
	var dataConfig config.GetConfigData
	err1 := json.Unmarshal(marshal, &dataConfig)
	if err1 != nil {
		log.Fatalf("初始化配置失败:%v", err1.Error())
	}
	router.GET("/ws", socket)
	err := router.Run(":" + dataConfig.BotPort)
	if err != nil {
		log.Fatalf("启动失败请检查配置文件，发生异常:%v", err.Error())
	}

}
