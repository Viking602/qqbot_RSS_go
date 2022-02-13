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

package msg

type Search struct {
	NeedLogin bool `json:"needLogin"`
	Result    struct {
		Songs []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
			Pst  int    `json:"pst"`
			T    int    `json:"t"`
			Ar   []struct {
				Id    int      `json:"id"`
				Name  string   `json:"name"`
				Tns   []string `json:"tns"`
				Alias []string `json:"alias"`
				Alia  []string `json:"alia,omitempty"`
			} `json:"ar"`
			Alia []string    `json:"alia"`
			Pop  float64     `json:"pop"`
			St   int         `json:"st"`
			Rt   *string     `json:"rt"`
			Fee  int         `json:"fee"`
			V    int         `json:"v"`
			Crbt interface{} `json:"crbt"`
			Cf   string      `json:"cf"`
			Al   struct {
				Id     int      `json:"id"`
				Name   string   `json:"name"`
				PicUrl string   `json:"picUrl"`
				Tns    []string `json:"tns"`
				Pic    int64    `json:"pic"`
				PicStr string   `json:"pic_str,omitempty"`
			} `json:"al"`
			Dt int `json:"dt"`
			H  struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"h"`
			M struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"m"`
			L struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"l"`
			A                    interface{}   `json:"a"`
			Cd                   string        `json:"cd"`
			No                   int           `json:"no"`
			RtUrl                interface{}   `json:"rtUrl"`
			Ftype                int           `json:"ftype"`
			RtUrls               []interface{} `json:"rtUrls"`
			DjId                 int           `json:"djId"`
			Copyright            int           `json:"copyright"`
			SId                  int           `json:"s_id"`
			Mark                 int64         `json:"mark"`
			OriginCoverType      int           `json:"originCoverType"`
			OriginSongSimpleData *struct {
				SongId  int    `json:"songId"`
				Name    string `json:"name"`
				Artists []struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"artists"`
				AlbumMeta struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"albumMeta"`
			} `json:"originSongSimpleData"`
			ResourceState   bool        `json:"resourceState"`
			Version         int         `json:"version"`
			Single          int         `json:"single"`
			NoCopyrightRcmd interface{} `json:"noCopyrightRcmd"`
			Mst             int         `json:"mst"`
			Cp              int         `json:"cp"`
			Mv              int         `json:"mv"`
			Rtype           int         `json:"rtype"`
			Rurl            interface{} `json:"rurl"`
			PublishTime     int64       `json:"publishTime"`
			Privilege       struct {
				Id                 int         `json:"id"`
				Fee                int         `json:"fee"`
				Payed              int         `json:"payed"`
				St                 int         `json:"st"`
				Pl                 int         `json:"pl"`
				Dl                 int         `json:"dl"`
				Sp                 int         `json:"sp"`
				Cp                 int         `json:"cp"`
				Subp               int         `json:"subp"`
				Cs                 bool        `json:"cs"`
				Maxbr              int         `json:"maxbr"`
				Fl                 int         `json:"fl"`
				Toast              bool        `json:"toast"`
				Flag               int         `json:"flag"`
				PreSell            bool        `json:"preSell"`
				PlayMaxbr          int         `json:"playMaxbr"`
				DownloadMaxbr      int         `json:"downloadMaxbr"`
				Rscl               interface{} `json:"rscl"`
				FreeTrialPrivilege struct {
					ResConsumable  bool `json:"resConsumable"`
					UserConsumable bool `json:"userConsumable"`
				} `json:"freeTrialPrivilege"`
				ChargeInfoList []struct {
					Rate          int         `json:"rate"`
					ChargeUrl     interface{} `json:"chargeUrl"`
					ChargeMessage interface{} `json:"chargeMessage"`
					ChargeType    int         `json:"chargeType"`
				} `json:"chargeInfoList"`
			} `json:"privilege"`
			Tns []string `json:"tns,omitempty"`
		} `json:"songs"`
		SongCount int `json:"songCount"`
	} `json:"result"`
	Code int `json:"code"`
}
