package msg

type VideInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizId   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State     int `json:"state"`
		Duration  int `json:"duration"`
		MissionId int `json:"mission_id"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			ArgueMsg   string `json:"argue_msg"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		SeasonId int  `json:"season_id"`
		NoCache  bool `json:"no_cache"`
		Pages    []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			FirstFrame string `json:"first_frame"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
		Label struct {
			Type int `json:"type"`
		} `json:"label"`
		UgcSeason struct {
			Id        int    `json:"id"`
			Title     string `json:"title"`
			Cover     string `json:"cover"`
			Mid       int    `json:"mid"`
			Intro     string `json:"intro"`
			SignState int    `json:"sign_state"`
			Attribute int    `json:"attribute"`
			Sections  []struct {
				SeasonId int    `json:"season_id"`
				Id       int    `json:"id"`
				Title    string `json:"title"`
				Type     int    `json:"type"`
				Episodes []struct {
					SeasonId  int    `json:"season_id"`
					SectionId int    `json:"section_id"`
					Id        int    `json:"id"`
					Aid       int    `json:"aid"`
					Cid       int    `json:"cid"`
					Title     string `json:"title"`
					Attribute int    `json:"attribute"`
					Arc       struct {
						Aid       int    `json:"aid"`
						Videos    int    `json:"videos"`
						TypeId    int    `json:"type_id"`
						TypeName  string `json:"type_name"`
						Copyright int    `json:"copyright"`
						Pic       string `json:"pic"`
						Title     string `json:"title"`
						Pubdate   int    `json:"pubdate"`
						Ctime     int    `json:"ctime"`
						Desc      string `json:"desc"`
						State     int    `json:"state"`
						Duration  int    `json:"duration"`
						Rights    struct {
							Bp            int `json:"bp"`
							Elec          int `json:"elec"`
							Download      int `json:"download"`
							Movie         int `json:"movie"`
							Pay           int `json:"pay"`
							Hd5           int `json:"hd5"`
							NoReprint     int `json:"no_reprint"`
							Autoplay      int `json:"autoplay"`
							UgcPay        int `json:"ugc_pay"`
							IsCooperation int `json:"is_cooperation"`
							UgcPayPreview int `json:"ugc_pay_preview"`
						} `json:"rights"`
						Author struct {
							Mid  int    `json:"mid"`
							Name string `json:"name"`
							Face string `json:"face"`
						} `json:"author"`
						Stat struct {
							Aid        int    `json:"aid"`
							View       int    `json:"view"`
							Danmaku    int    `json:"danmaku"`
							Reply      int    `json:"reply"`
							Fav        int    `json:"fav"`
							Coin       int    `json:"coin"`
							Share      int    `json:"share"`
							NowRank    int    `json:"now_rank"`
							HisRank    int    `json:"his_rank"`
							Like       int    `json:"like"`
							Dislike    int    `json:"dislike"`
							Evaluation string `json:"evaluation"`
							ArgueMsg   string `json:"argue_msg"`
						} `json:"stat"`
						Dynamic   string `json:"dynamic"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
						DescV2 interface{} `json:"desc_v2"`
					} `json:"arc"`
					Page struct {
						Cid       int    `json:"cid"`
						Page      int    `json:"page"`
						From      string `json:"from"`
						Part      string `json:"part"`
						Duration  int    `json:"duration"`
						Vid       string `json:"vid"`
						Weblink   string `json:"weblink"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
					} `json:"page"`
					Bvid string `json:"bvid"`
				} `json:"episodes"`
			} `json:"sections"`
			Stat struct {
				SeasonId int `json:"season_id"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Fav      int `json:"fav"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
			} `json:"stat"`
			EpCount    int `json:"ep_count"`
			SeasonType int `json:"season_type"`
		} `json:"ugc_season"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			UrlImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
			Honor []struct {
				Aid                int    `json:"aid"`
				Type               int    `json:"type"`
				Desc               string `json:"desc"`
				WeeklyRecommendNum int    `json:"weekly_recommend_num"`
			} `json:"honor"`
		} `json:"honor_reply"`
	} `json:"data"`
}
