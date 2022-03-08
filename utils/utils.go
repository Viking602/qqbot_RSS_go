package utils

import (
	log "github.com/sirupsen/logrus"
	"io"
	"math"
	"net/http"
	"time"
)

func ReTime(pub string) string {
	t1, _ := time.Parse(time.RFC1123, pub)
	return time.Unix(t1.Unix(), 0).Format("2006-01-02 15:04:05")
}

func CheckCode(uri string) int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Errorf("请求发生异常%v", err.Error())
		return 404
	}
	response, resErr := client.Do(req)
	if resErr != nil {
		log.Errorf("发生异常%v", resErr)
		return 404
	}
	defer func(Body io.ReadCloser) {
		closeErr := Body.Close()
		if closeErr != nil {
			log.Errorf("关闭链接时发生异常:%v", closeErr.Error())
		}
	}(response.Body)
	return response.StatusCode
}

func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.0) * 1e-2
}
