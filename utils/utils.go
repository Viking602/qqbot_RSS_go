package utils

import (
	"io"
	"log"
	"net/http"
	"time"
)

func ReTime(pub string) string {
	t1, _ := time.Parse(time.RFC1123, pub)
	return time.Unix(t1.Unix(), 0).Format("2006-01-02 15:04:05")
}

func CheckCode(uri string) int {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Printf("请求发生异常%v", err)
	}
	response, resErr := client.Do(req)
	if resErr != nil {
		log.Printf("发生异常%v", resErr)
	}
	defer func(Body io.ReadCloser) {
		closeErr := Body.Close()
		if closeErr != nil {
			log.Printf("关闭链接时发生异常:%v", closeErr.Error())
		}
	}(response.Body)
	return response.StatusCode
}
