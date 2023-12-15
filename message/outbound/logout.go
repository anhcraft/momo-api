package outbound

type LogoutBody struct {
	AppCode     string    `json:"appCode"`
	AppId       string    `json:"appId"`
	AppVer      int       `json:"appVer"`
	BuildNumber int       `json:"buildNumber"`
	Channel     string    `json:"channel"`
	CmdId       string    `json:"cmdId"`
	DeviceOs    string    `json:"deviceOS"`
	Lang        string    `json:"lang"`
	MomoMsg     LogoutMsg `json:"momoMsg"`
	MsgType     string    `json:"msgType"`
	Time        int64     `json:"time"`
	User        string    `json:"user"`
}

type LogoutMsg struct {
	Class string `json:"_class"`
}
