package outbound

// transferMoney, getTranHisMsgExtras
type MoneyTransferInitBody struct {
	AppCode        string                         `json:"appCode"`
	AppId          string                         `json:"appId"`
	AppVer         int                            `json:"appVer"`
	BuildNumber    int                            `json:"buildNumber"`
	Channel        string                         `json:"channel"`
	CmdId          string                         `json:"cmdId"`
	DeviceOs       string                         `json:"deviceOS"`
	Extra          MoneyTransferInitExtra         `json:"extra"`
	Lang           string                         `json:"lang"`
	MomoMsg        MoneyTransferInitMsg           `json:"momoMsg"`
	MsgType        string                         `json:"msgType"`
	ConfirmMsgType string                         `json:"confirmMsgType"`
	ConfirmClass   string                         `json:"confirmClass"`
	Time           int64                          `json:"time"`
	User           string                         `json:"user"`
	PaymentInfo    []MoneyTransferInitPaymentInfo `json:"paymentInfo"`
}

type MoneyTransferInitPaymentInfo struct {
	Title  string      `json:"title"`
	Value  interface{} `json:"value"`
	Format string      `json:"format"`
}

type MoneyTransferInitExtra struct {
	CheckSum string `json:"checkSum"`
}

type MoneyTransferInitMsg struct {
	Class       string                   `json:"_class"`
	ServiceId   string                   `json:"serviceId"`
	ServiceCode string                   `json:"serviceCode"`
	ClientTime  int64                    `json:"clientTime"`
	TranType    int                      `json:"tranType"`
	Comment     string                   `json:"comment"`
	Ref         string                   `json:"ref"`
	PartnerId   string                   `json:"partnerId"`
	Amount      float64                  `json:"amount"`
	Extras      string                   `json:"extras"`
	TranList    []MoneyTransferInitTrans `json:"tranList"`
	MoneySource int                      `json:"moneySource"`
}

type MoneyTransferInitTrans struct {
	PartnerName    string  `json:"partnerName"`
	PartnerId      string  `json:"partnerId"`
	OriginalAmount float64 `json:"originalAmount"`
	TransferSource string  `json:"transferSource"`
	MoneySource    int     `json:"moneySource"`
}
