package outbound

import "momo-api/message/common"

type RequestOTPBody struct {
	AppCode     string              `json:"appCode"`
	AppId       string              `json:"appId"`
	AppVer      int                 `json:"appVer"`
	BuildNumber int                 `json:"buildNumber"`
	Channel     string              `json:"channel"`
	CmdId       string              `json:"cmdId"`
	DeviceOs    string              `json:"deviceOS"`
	Extra       RequestOTPExtra     `json:"extra"`
	Lang        string              `json:"lang"`
	MomoMsg     common.RegDeviceMsg `json:"momoMsg"`
	MsgType     string              `json:"msgType"`
	Time        int64               `json:"time"`
	User        string              `json:"user"`
}

type RequestOTPExtra struct {
	DeviceToken          string `json:"DEVICE_TOKEN"`
	Idfa                 string `json:"IDFA"`
	ModelId              string `json:"MODELID"`
	OneSignalToken       string `json:"ONESIGNAL_TOKEN"`
	RequireHashStringOtp bool   `json:"REQUIRE_HASH_STRING_OTP"`
	SecureId             string `json:"SECUREID"`
	Simulator            bool   `json:"SIMULATOR"`
	Token                string `json:"TOKEN"`
	Action               string `json:"action"`
	IsVoice              bool   `json:"isVoice"`
	RKey                 string `json:"rkey"`
}
