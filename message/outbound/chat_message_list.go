package outbound

type ChatMessageListBody struct {
	RoomSource      string `json:"roomSource"`
	RoomID          string `json:"roomId"`
	UserID          string `json:"userId"`
	TargetMessageID string `json:"targetMessageId"`
	Limit           int    `json:"limit"`  // 30
	Action          int    `json:"action"` // 1
}
