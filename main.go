package main

import (
	"encoding/json"
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

func main() {
	config, _ := ioutil.ReadFile("config.yaml")
	err := yaml.Unmarshal(config, &setting)
	if err != nil {
		return
	}
	DB = InitDB()
	for true {
		urlData := queryUrl()
		var newData groupUrl
		for _, data := range urlData {
			err := json.Unmarshal([]byte(data), &newData)
			if err != nil {
				return
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
						log.Println(itm.Link, "发布时间在当前时间五分钟内，触发通知", "Time:", programTime, "发送QQ群ID:", newData.GroupCode)
						sendMsg(message, newData.GroupCode)
					} else {
						log.Println(itm.Link, "发布时间不在当前时间五分钟内，未触发通知", "Time", programTime, "预发送QQ群ID:", newData.GroupCode)
					}
				}
			}
		}
		time.Sleep(300 * time.Second)
	}

}
