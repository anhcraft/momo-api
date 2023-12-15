package inbound

import "encoding/json"

type NotificationListResponse struct {
	Message *NotificationListResponseMessage `json:"message"`
}

func (r *NotificationListResponse) Success() bool {
	return r.Message != nil && r.Message.ResponseInfo != nil && r.Message.ResponseInfo.ErrorCode == 0
}

func ParseNotificationListResponse(body []byte) *NotificationListResponse {
	data := &NotificationListResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}

type NotificationListResponseMessage struct {
	Data         *NotificationListResponseMessageData `json:"data"`
	ResponseInfo *GenericResponse                     `json:"responseInfo"`
}

type NotificationListResponseMessageData struct {
	Notifications []*NotificationSchema `json:"notifications"`
	Cursor        string                `json:"cursor"`
}

type NotificationSchema struct {
	Id                string  `json:"Id"`
	AllowOutApp       bool    `json:"allowOutApp"`
	Body              string  `json:"body"`
	BodyIOS           string  `json:"bodyIOS"`
	BtnStatus         float64 `json:"btnStatus"`
	BtnTitle          string  `json:"btnTitle"`
	Caption           string  `json:"caption"`
	Category          string  `json:"category"`
	CmdId             string  `json:"cmdId"`
	Extra             string  `json:"extra"`
	HtmlBody          string  `json:"htmlBody"`
	IsDeleted         bool    `json:"isDeleted"`
	IsOffPopup        bool    `json:"isOffPopup"`
	IsRead            bool    `json:"isRead"`
	IsTranSuccess     bool    `json:"isTranSuccess"`
	OS                string  `json:"os"`
	Prefix            string  `json:"prefix"`
	QOS               float64 `json:"qos"`
	ReceiverNumber    string  `json:"receiverNumber"`
	RefId             string  `json:"refId"`
	Sender            string  `json:"sender"`
	Sms               string  `json:"sms"`
	Status            float64 `json:"status"`
	Time              float64 `json:"time"`
	TimeString        string  `json:"timeString"`
	Token             string  `json:"token"`
	TranId            float64 `json:"tranId"`
	TransHisServiceId string  `json:"transHisServiceId"`
	Type              float64 `json:"type"`
	UserConfig        string  `json:"userConfig"`
}
