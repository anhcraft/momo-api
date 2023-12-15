package outbound

import "momo-api/message/common"

type VerifyOtpBody struct {
	AppCode     string              `json:"appCode"`
	AppId       string              `json:"appId"`
	AppVer      int                 `json:"appVer"`
	BuildNumber int                 `json:"buildNumber"`
	Channel     string              `json:"channel"`
	CmdId       string              `json:"cmdId"`
	DeviceOs    string              `json:"deviceOS"`
	Extra       VerifyOTPExtra      `json:"extra"`
	Lang        string              `json:"lang"`
	MomoMsg     common.RegDeviceMsg `json:"momoMsg"`
	MsgType     string              `json:"msgType"`
	Time        int64               `json:"time"`
	User        string              `json:"user"`
}

type VerifyOTPExtra struct {
	DeviceToken    string `json:"DEVICE_TOKEN"`
	Idfa           string `json:"IDFA"`
	ModelId        string `json:"MODELID"`
	OneSignalToken string `json:"ONESIGNAL_TOKEN"`
	SecureId       string `json:"SECUREID"`
	Simulator      bool   `json:"SIMULATOR"`
	Token          string `json:"TOKEN"`
	OHash          string `json:"ohash"`
}
