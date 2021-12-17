package img

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func SauceNAO(imgUrl string) []byte {
	params := url.Values{}
	Url, err := url.Parse("https://saucenao.com/search.php")
	if err != nil {
		log.Printf("err:%v", err)
	}
	params.Set("db", "999")
	params.Set("output_type", "2")
	params.Set("testmode", "1")
	params.Set("numres", "5")
	params.Set("url", imgUrl)
	params.Set("api_key", "f1fe0e2fd1d85dd206ee6a04df595160e5f4d323")
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
	fmt.Println(string(body))
	return body
}
