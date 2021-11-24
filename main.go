package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func reTime(pub string) string {
	t1, _ := time.Parse(time.RFC1123, pub)
	return time.Unix(t1.Unix(), 0).Format("2006-01-02 15:04:05")
}

func sendMsg(msg string, groupId int) {
	params := url.Values{}
	path := "/send_msg?"
	Url, err := url.Parse(setting.QqBot.Url + path)
	if err != nil {
		return
	}
	params.Set("access_token", setting.QqBot.AccessToken)
	params.Set("message_type", "group")
	params.Set("group_id", strconv.Itoa(groupId))
	params.Set("message", msg)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	log.Println(resp.Status)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
}

func reports(c *gin.Context) {
	jsonData := make(map[string]interface{})
	err := c.BindJSON(&jsonData)
	if err != nil {
		log.Fatal("异常消息：", err)
	}
	postType := jsonData["post_type"]
	switch postType {
	case "message":
		messageType := jsonData["message_type"]
		switch messageType {
		case "private":
			private, privateErr := json.Marshal(jsonData)
			var privateData PrivateMsg
			if privateErr != nil {
				log.Fatalf("初始化私聊消息异常:%v", privateErr.Error())
			}
			jsonErr := json.Unmarshal(private, &privateData)
			if jsonErr != nil {
				log.Fatalf("序列化私聊消息JSON异常%v:", jsonErr.Error())
			}
			fmt.Println(privateData.Message, privateData.Sender.Nickname, privateData.Sender.UserId)
		case "group":
			group, groupErr := json.Marshal(jsonData)
			if groupErr != nil {
				log.Fatalf("初始化群消息异常:%v", groupErr.Error())
			}
			var groupData GroupMsg
			jsonErr := json.Unmarshal(group, &groupData)
			if jsonErr != nil {
				log.Fatalf("序列化群消息JSON异常:%v", jsonErr.Error())
			}
			groupMsg(groupData.Message, groupData.GroupId, groupData.Sender.UserId)
			fmt.Println(groupData.GroupId, groupData.Message, groupData.Sender.Nickname, groupData.Sender.UserId)
		}
	case "meta_event":
		mateEvent, _ := json.Marshal(jsonData)
		var mateEventData LiveEvent
		jsonErr := json.Unmarshal(mateEvent, &mateEventData)
		if jsonErr != nil {
			log.Fatalf("序列化JSON异常:%v", jsonErr.Error())
		}
		log.Println("最后一次接收消息时间:", time.Unix(int64(mateEventData.Status.Stat.LastMessageTime), 0).Format("2006-01-02 15:04:05"))

	}

}

func sell() {
	for true {
		urlData := queryUrl()
		var newData groupUrl
		for _, data := range urlData {
			err := json.Unmarshal([]byte(data), &newData)
			if err != nil {
				log.Fatalf("数据异常:%v", err.Error())
			}
			fp := gofeed.NewParser()
			feed, _ := fp.ParseURL(newData.Url)
			for nm, itm := range feed.Items {
				if nm == 0 {
					programTime := reTime(itm.Published)
					message := feed.Title + "\n" +
						"标题:" + itm.Title + "\n" +
						"链接:" + itm.Link + "\n" +
						"日期:" + programTime
					tm := time.Now().Unix() - 300
					nowTime := time.Unix(tm, 0).Format("2006-01-02 15:04:05")
					if programTime > nowTime {
						log.Println(itm.Link, "查看第一条消息", "发布时间:", programTime, "触发通知，群ID:", newData.GroupCode)
						fmt.Println(message, newData.GroupCode)
					} else {
						log.Println(itm.Link, "查看第一条消息", "发布时间", programTime, "未触发通知，群ID:", newData.GroupCode)
					}
				}
			}
		}
		time.Sleep(300 * time.Second)
	}
}

func main() {
	route := gin.Default()
	config, _ := ioutil.ReadFile("config.yaml")
	err := yaml.Unmarshal(config, &setting)
	if err != nil {
		log.Fatalf("初始化配置失败:%v", err.Error())
	}
	DB = InitDB()
	go sell()

	v1 := route.Group("v1")
	{
		v1.POST("/reports", reports)
	}
	err = route.Run()
	if err != nil {
		log.Fatalf("启动失败:%v", err.Error())
	}

}
