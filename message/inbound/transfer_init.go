package inbound

import (
	"encoding/json"
)

type MoneyTransferInitResponse struct {
	MomoMsg   MoneyTransferInitResponseMsg   `json:"momoMsg"`
	Time      int64                          `json:"time"`
	User      string                         `json:"user"`
	Pass      string                         `json:"pass"`
	CmdID     string                         `json:"cmdId"`
	Lang      string                         `json:"lang"`
	MsgType   string                         `json:"msgType"`
	Result    bool                           `json:"result"`
	ErrorCode int                            `json:"errorCode"`
	ErrorDesc string                         `json:"errorDesc"`
	AppCode   string                         `json:"appCode"`
	AppVer    int                            `json:"appVer"`
	Channel   string                         `json:"channel"`
	DeviceOS  string                         `json:"deviceOS"`
	Path      string                         `json:"path"`
	Session   string                         `json:"session"`
	Extra     MoneyTransferInitResponseExtra `json:"extra"`
}

func (r *MoneyTransferInitResponse) Success() bool {
	return r.ErrorCode == 0
}

func ParseMoneyTransferInitResponse(body []byte) *MoneyTransferInitResponse {
	data := &MoneyTransferInitResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}

type MoneyTransferInitResponseMsg struct {
	BankInReply MoneyTransferInitBankInReply `json:"bankInReply"`
	ReplyMsgs   []MoneyTransferInitReplyMsgs `json:"replyMsgs"`
	IsLixi      bool                         `json:"isLixi"`
	Lixi        bool                         `json:"lixi"`
	Class       string                       `json:"_class"`
}

type MoneyTransferInitResponseExtra struct {
	TotalAmount            string `json:"totalAmount"`
	OriginalAmount         string `json:"originalAmount"`
	OriginalClass          string `json:"originalClass"`
	OriginalPhone          string `json:"originalPhone"`
	TotalFee               string `json:"totalFee"`
	CheckSum               string `json:"checkSum"`
	GetUserInfoTaskRequest string `json:"GetUserInfoTaskRequest"`
}

type MoneyTransferInitEnableOptions struct {
	Voucher  bool   `json:"voucher"`
	Discount bool   `json:"discount"`
	Prepaid  bool   `json:"prepaid"`
	Desc     string `json:"desc"`
}

type MoneyTransferInitBankRepTranHis struct {
	OtpType       string                         `json:"otpType"`
	IPAddress     string                         `json:"ipAddress"`
	EnableOptions MoneyTransferInitEnableOptions `json:"enableOptions"`
	Class         string                         `json:"_class"`
}

type MoneyTransferInitBankInReply struct {
	TranHisMsg MoneyTransferInitBankRepTranHis `json:"tranHisMsg"`
	Class      string                          `json:"_class"`
}

type MoneyTransferInitTranHisMsg struct {
	ID             string                         `json:"ID"`
	User           string                         `json:"user"`
	CommandInd     string                         `json:"commandInd"`
	TranID         float64                        `json:"tranId"`
	ClientTime     float64                        `json:"clientTime"`
	AckTime        float64                        `json:"ackTime"`
	TranType       int                            `json:"tranType"`
	Io             int                            `json:"io"`
	PartnerID      string                         `json:"partnerId"`
	PartnerName    string                         `json:"partnerName"`
	Amount         float64                        `json:"amount"`
	Status         int                            `json:"status"`
	OwnerNumber    string                         `json:"ownerNumber"`
	MoneySource    int                            `json:"moneySource"`
	Desc           string                         `json:"desc"`
	OriginalAmount float64                        `json:"originalAmount"`
	Quantity       int                            `json:"quantity"`
	LastUpdate     float64                        `json:"lastUpdate"`
	Share          string                         `json:"share"`
	ReceiverType   int                            `json:"receiverType"`
	Extras         string                         `json:"extras"`
	Channel        string                         `json:"channel"`
	OtpType        string                         `json:"otpType"`
	TransferSource string                         `json:"transferSource"`
	IPAddress      string                         `json:"ipAddress"`
	EnableOptions  MoneyTransferInitEnableOptions `json:"enableOptions"`
	Class          string                         `json:"_class"`
}

type MoneyTransferInitReplyMsgs struct {
	ID         string                      `json:"ID"`
	TransID    float64                     `json:"transId"`
	IsSucess   bool                        `json:"isSucess"`
	TranHisMsg MoneyTransferInitTranHisMsg `json:"tranHisMsg"`
	ID0        string                      `json:"id"`
	Class      string                      `json:"_class"`
}
