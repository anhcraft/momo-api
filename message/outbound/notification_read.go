package outbound

type NotificationReadBody struct {
	ID     string `json:"Id"`
	IsRead bool   `json:"isRead"`
}
