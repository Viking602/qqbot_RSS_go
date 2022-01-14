package handlers

import (
	"encoding/json"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"qqbot-RSS-go/db"
	"qqbot-RSS-go/modles/msg"
	"qqbot-RSS-go/services/bilibili"
	"qqbot-RSS-go/services/hiapi"
	"qqbot-RSS-go/services/img"
	"qqbot-RSS-go/utils"
	"strconv"
)

func CommandAddRss(uri string, botUid int64, groupId int, userId int) string {
	if uri != "rss-add" && uri != "" {
		fp := gofeed.NewParser()
		rspCode := utils.CheckCode(uri)
		if rspCode == 200 {
			feed, rssErr := fp.ParseURL(uri)
			if rssErr != nil {
				logrus.Error("非法RSS格式:%v", rssErr.Error())
				return "非法RSS格式"
			}
			result := db.InsertUrl(uri, feed.Title, botUid, groupId, userId)
			if result == true {
				return feed.Title + "订阅成功"
			} else {
				return feed.Title + "添加失败，订阅已经存在或注册异常"
			}
		} else {
			return uri + "无法访问或错误的URL"
		}
	} else {
		return "使用方法:rss-add RSS订阅URL\n建议使用https://rss.vark.fun获取RSS信息"
	}
}

func CommandAddLive(roomCode string, botUid int64, groupId int, userId int) string {
	if roomCode != "rss-live" && roomCode != "" {
		roomInfo := bilibili.LiveInfo(roomCode)
		var room msg.BiliLiveInfo
		err := json.Unmarshal(roomInfo, &room)
		if err != nil {
			logrus.Error("序列化JSON发生异常:%v", err.Error())
			return "房间号码错误"
		}
		upData := bilibili.GetUpInfo(strconv.Itoa(room.Data.Uid))
		var upInfo msg.UpInfo
		err = json.Unmarshal(upData, &upInfo)
		if err != nil {
			logrus.Error("序列化JSON发生异常:%v", err.Error())
			return "用户信息错误"
		}
		result := db.InsertRoom(room.Data.RoomId, upInfo.Data.Name, botUid, groupId, userId)
		if result == true {
			return upInfo.Data.Name + "直播间订阅成功"
		} else {
			return upInfo.Data.Name + "直播间订阅失败，已存在或注册异常"
		}
	} else {
		return "使用方法:rss-live bilibili直播间房间号"
	}
}

func CommandDelRss(botUid int64, groupId int, urlName string, createUserId int) string {
	if urlName != "rss-del" && urlName != "" {
		result := db.DelRss(botUid, groupId, urlName, createUserId)
		if result == true {
			return urlName + "取消订阅成功"
		} else {
			return urlName + "取消订阅失败，订阅不存在或权限不足"
		}
	} else {
		return "使用方法:rss-del 订阅名称"
	}
}

func CommandDelLive(botUid int64, groupId int, roomCode string, createUserId int) string {
	if roomCode != "rss-live-del" && roomCode != "" {
		result := db.DelLive(botUid, groupId, roomCode, createUserId)
		if result == true {
			return roomCode + "直播订阅取消成功"
		} else {
			return roomCode + "直播订阅取消失败，订阅不存在或权限不足"
		}
	} else {
		return "使用方法:rss-live-del bilibili直播房间号"
	}
}

func CommandNAO(url string) []string {
	logrus.Infof("开始对%v图片进行检索", url)
	data := img.SauceNAO(url)
	var saucenao msg.SauceNAO
	err := json.Unmarshal(data, &saucenao)
	if err != nil {
		logrus.Error("发生异常:%v", err.Error())
	}
	var result []string
	for _, imgs := range saucenao.Results {
		msgData := `[CQ:image,file=` + imgs.Header.Thumbnail + `]\n` +
			`相似度:` + imgs.Header.Similarity + `\n` +
			`标题:` + imgs.Data.Title + `\n` +
			`链接:`
		result = append(result, msgData)
		for _, i := range imgs.Data.ExtUrls {
			result = append(result, i)
			break
		}
	}
	return result
}

func CommandSearchMusic(s string) string {
	data := hiapi.Search163(s)
	var musicData msg.Search
	err := json.Unmarshal(data, &musicData)
	if err != nil {
		logrus.Errorf("发生异常:%v", err.Error())
	}
	var result string
	if musicData.Result.SongCount != 0 {
		for _, songs := range musicData.Result.Songs {
			msgData := `[CQ:music,type=163,id=` + strconv.Itoa(songs.Id) + `]`
			result = msgData
			break
		}
	} else {
		return "搜索失败结果为0"
	}
	return result
}
