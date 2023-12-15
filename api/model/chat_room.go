package model

type ChatRoom struct {
	Id            string   `json:"id,omitempty"`
	Source        string   `json:"source,omitempty"` // SOCIAL, ESCROW, OA
	Type          string   `json:"type,omitempty"`   // NORMAL_ROOM, REQUEST_ROOM, MONEY_POOL
	IsGroup       bool     `json:"isGroup,omitempty"`
	IsBlock       bool     `json:"isBlock,omitempty"`
	IsHide        bool     `json:"isHide,omitempty"`
	Members       []string `json:"members,omitempty"`
	Creator       string   `json:"creator,omitempty"` // User-Id
	Unread        bool     `json:"unread,omitempty"`
	LastMessageAt int64    `json:"lastMessageAt,omitempty"`
	LastMessageID string   `json:"lastMessageID,omitempty"`
	IsMute        bool     `json:"isMute,omitempty"`
	LastMessage   string   `json:"lastMessage,omitempty"`
}
