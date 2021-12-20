package db

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"qqbot-RSS-go/modles/config"
)

var DB *sql.DB

func InitDB() *sql.DB {
	data := config.GetConfig()
	marshal, err2 := json.Marshal(data)
	if err2 != nil {
		log.Fatal("初始化配置失败", err2.Error())
	}
	var dataConfig config.GetConfigData
	err1 := json.Unmarshal(marshal, &dataConfig)
	if err1 != nil {
		log.Fatal("初始化配置失败", err1.Error())
	}
	dsn := dataConfig.DbUser + ":" + dataConfig.DbPassWord + "@tcp(" + dataConfig.DbHost + ":" + dataConfig.DbPort + ")/" + dataConfig.DbName
	openDb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err = openDb.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Info("初始化连接数据库成功")
	return openDb
}

func InsertMsgId(msgInfo string, uri string, groupCode int, botUid int64) {
	updateExec, err := DB.Exec("update send_info as a inner join group_info as b on a.GroupId = b.GroupId inner join bot_info as c on b.BotId = c.BotId set a.MsgInfo = ? where a.Url =? and b.GroupCode = ? and c.BotUid = ?;", msgInfo, uri, groupCode, botUid)
	if err != nil {
		log.Errorf("更新异常%v:", err.Error())
	}
	affected, err := updateExec.RowsAffected()
	if err != nil {
		log.Warn(err.Error())
	}
	if affected == 0 {
		_, err := DB.Exec("insert into send_info (url, MsgInfo, groupid, createdatetime, updatetime)SELECT ?, ?, GroupId, now(), now() from group_info as a inner join bot_info as b on a.BotId = b.BotId where GroupCode = ? and b.BotUid = ?", uri, msgInfo, groupCode, botUid)
		if err != nil {
			log.Errorf("插入数据时发生异常:%v", err.Error())
		}
	}
}

func InsertUrl(uri string, uriName string, botUid int64, groupId int, userId int) bool {
	InsertExec, err := DB.Exec("insert into url_info (url, urlname, status, groupid, botid, createdatetime, rsstypeid, CreateUserId)SELECT ?, ?, 1, groupId, bi.BotId, now(), 1, ? from group_info as gi , bot_info as bi where bi.BotUid = ? and gi.GroupCode = ? and gi.BotId = bi.BotId and not exists(select 1 from url_info as ui where ui.GroupId = gi.GroupId and ui.BotId = bi.BotId and ui.Url =?)", uri, uriName, userId, botUid, groupId, uri)
	if err != nil {
		log.Errorf("发生异常:%v", err.Error())
	}
	affected, err := InsertExec.RowsAffected()
	if err != nil {
		log.Errorf("获取执行结果异常:%v", err.Error())
	}
	if affected == 1 {
		return true
	} else {
		return false
	}
}

func InsertRoom(roomCode int, roomName string, botUid int64, groupId int, userId int) bool {
	InsertExec, err := DB.Exec("insert into room_info (roomcode, botid, groupid, rsstypeid, status, createdatetime, roomname, CreateUserId)SELECT ?,bi.botid, gi.groupid, 2, 1, now(), ?, ? from bot_info as bi, group_info as gi where bi.BotUid = ? and gi.GroupCode = ? and gi.BotId = bi.BotId and not exists(select 1 from room_info as ri where ri.GroupId = gi.GroupId and ri.BotId = bi.BotId and ri.RoomCode = ? )", roomCode, roomName, userId, botUid, groupId, roomCode)
	if err != nil {
		log.Errorf("发生异常:%v", err.Error())
	}
	affected, err := InsertExec.RowsAffected()
	if err != nil {
		log.Errorf("获取执行结果异常:%v", err.Error())
	}
	if affected == 1 {
		return true
	} else {
		return false
	}
}

func DelRss(botUid int64, groupId int, urlName string, createUserId int) bool {
	DelExec, err := DB.Exec("delete ui from url_info as ui inner join bot_info bi on ui.BotId = bi.BotId inner join group_info gi on bi.BotId = gi.BotId where bi.BotUid = ? and gi.GroupCode = ? and ui.UrlName = ? and ui.CreateUserId = ? and ui.GroupId = gi.GroupId", botUid, groupId, urlName, createUserId)
	if err != nil {
		log.Errorf("发生异常:%v", err.Error())
	}
	affected, err := DelExec.RowsAffected()
	if err != nil {
		log.Errorf("获取执行结果异常:%v", err.Error())
	}
	if affected == 1 {
		return true
	} else {
		return false
	}
}

func DelLive(botUid int64, groupId int, roomCode string, createUserId int) bool {
	DelExec, err := DB.Exec("delete ri from room_info as ri inner join bot_info bi on ri.BotId = bi.BotId inner join group_info gi on bi.BotId = gi.BotId where bi.BotUid = ? and gi.GroupCode = ? and ri.RoomCode = ? and ri.CreateUserId = ? and ri.GroupId = gi.GroupId", botUid, groupId, roomCode, createUserId)
	if err != nil {
		log.Errorf("发生异常:%v", err.Error())
	}
	affected, err := DelExec.RowsAffected()
	if err != nil {
		log.Errorf("获取执行结果异常:%v", err.Error())
	}
	if affected == 1 {
		return true
	} else {
		return false
	}
}
