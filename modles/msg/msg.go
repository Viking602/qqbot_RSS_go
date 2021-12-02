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
