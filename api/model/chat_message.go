package model

type ChatMessage struct {
	Id          string                  `json:"id,omitempty"`
	Sender      string                  `json:"sender,omitempty"`
	Room        string                  `json:"room,omitempty"`
	Source      string                  `json:"source,omitempty"` // SOCIAL
	RequestId   string                  `json:"requestId,omitempty"`
	MsgType     string                  `json:"msgType,omitempty"`    // TEXT, INFORMATIONAL_CARD
	CampaignID  string                  `json:"campaignID,omitempty"` // SOCIAL#, P2P#TRANSFER_MONEY, #INLINE#
	CreatedAt   int64                   `json:"createdAt,omitempty"`
	UpdatedAt   int64                   `json:"updatedAt,omitempty"`
	IsDelete    bool                    `json:"isDelete,omitempty"`
	Text        string                  `json:"text,omitempty"`
	InfoCard    *ChatMessageInfoCard    `json:"infoCard,omitempty"`
	Transaction *ChatMessageTransaction `json:"transaction,omitempty"`
}

type ChatMessageInfoCard struct {
	Title       string            `json:"title,omitempty"`
	Status      map[string]string `json:"status,omitempty"`
	Amount      float64           `json:"amount,omitempty"`
	Content     string            `json:"content,omitempty"`
	StickerURL  string            `json:"stickerURL,omitempty"`
	BottomImage string            `json:"bottomImage,omitempty"`
}

type ChatMessageTransaction struct {
	Id      string  `json:"id,omitempty"`
	Amount  float64 `json:"amount,omitempty"`
	Message string  `json:"message,omitempty"`
	Sender  string  `json:"sender,omitempty"`
}
