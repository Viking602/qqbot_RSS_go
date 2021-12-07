package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"qqbot-RSS-go/db"
	"qqbot-RSS-go/modles/msg"
	"qqbot-RSS-go/services/bilibili"
	"qqbot-RSS-go/utils"
	"strconv"
)

func AddRss(uri string, botUid int64, groupId int) string {
	fp := gofeed.NewParser()
	rspCode := utils.CheckCode(uri)
	if rspCode == 200 {
		feed, rssErr := fp.ParseURL(uri)
		if rssErr != nil {
			log.Printf("非法RSS格式:%v", rssErr.Error())
			return "非法RSS格式"
		}
		result := db.InsertUrl(uri, feed.Title, botUid, groupId)
		if result == true {
			return feed.Title + "添加成功"
		} else {
			return feed.Title + "添加失败，订阅已经存在或注册异常"
		}
	} else {
		return uri + "无法访问或错误的URL"
	}
}

func AddLive(roomCode string, botUid int64, groupId int) string {
	roomInfo := bilibili.LiveInfo(roomCode)
	var room msg.BiliLiveInfo
	fmt.Println(room.Code)
	err := json.Unmarshal(roomInfo, &room)
	if err != nil {
		log.Printf("序列化JSON发生异常:%v", err.Error())
		return "房间号码错误"
	}
	upData := bilibili.GetUpInfo(strconv.Itoa(room.Data.Uid))
	var upInfo msg.UpInfo
	err = json.Unmarshal(upData, &upInfo)
	if err != nil {
		log.Printf("序列化JSON发生异常:%v", err.Error())
		return "用户信息错误"
	}
	result := db.InsertRoom(room.Data.RoomId, upInfo.Data.Name, botUid, groupId)
	if result == true {
		return upInfo.Data.Name + "直播间订阅成功"
	} else {
		return upInfo.Data.Name + "直播间订阅失败，已存在或注册异常"
	}

}
