package bilibili

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func LiveInfo(roomId string) []byte {
	params := url.Values{}
	Url, err := url.Parse("https://api.live.bilibili.com/room/v1/Room/get_info")
	if err != nil {
		log.Printf("err:%v", err)
	}
	params.Set("room_id", roomId)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, respErr := http.Get(urlPath)
	if respErr != nil {
		fmt.Println(respErr)
		return []byte(respErr.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("关闭连接时发生异常:%v", err)
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func GetUpInfo(mid string) []byte {
	params := url.Values{}
	Url, err := url.Parse("https://api.bilibili.com/x/space/acc/info")
	if err != nil {
		log.Printf("发生异常:%v", err.Error())
	}
	params.Set("mid", mid)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, respErr := http.Get(urlPath)
	if respErr != nil {
		log.Printf("发生异常:%v", respErr)
		return []byte(respErr.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("关闭连接时发生异常:%v", err)
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
