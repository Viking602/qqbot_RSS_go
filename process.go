package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PrivateMsg struct {
	Font        int    `json:"font"`
	Message     string `json:"message"`
	MessageId   int    `json:"message_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	RawMessage  string `json:"raw_message"`
	SelfId      int64  `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
		UserId   int    `json:"user_id"`
	} `json:"sender"`
	SubType  string `json:"sub_type"`
	TargetId int64  `json:"target_id"`
	Time     int    `json:"time"`
	UserId   int    `json:"user_id"`
}

type GroupMsg struct {
	Anonymous   interface{} `json:"anonymous"`
	Font        int         `json:"font"`
	GroupId     int         `json:"group_id"`
	Message     string      `json:"message"`
	MessageId   int         `json:"message_id"`
	MessageSeq  int         `json:"message_seq"`
	MessageType string      `json:"message_type"`
	PostType    string      `json:"post_type"`
	RawMessage  string      `json:"raw_message"`
	SelfId      int64       `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserId   int    `json:"user_id"`
	} `json:"sender"`
	SubType string `json:"sub_type"`
	Time    int    `json:"time"`
	UserId  int    `json:"user_id"`
}

type LiveEvent struct {
	Interval      int    `json:"interval"`
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfId        int64  `json:"self_id"`
	Status        struct {
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
		Stat           struct {
			DisconnectTimes int `json:"disconnect_times"`
			LastMessageTime int `json:"last_message_time"`
			LostTimes       int `json:"lost_times"`
			MessageReceived int `json:"message_received"`
			MessageSent     int `json:"message_sent"`
			PacketLost      int `json:"packet_lost"`
			PacketReceived  int `json:"packet_received"`
			PacketSent      int `json:"packet_sent"`
		} `json:"stat"`
	} `json:"status"`
	Time int `json:"time"`
}

func groupMsg(message string, groupId int, userId int) {
	switch message {
	case "rss-all":
		groupData := queryGroup(groupId)
		var urlInfo groupUrl
		var result []string
		for _, url := range groupData {
			_ = json.Unmarshal([]byte(url), &urlInfo)
			result = append(result, urlInfo.UrlName)
		}
		fmt.Println(strings.Join(result, ","))
		msgData := "当前订阅:\n" +
			strings.Join(result, "\n")
		sendMsg(msgData, groupId)
	case "rss-status":
		msgData := "正在运行"
		sendMsg(msgData, groupId)
	case "rss-help":
		msgData := "帮助\n" +
			"rss-all\t查询本群订阅信息\n" +
			"rss-about\t关于\n" +
			"rss-status\t运行状态" +
			"正在开发中的功能:\n" +
			"rss-init\t注册群信息\n" +
			"rss-reg --url\t添加订阅\n" +
			"rss-del --name\t删除订阅\n"
		sendMsg(msgData, groupId)
	case "rss-about":
		msgData := "关于本插件:\n" +
			"依托于go-cqhttp运行\n" +
			"订阅消息来源于:\n" +
			"https://docs.rsshub.app \n" +
			"当前版本:DEV20211125\n" +
			"有问题请联系[CQ:at,qq=1900097700]"
		sendMsg(msgData, groupId)
	default:

	}
}
