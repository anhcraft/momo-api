package outbound

type MoneyTransferConfirmBody struct {
	AppCode        string                            `json:"appCode"`
	AppId          string                            `json:"appId"`
	AppVer         int                               `json:"appVer"`
	BuildNumber    int                               `json:"buildNumber"`
	Channel        string                            `json:"channel"`
	CmdId          string                            `json:"cmdId"`
	DeviceOs       string                            `json:"deviceOS"`
	Extra          MoneyTransferConfirmExtra         `json:"extra"`
	Lang           string                            `json:"lang"`
	MomoMsg        MoneyTransferConfirmMsg           `json:"momoMsg"`
	MsgType        string                            `json:"msgType"`
	ConfirmMsgType string                            `json:"confirmMsgType"`
	ConfirmClass   string                            `json:"confirmClass"`
	Time           int64                             `json:"time"`
	User           string                            `json:"user"`
	PaymentInfo    []MoneyTransferConfirmPaymentInfo `json:"paymentInfo"`
}

type MoneyTransferConfirmPaymentInfo struct {
	Title  string      `json:"title"`
	Value  interface{} `json:"value"`
	Format string      `json:"format"`
}

type MoneyTransferConfirmExtra struct {
	CheckSum string `json:"checkSum"`
}

type MoneyTransferConfirmMsg struct {
	Class                  string                      `json:"_class"`
	ServiceId              string                      `json:"serviceId"`
	ServiceCode            string                      `json:"serviceCode"`
	ClientTime             int64                       `json:"clientTime"`
	TranType               int                         `json:"tranType"`
	Comment                string                      `json:"comment"`
	Ref                    string                      `json:"ref"`
	PartnerId              string                      `json:"partnerId"`
	Amount                 float64                     `json:"amount"`
	Extras                 string                      `json:"extras"`
	Id                     string                      `json:"id"`
	Ids                    []string                    `json:"ids"`
	TranList               []MoneyTransferConfirmTrans `json:"tranList"`
	MoneySource            int                         `json:"moneySource"`
	TotalAmount            string                      `json:"totalAmount"`
	OriginalAmount         string                      `json:"originalAmount"`
	OriginalClass          string                      `json:"originalClass"`
	OriginalPhone          string                      `json:"originalPhone"`
	TotalFee               string                      `json:"totalFee"`
	CheckSum               string                      `json:"checkSum"`
	GetUserInfoTaskRequest string                      `json:"GetUserInfoTaskRequest"`
}

type MoneyTransferConfirmTrans struct {
	PartnerName    string  `json:"partnerName"`
	PartnerId      string  `json:"partnerId"`
	OriginalAmount float64 `json:"originalAmount"`
	TransferSource string  `json:"transferSource"`
	MoneySource    int     `json:"moneySource"`
}
