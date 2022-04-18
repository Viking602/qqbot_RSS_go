package query

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"qqbot-RSS-go/db"
)

type GroupUrl struct {
	GroupCode int
	Url       string
	UrlName   string
}

type MsgInfo struct {
	Info string
}

type Owner struct {
	OwnerUid int64
}

type RoomName struct {
	RoomName string
}

type RoomInfo struct {
	RoomCode  string
	GroupCode int
}

func Url(rssType int, userId int64) []string {
	var result []string
	row, err := db.DB.Query("select a.GroupCode, b.Url from group_info as a INNER JOIN url_info as b ON a.GroupId = b.GroupId inner join bot_info as d on a.BotId = d.BotId where a.Status = 1 and b.status = 1 and d.Status = 1 and b.RssTypeId = ? and d.BotUid = ? and b.BotId = d.BotId", rssType, userId)
	if err != nil {
		log.Error(err)
	}
	for row.Next() {
		var gUrl GroupUrl
		err := row.Scan(&gUrl.GroupCode, &gUrl.Url)
		if err != nil {
			log.Error(err)
		}
		dict, _ := json.Marshal(gUrl)
		data := string(dict)
		result = append(result, data)
	}
	return result
}

func Group(groupId int, botUid int64) []string {
	var result []string
	row, err := db.DB.Query("select b.UrlName from group_info as a INNER JOIN url_info as b ON a.GroupId = b.GroupId inner join bot_info as c on a.BotId = c.BotId where a.Status = 1 and b.status = 1 and a.GroupCode = ? and c.BotUid = ? and b.BotId = a.BotId", groupId, botUid)
	if err != nil {
		log.Error(err)
	}
	for row.Next() {
		var gUrl GroupUrl
		err := row.Scan(&gUrl.UrlName)
		if err != nil {
			log.Error(err)
		}
		dict, _ := json.Marshal(gUrl)
		data := string(dict)
		result = append(result, data)
	}
	return result
}

func LiveGroup(groupId int, botUid int64) []string {
	var result []string
	row, err := db.DB.Query("select RoomName from room_info as ri inner join bot_info bi on ri.BotId = bi.BotId inner join group_info gi on bi.BotId = gi.BotId where bi.BotUid = ? and gi.GroupCode = ? and ri.Status = 1 and gi.Status = 1 and bi.Status = 1 and ri.GroupId = gi.GroupId", botUid, groupId)
	if err != nil {
		log.Error(err)
	}
	for row.Next() {
		var roomName RoomName
		err := row.Scan(&roomName.RoomName)
		if err != nil {
			log.Error(err)
		}
		dict, _ := json.Marshal(roomName)
		data := string(dict)
		result = append(result, data)
	}
	return result
}

func SendInfo(uri string, groupCode int, botUid int64) string {
	var msg MsgInfo
	err := db.DB.QueryRow("select MsgInfo from send_info as a inner join group_info as b on  a.GroupId = b.GroupId inner join bot_info as c on b.BotId = c.BotId where a.Url = ? and b.GroupCode = ?  and c.BotUid = ?", uri, groupCode, botUid).Scan(&msg.Info)
	if err != nil {
		log.Error(err)
	}
	return msg.Info
}

func GetRoomCode(botUid int64) []string {
	var result []string
	row, err := db.DB.Query("select RoomCode,c.GroupCode from room_info as a inner join bot_info as b on a.BotId = b.BotId inner join group_info as c on a.GroupId = c.GroupId where b.BotUid = ? and RssTypeId = 2", botUid)
	if err != nil {
		log.Errorf("查询房间号时出现异常:%v:", err.Error())
		return nil
	}
	for row.Next() {
		var room RoomInfo
		err := row.Scan(&room.RoomCode, &room.GroupCode)
		if err != nil {
			log.Errorf("查询房间号时出现异常:%v", err.Error())
		}
		dict, _ := json.Marshal(room)
		result = append(result, string(dict))
	}
	return result
}

func GetBotOwner(botUid int64) int64 {
	var msg Owner
	err := db.DB.QueryRow("SELECT OwnerUid FROM `bot_info` WHERE BotUid = ?", botUid).Scan(&msg.OwnerUid)
	if err != nil {
		log.Error(err)
	}
	return msg.OwnerUid
}
