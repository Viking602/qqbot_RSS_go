package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"qqbot-RSS-go/modles/config"
)

var DB *sql.DB

func InitDB() *sql.DB {
	dsn := config.Setting.Mysql.User + ":" + config.Setting.Mysql.Password + "@tcp(" + config.Setting.Mysql.Host + ":" + config.Setting.Mysql.Port + ")/" + config.Setting.Mysql.Db
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
	log.Println("初始化连接数据库成功")
	return openDb
}

func InsertMsgId(msgInfo string, uri string, groupCode int, botUid int64) {
	updateExec, err := DB.Exec("update send_info as a inner join group_info as b on a.GroupId = b.GroupId inner join bot_info as c on b.BotId = c.BotId set a.MsgInfo = ? where a.Url =? and b.GroupCode = ? and c.BotUid = ?;", msgInfo, uri, groupCode, botUid)
	if err != nil {
		log.Printf("更新异常%v:", err.Error())
	}
	affected, err := updateExec.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if affected == 0 {
		_, err := DB.Exec("insert into send_info (url, MsgInfo, groupid, createdatetime, updatetime)SELECT ?, ?, GroupId, now(), now() from group_info as a inner join bot_info as b on a.BotId = b.BotId where GroupCode = ? and b.BotUid = ?", uri, msgInfo, groupCode, botUid)
		if err != nil {
			log.Printf("插入数据时发生异常:%v", err)
		}
	}
}
