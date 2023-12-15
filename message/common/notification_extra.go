package common

type NotificationExtraTransaction struct {
	TranId       string  `json:"tranId"`
	Comment      string  `json:"comment"`
	PartnerId    string  `json:"partnerId"`
	PartnerName  string  `json:"partnerName"`
	Amount       float64 `json:"amount"`
	ReceiverType int     `json:"receiverType"`
}
