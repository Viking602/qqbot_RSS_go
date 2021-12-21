package bilibili

import (
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func LiveInfo(roomId string) []byte {
	params := url.Values{}
	Url, err := url.Parse("https://api.live.bilibili.com/room/v1/Room/get_info")
	if err != nil {
		log.Errorf("err:%v", err)
	}
	params.Set("room_id", roomId)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, respErr := http.Get(urlPath)
	if respErr != nil {
		log.Error(respErr.Error())
		return []byte(respErr.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("关闭连接时发生异常:%v", err)
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func GetUpInfo(mid string) []byte {
	params := url.Values{}
	Url, err := url.Parse("https://api.bilibili.com/x/space/acc/info")
	if err != nil {
		log.Errorf("发生异常:%v", err.Error())
	}
	params.Set("mid", mid)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, respErr := http.Get(urlPath)
	if respErr != nil {
		log.Errorf("发生异常:%v", respErr)
		return []byte(respErr.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("关闭连接时发生异常:%v", err)
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
