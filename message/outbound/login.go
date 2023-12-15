package outbound

type LoginBody struct {
	AppCode     string     `json:"appCode"`
	AppId       string     `json:"appId"`
	AppVer      int        `json:"appVer"`
	BuildNumber int        `json:"buildNumber"`
	Channel     string     `json:"channel"`
	CmdId       string     `json:"cmdId"`
	DeviceOs    string     `json:"deviceOS"`
	Extra       LoginExtra `json:"extra"`
	Lang        string     `json:"lang"`
	MomoMsg     LoginMsg   `json:"momoMsg"`
	MsgType     string     `json:"msgType"`
	Pass        string     `json:"pass"`
	Time        int64      `json:"time"`
	User        string     `json:"user"`
}

type LoginExtra struct {
	DeviceToken    string `json:"DEVICE_TOKEN"`
	Idfa           string `json:"IDFA"`
	ModelId        string `json:"MODELID"`
	OneSignalToken string `json:"ONESIGNAL_TOKEN"`
	SecureId       string `json:"SECUREID"`
	Simulator      bool   `json:"SIMULATOR"`
	Token          string `json:"TOKEN"`
	CheckSum       string `json:"checkSum"`
	PHash          string `json:"pHash"`
}

type LoginMsg struct {
	Class   string `json:"_class"`
	IsSetup bool   `json:"isSetup"`
}
