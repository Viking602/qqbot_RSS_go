package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"qqbot-RSS-go/bot"
	"qqbot-RSS-go/modles/query"
	"strings"
)

func GroupMsg(message string, groupId int, botUid int64, ws *websocket.Conn, mt int) {
	switch message {
	case "rss-all":
		groupData := query.Group(groupId, botUid)
		var urlInfo query.GroupUrl
		var result []string
		for _, url := range groupData {
			_ = json.Unmarshal([]byte(url), &urlInfo)
			result = append(result, urlInfo.UrlName)
		}
		fmt.Println(strings.Join(result, ","))
		msgData := "当前订阅:\n" +
			strings.Join(result, "\n")
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-status":
		msgData := "正在运行"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-help":
		msgData := "help:\n" +
			"rss-all\t查询本群订阅信息\n" +
			"rss-about\t关于\n" +
			"rss-status\t运行状态" +
			"正在开发中的功能:\n" +
			"rss-init\t注册群信息\n" +
			"rss-reg --url\t添加订阅\n" +
			"rss-del --name\t删除订阅\n"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "rss-about":
		msgData := "about:\n" +
			"基于onebot-11标准实现\n" +
			"订阅消息来源于:\n" +
			"https://docs.rsshub.app \n" +
			"当前版本:DEV20211129\n" +
			"有问题请联系[CQ:at,qq=1900097700]"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	case "确实":
		msgData := "[CQ:image,file=48008a017562dce6bb7e77cceb7af475.image,url=https://gchat.qpic.cn/gchatpic_new/1900097700/725315770-3053488658-48008A017562DCE6BB7E77CCEB7AF475/0?term=3,subType=0]"
		bot.SendGroupMessageSocket(groupId, msgData, mt, ws)
	default:

	}
}
