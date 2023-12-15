package inbound

import "encoding/json"

type TransactionBrowseResponse struct {
	Time         int64                       `json:"time"`
	StatusCode   int                         `json:"statusCode"`
	ResultCode   int                         `json:"resultCode"`
	Message      string                      `json:"message"`
	MomoMsg      []*TransactionBrowseMomoMsg `json:"momoMsg"`
	CurrentLimit int                         `json:"currentLimit"`
	ListOver     bool                        `json:"listOver"`
}

func (r *TransactionBrowseResponse) Success() bool {
	return r.StatusCode == 200
}

func ParseTransactionBrowseResponse(body []byte) *TransactionBrowseResponse {
	data := &TransactionBrowseResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}

// vn.momo.core.modules.storage.relational.schemes.TranHisNewSchema
type TransactionBrowseMomoMsg struct {
	ID                  string      `json:"id"`        // unknown ID
	Username            string      `json:"username"`  // Phone, User ID
	TransID             int64       `json:"transId"`   // Trans ID
	ServiceID           string      `json:"serviceId"` // transfer_masking, transfer_p2p, banklink_cashin, or 3rd apps (TIKI042017)
	LastUpdate          int64       `json:"lastUpdate"`
	TransCategory       string      `json:"transCategory"` // th_p2p_w2w_receive_cat3, th_receive_cat1, th_AppMoMo_cat2, th_cashin_nhlk_cat2, th_topupcard_viettel_cat3
	Status              int         `json:"status"`        // 2
	BillID              string      `json:"billId"`        // nullable
	Quantity            int         `json:"quantity"`
	TotalOriginalAmount int         `json:"totalOriginalAmount"`
	TotalAmount         int         `json:"totalAmount"`
	UseVoucher          int         `json:"useVoucher"`
	UsePoint            interface{} `json:"usePoint"`
	TotalFee            int         `json:"totalFee"`
	SourceID            string      `json:"sourceId"`   // Sender User ID or Bank ID (e.g "acb.bank")
	SourceName          string      `json:"sourceName"` // Sender Name or Bank Name
	TargetID            string      `json:"targetId"`   // Target User ID or App ID
	TargetName          string      `json:"targetName"` // may be == TargetId, or App Name, "vttiocta_mathe_vt"
	UsePrepaid          int         `json:"usePrepaid"`
	Io                  int         `json:"io"`
	MoneySource         int         `json:"moneySource"` // 1 (from MoMo)
	PostBalance         int         `json:"postBalance"` // new balance
	TotalSubTotalAmount interface{} `json:"totalSubTotalAmount"`
	TotalFeeSof         int         `json:"totalFeeSof"`
	TotalDiscountAmount int         `json:"totalDiscountAmount"`
	TotalDiscountFee    interface{} `json:"totalDiscountFee"`
	MoneySourceName     string      `json:"moneySourceName"` // "Ví momo", Bank Name
	PurchaseID          string      `json:"purchaseId"`
	CashbackIncr        int         `json:"cashbackIncr"`
	PointIncr           int         `json:"pointIncr"`
	BonusIncr           int         `json:"bonusIncr"`
	PrepaidIds          interface{} `json:"prepaidIds"`
	GiftPrepaidID       interface{} `json:"giftPrepaidId"`
	ServiceName         string      `json:"serviceName"` // User ID, "TIKI"
	ErrorType           interface{} `json:"errorType"`
	ErrorCode           int         `json:"errorCode"`
	ErrorDesc           string      `json:"errorDesc"`
	TranshisData        string      `json:"transhisData"`
	DiscountAmount      int         `json:"discountAmount"`
	CreatedAt           int64       `json:"createdAt"`
	DirectDiscount      int         `json:"directDiscount"`
	IsDeleted           bool        `json:"isDeleted"`
}
