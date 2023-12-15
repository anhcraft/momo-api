package outbound

import "momo-api/message/common"

type VerifyPhoneBody struct {
	AppCode     string              `json:"appCode"`
	AppId       string              `json:"appId"`
	AppVer      int                 `json:"appVer"`
	BuildNumber int                 `json:"buildNumber"`
	Channel     string              `json:"channel"`
	CmdId       string              `json:"cmdId"`
	DeviceOs    string              `json:"deviceOS"`
	Lang        string              `json:"lang"`
	MomoMsg     common.RegDeviceMsg `json:"momoMsg"`
	MsgType     string              `json:"msgType"`
	Time        int64               `json:"time"`
	User        string              `json:"user"`
}
