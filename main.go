package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"qqbot-RSS-go/db"
	"qqbot-RSS-go/modles/config"
	"qqbot-RSS-go/modles/msg"
	"qqbot-RSS-go/services"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func socket(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("连接异常:%v", err)
		return
	}
	defer func(ws *websocket.Conn) {
		closeErr := ws.Close()
		if err != nil {
			log.Printf("关闭连接发生异常:%v", closeErr.Error())
		}
	}(ws)
	for {
		mt, message, msgErr := ws.ReadMessage()
		log.Printf("Socket消息:%v", string(message))
		if msgErr != nil {
			log.Printf("连接地址%v 已断开:%v", ws.RemoteAddr(), msgErr)
			break
		}
		// 初始化解析
		var data map[string]interface{}
		dataErr := json.Unmarshal(message, &data)
		if dataErr != nil {
			log.Printf("序列化JSON异常:%v", dataErr)
			break
		}
		//获取消息ID
		//msgData, msgErr := json.Marshal(data)
		//if msgErr != nil{
		//	log.Printf(msgErr.Error())
		//}
		//var MsgData msg.ResultMsg
		//jsonErr = json.Unmarshal(msgData, &MsgData)
		//if jsonErr != nil{
		//	log.Printf(jsonErr.Error())
		//}
		//消息ID获取结束
		postType := data["post_type"]
		switch postType {
		case "message":
			messageType := data["message_type"]
			switch messageType {
			case "group":
				var groupMsg msg.GroupMsg
				groupMsgErr := json.Unmarshal(message, &groupMsg)
				if groupMsgErr != nil {
					log.Printf("解析群消息异常:%v", groupMsgErr)
				}
				go services.GroupMsg(groupMsg.Message, groupMsg.GroupId, groupMsg.SelfId, ws, mt)
			}
		case "meta_event":
			metaEventType := data["meta_event_type"]
			switch metaEventType {
			case "heartbeat":
				var botInfo msg.LiveEvent
				UnErr := json.Unmarshal(message, &botInfo)
				if UnErr != nil {
					log.Printf("Error:%v", UnErr.Error())
				}
				go services.Sell(botInfo.SelfId, mt, ws)
				//go services.BiliLive(botInfo.SelfId, mt, ws)
				go services.NewBilLive(botInfo.SelfId, ws, mt)
			case "lifecycle":
				var lifecycle msg.Lifecycle
				lifecycleErr := json.Unmarshal(message, &lifecycle)
				if lifecycleErr != nil {
					log.Printf("解析异常:%v", lifecycleErr.Error())
				}
				log.Printf("UID:%v 已建立连接", lifecycle.SelfId)
			}
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

	router.GET("/ws", socket)
	err := router.Run(":" + config.Setting.QqBot.Port)
	if err != nil {
		log.Fatalf("启动失败请检查配置文件，发生异常:%v", err.Error())
	}

}
