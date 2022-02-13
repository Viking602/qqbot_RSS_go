/*
 * Copyright 2021 Viking602
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package hibiapi

import (
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiUrl = "https://api.obfs.dev"

func Search163(s string) []byte {
	params := url.Values{}
	path := "/api/netease/search"
	Url, err := url.Parse(apiUrl + path)
	if err != nil {
		log.Errorf("发生异常:%v", err.Error())
	}
	params.Set("s", s)
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
