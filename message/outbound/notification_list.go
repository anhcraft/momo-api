package outbound

type NotificationListBody struct {
	Cursor   string `json:"cursor"`
	FromTime int64  `json:"fromTime"`
	Limit    int    `json:"limit"`
	ToTime   int64  `json:"toTime"`
	UserId   string `json:"userId"`
}
