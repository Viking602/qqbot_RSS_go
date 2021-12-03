package msg

type Lifecycle struct {
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfId        int64  `json:"self_id"`
	SubType       string `json:"sub_type"`
	Time          int    `json:"time"`
}

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

type ResultMsg struct {
	Data struct {
		MessageId int `json:"message_id"`
	} `json:"data"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

type Msg struct {
	Data struct {
		Group       bool   `json:"group"`
		GroupId     int    `json:"group_id"`
		Message     string `json:"message"`
		MessageId   int    `json:"message_id"`
		MessageIdV2 string `json:"message_id_v2"`
		MessageSeq  int    `json:"message_seq"`
		MessageType string `json:"message_type"`
		RealId      int    `json:"real_id"`
		Sender      struct {
			Nickname string `json:"nickname"`
			UserId   int64  `json:"user_id"`
		} `json:"sender"`
		Time int `json:"time"`
	} `json:"data"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

type BiliLiveInfo struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Uid              int      `json:"uid"`
		RoomId           int      `json:"room_id"`
		ShortId          int      `json:"short_id"`
		Attention        int      `json:"attention"`
		Online           int      `json:"online"`
		IsPortrait       bool     `json:"is_portrait"`
		Description      string   `json:"description"`
		LiveStatus       int      `json:"live_status"`
		AreaId           int      `json:"area_id"`
		ParentAreaId     int      `json:"parent_area_id"`
		ParentAreaName   string   `json:"parent_area_name"`
		OldAreaId        int      `json:"old_area_id"`
		Background       string   `json:"background"`
		Title            string   `json:"title"`
		UserCover        string   `json:"user_cover"`
		Keyframe         string   `json:"keyframe"`
		IsStrictRoom     bool     `json:"is_strict_room"`
		LiveTime         string   `json:"live_time"`
		Tags             string   `json:"tags"`
		IsAnchor         int      `json:"is_anchor"`
		RoomSilentType   string   `json:"room_silent_type"`
		RoomSilentLevel  int      `json:"room_silent_level"`
		RoomSilentSecond int      `json:"room_silent_second"`
		AreaName         string   `json:"area_name"`
		Pendants         string   `json:"pendants"`
		AreaPendants     string   `json:"area_pendants"`
		HotWords         []string `json:"hot_words"`
		HotWordsStatus   int      `json:"hot_words_status"`
		Verify           string   `json:"verify"`
		NewPendants      struct {
			Frame struct {
				Name       string `json:"name"`
				Value      string `json:"value"`
				Position   int    `json:"position"`
				Desc       string `json:"desc"`
				Area       int    `json:"area"`
				AreaOld    int    `json:"area_old"`
				BgColor    string `json:"bg_color"`
				BgPic      string `json:"bg_pic"`
				UseOldArea bool   `json:"use_old_area"`
			} `json:"frame"`
			Badge struct {
				Name     string `json:"name"`
				Position int    `json:"position"`
				Value    string `json:"value"`
				Desc     string `json:"desc"`
			} `json:"badge"`
			MobileFrame struct {
				Name       string `json:"name"`
				Value      string `json:"value"`
				Position   int    `json:"position"`
				Desc       string `json:"desc"`
				Area       int    `json:"area"`
				AreaOld    int    `json:"area_old"`
				BgColor    string `json:"bg_color"`
				BgPic      string `json:"bg_pic"`
				UseOldArea bool   `json:"use_old_area"`
			} `json:"mobile_frame"`
			MobileBadge interface{} `json:"mobile_badge"`
		} `json:"new_pendants"`
		UpSession            string `json:"up_session"`
		PkStatus             int    `json:"pk_status"`
		PkId                 int    `json:"pk_id"`
		BattleId             int    `json:"battle_id"`
		AllowChangeAreaTime  int    `json:"allow_change_area_time"`
		AllowUploadCoverTime int    `json:"allow_upload_cover_time"`
		StudioInfo           struct {
			Status     int           `json:"status"`
			MasterList []interface{} `json:"master_list"`
		} `json:"studio_info"`
	} `json:"data"`
}

type UpInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Mid       int    `json:"mid"`
		Name      string `json:"name"`
		Sex       string `json:"sex"`
		Face      string `json:"face"`
		FaceNft   int    `json:"face_nft"`
		Sign      string `json:"sign"`
		Rank      int    `json:"rank"`
		Level     int    `json:"level"`
		Jointime  int    `json:"jointime"`
		Moral     int    `json:"moral"`
		Silence   int    `json:"silence"`
		Coins     int    `json:"coins"`
		FansBadge bool   `json:"fans_badge"`
		FansMedal struct {
			Show  bool `json:"show"`
			Wear  bool `json:"wear"`
			Medal struct {
				Uid              int    `json:"uid"`
				TargetId         int    `json:"target_id"`
				MedalId          int    `json:"medal_id"`
				Level            int    `json:"level"`
				MedalName        string `json:"medal_name"`
				MedalColor       int    `json:"medal_color"`
				Intimacy         int    `json:"intimacy"`
				NextIntimacy     int    `json:"next_intimacy"`
				DayLimit         int    `json:"day_limit"`
				MedalColorStart  int    `json:"medal_color_start"`
				MedalColorEnd    int    `json:"medal_color_end"`
				MedalColorBorder int    `json:"medal_color_border"`
				IsLighted        int    `json:"is_lighted"`
				LightStatus      int    `json:"light_status"`
				WearingStatus    int    `json:"wearing_status"`
				Score            int    `json:"score"`
			} `json:"medal"`
		} `json:"fans_medal"`
		Official struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"official"`
		Vip struct {
			Type       int   `json:"type"`
			Status     int   `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int   `json:"vip_pay_type"`
			ThemeType  int   `json:"theme_type"`
			Label      struct {
				Path        string `json:"path"`
				Text        string `json:"text"`
				LabelTheme  string `json:"label_theme"`
				TextColor   string `json:"text_color"`
				BgStyle     int    `json:"bg_style"`
				BgColor     string `json:"bg_color"`
				BorderColor string `json:"border_color"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptUrl string `json:"avatar_subscript_url"`
		} `json:"vip"`
		Pendant struct {
			Pid               int    `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int    `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
		} `json:"pendant"`
		Nameplate struct {
			Nid        int    `json:"nid"`
			Name       string `json:"name"`
			Image      string `json:"image"`
			ImageSmall string `json:"image_small"`
			Level      string `json:"level"`
			Condition  string `json:"condition"`
		} `json:"nameplate"`
		UserHonourInfo struct {
			Mid    int           `json:"mid"`
			Colour interface{}   `json:"colour"`
			Tags   []interface{} `json:"tags"`
		} `json:"user_honour_info"`
		IsFollowed bool   `json:"is_followed"`
		TopPhoto   string `json:"top_photo"`
		Theme      struct {
		} `json:"theme"`
		SysNotice struct {
		} `json:"sys_notice"`
		LiveRoom struct {
			RoomStatus    int    `json:"roomStatus"`
			LiveStatus    int    `json:"liveStatus"`
			Url           string `json:"url"`
			Title         string `json:"title"`
			Cover         string `json:"cover"`
			Online        int    `json:"online"`
			Roomid        int    `json:"roomid"`
			RoundStatus   int    `json:"roundStatus"`
			BroadcastType int    `json:"broadcast_type"`
		} `json:"live_room"`
		Birthday   string      `json:"birthday"`
		School     interface{} `json:"school"`
		Profession struct {
			Name string `json:"name"`
		} `json:"profession"`
		Tags   []string `json:"tags"`
		Series struct {
			UserUpgradeStatus int  `json:"user_upgrade_status"`
			ShowUpgradeWindow bool `json:"show_upgrade_window"`
		} `json:"series"`
	} `json:"data"`
}
