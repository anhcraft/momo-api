package inbound

import "encoding/json"

type TransactionDetailResponse struct {
	Time       int64                     `json:"time"`
	StatusCode int                       `json:"statusCode"`
	ResultCode int                       `json:"resultCode"`
	Message    string                    `json:"message"`
	MomoMsg    *TransactionDetailMomoMsg `json:"momoMsg"`
}

func (r *TransactionDetailResponse) Success() bool {
	return r.StatusCode == 200
}

func ParseTransactionDetailResponse(body []byte) *TransactionDetailResponse {
	data := &TransactionDetailResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}

type TransactionDetailMomoMsg struct {
	ID                    string      `json:"id"`       // unknown id
	Username              string      `json:"username"` // user id, user phone
	TransID               int64       `json:"transId"`
	ServiceID             string      `json:"serviceId"` // see trans browse for more info
	LastUpdate            int64       `json:"lastUpdate"`
	TransCategory         string      `json:"transCategory"` // see trans browse for more info
	Status                int         `json:"status"`        // see trans browse for more info
	BillID                string      `json:"billId"`
	Quantity              int         `json:"quantity"`
	TotalOriginalAmount   int         `json:"totalOriginalAmount"`
	TotalAmount           int         `json:"totalAmount"`
	UseVoucher            int         `json:"useVoucher"`
	UsePoint              interface{} `json:"usePoint"`
	TotalFee              int         `json:"totalFee"`
	SourceID              string      `json:"sourceId"`   // see trans browse for more info
	SourceName            string      `json:"sourceName"` // see trans browse for more info
	TargetID              string      `json:"targetId"`   // see trans browse for more info
	TargetName            string      `json:"targetName"` // see trans browse for more info
	UsePrepaid            int         `json:"usePrepaid"`
	Io                    int         `json:"io"`
	MoneySource           int         `json:"moneySource"`
	PostBalance           int         `json:"postBalance"`
	TotalSubTotalAmount   interface{} `json:"totalSubTotalAmount"`
	TotalFeeSof           int         `json:"totalFeeSof"`
	TotalDiscountAmount   int         `json:"totalDiscountAmount"`
	TotalDiscountFee      interface{} `json:"totalDiscountFee"`
	MoneySourceName       string      `json:"moneySourceName"`
	PurchaseID            string      `json:"purchaseId"`
	CashbackIncr          int         `json:"cashbackIncr"`
	PointIncr             int         `json:"pointIncr"`
	BonusIncr             int         `json:"bonusIncr"`
	PrepaidIds            interface{} `json:"prepaidIds"`
	GiftPrepaidID         interface{} `json:"giftPrepaidId"`
	ServiceName           string      `json:"serviceName"`
	ErrorType             interface{} `json:"errorType"`
	ErrorCode             int         `json:"errorCode"`
	ErrorDesc             string      `json:"errorDesc"`
	TranshisData          string      `json:"transhisData"`
	DiscountAmount        int         `json:"discountAmount"`
	CreatedAt             int64       `json:"createdAt"`
	DirectDiscount        int         `json:"directDiscount"`
	HasOldData            bool        `json:"hasOldData"`
	RefundTransID         string      `json:"refundTransId"`
	RefundAmount          interface{} `json:"refundAmount"`
	RefundMoneySourceName interface{} `json:"refundMoneySourceName"`
	PartnerAvatarURL      interface{} `json:"partnerAvatarUrl"`
	CbAmountLater         int         `json:"cbAmountLater"`
	OldData               interface{} `json:"oldData"`
	Channel               string      `json:"channel"`
	CallSpecificResult    bool        `json:"callSpecificResult"`
	DefaultTransData      interface{} `json:"defaultTransData"`
	ParentID              interface{} `json:"parentId"`
	InsTenors             int         `json:"insTenors"`
	ServiceData           string      `json:"serviceData"` // stringified TransactionDetailP2PTransServiceData
	IsDeleted             bool        `json:"isDeleted"`
}

type TransactionDetailP2PTransServiceData struct {
	TransID             int64       `json:"TRANS_ID"`
	OwnerNumber         string      `json:"OWNER_NUMBER"`
	ReceiverNumber      string      `json:"RECEIVER_NUMBER"`
	ReceiverType        int         `json:"RECEIVER_TYPE"`
	SocialUserID        interface{} `json:"SOCIAL_USER_ID"`
	ChatID              interface{} `json:"CHAT_ID"`
	TransferSource      interface{} `json:"TRANSFER_SOURCE"`
	Stickers            interface{} `json:"STICKERS"`
	MoneyRequestID      interface{} `json:"MONEY_REQUEST_ID"`
	MoneyRequestIds     interface{} `json:"MONEY_REQUEST_IDS"`
	MoneyRequestGroupID interface{} `json:"MONEY_REQUEST_GROUP_ID"`
	ThemeURL            string      `json:"THEME_URL"`
	BillID              interface{} `json:"BILL_ID"`
	OrderInfo           interface{} `json:"ORDER_INFO"`
	ShippingInfo        interface{} `json:"SHIPPING_INFO"`
	LastUpdate          int64       `json:"LAST_UPDATE"`
	CommentValue        string      `json:"COMMENT_VALUE"`
	PartnerAction       interface{} `json:"PARTNER_ACTION"`
	OrderID             interface{} `json:"ORDER_ID"`
	LinkID              interface{} `json:"LINK_ID"`
	GiftType            interface{} `json:"GIFT_TYPE"`
	ServiceID           string      `json:"SERVICE_ID"`
	EventType           interface{} `json:"EVENT_TYPE"`
	Amount              interface{} `json:"AMOUNT"`
	AutoInvestment      interface{} `json:"AUTO_INVESTMENT"`
	TransHisServiceID   string      `json:"transHisServiceId"`
	ServiceID0          string      `json:"serviceId"`
}
