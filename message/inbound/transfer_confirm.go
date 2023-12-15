package inbound

import (
	"encoding/json"
)

type MoneyTransferConfirmResponse struct {
	Time      int64  `json:"time"`
	User      string `json:"user"`
	Pass      string `json:"pass"`
	CmdID     string `json:"cmdId"`
	Lang      string `json:"lang"`
	MsgType   string `json:"msgType"`
	Result    bool   `json:"result"`
	ErrorCode int    `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
	AppCode   string `json:"appCode"`
	AppVer    int    `json:"appVer"`
	Channel   string `json:"channel"`
	DeviceOS  string `json:"deviceOS"`
	Path      string `json:"path"`
	Session   string `json:"session"`
}

func (r *MoneyTransferConfirmResponse) Success() bool {
	return r.ErrorCode == 0 || r.ErrorCode == 9000
}

func ParseMoneyTransferConfirmResponse(body []byte) *MoneyTransferConfirmResponse {
	data := &MoneyTransferConfirmResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}
