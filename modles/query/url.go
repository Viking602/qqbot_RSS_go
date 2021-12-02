package query

import (
	"encoding/json"
	"log"
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

func Url(rssType int, userId int64) []string {
	var result []string
	row, err := db.DB.Query("select a.GroupCode, b.Url from group_info as a INNER JOIN url_info as b ON a.GroupId = b.GroupId inner join rss_type as c on b.RssTypeId = c.RssTypeId inner join bot_info as d on a.BotId = d.BotId where a.Status = 1 and b.status = 1 and d.Status = 1 and c.RssTypeId = ? and d.BotUid = ? and b.BotId = d.BotId", rssType, userId)
	if err != nil {
		log.Println(err)
	}
	for row.Next() {
		var gUrl GroupUrl
		err := row.Scan(&gUrl.GroupCode, &gUrl.Url)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
	}
	for row.Next() {
		var gUrl GroupUrl
		err := row.Scan(&gUrl.UrlName)
		if err != nil {
			log.Println(err)
		}
		dict, _ := json.Marshal(gUrl)
		data := string(dict)
		result = append(result, data)
	}
	return result
}

func SendInfo(uri string, groupCode int, botUid int64) string {
	var msg MsgInfo
	err := db.DB.QueryRow("select MsgInfo from send_info as a inner join group_info as b on  a.GroupId = b.GroupId inner join bot_info as c on b.BotId = c.BotId where a.Url = ? and b.GroupCode = ?  and c.BotUid = ?", uri, groupCode, botUid).Scan(&msg.Info)
	if err != nil {
		log.Println(err)
	}
	return msg.Info
}
