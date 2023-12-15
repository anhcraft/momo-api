package inbound

import "encoding/json"

type ChatMessageListResponse struct {
	Success bool                         `json:"success"`
	JSON    *ChatMessageListResponseJSON `json:"json"`
}

func ParseChatMessageListResponse(body []byte) *ChatMessageListResponse {
	data := &ChatMessageListResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}

type ChatMessageListResponseJSON struct {
	Messages        []*ChatMessageListResponseMessage `json:"messages"`
	LastReadMessage map[string]string                 `json:"lastReadMessage"` // Map<User-Id, Time>
}

type ChatMessageListResponseMessage struct {
	SenderID      string                               `json:"senderId"`
	RoomID        string                               `json:"roomId"`
	MessageID     string                               `json:"messageId"`
	MessageSource string                               `json:"messageSource"` // SOCIAL
	RequestID     string                               `json:"requestId"`
	ActorID       string                               `json:"actorId"`
	MessageType   string                               `json:"messageType"` // TEXT, INFORMATIONAL_CARD
	CampaignID    string                               `json:"campaignId"`  // SOCIAL#, P2P#TRANSFER_MONEY, #INLINE#
	Mention       []interface{}                        `json:"mention"`
	TemplateData  *ChatMessageListResponseTemplateData `json:"templateData"`
	Text          string                               `json:"text"`
	ActionData    *ChatMessageListResponseActionData   `json:"actionData"`
	HideFrom      []interface{}                        `json:"hideFrom"`
	CreatedAt     int64                                `json:"createdAt"`
	UpdatedAt     int64                                `json:"updatedAt"`
	IsDelete      bool                                 `json:"isDelete"`
}

type ChatMessageListResponseTemplateData struct {
	Items *ChatMessageListResponseItems `json:"items"`
}

type ChatMessageListResponseItems struct {
	Default *ChatMessageListResponseDefault `json:"default"`
}

type ChatMessageListResponseDefault struct {
	Body interface{} `json:"body"` // ChatMessageListResponseGenericBody, ChatMessageListResponseCardBody
}

type ChatMessageListResponseGenericBody struct {
	Content struct {
		Default string `json:"default"`
	} `json:"content"`
}

type ChatMessageListResponseCardBody struct {
	Title       string            `json:"title"`
	Status      map[string]string `json:"status"`
	Amount      interface{}       `json:"amount"` // either string or float
	Content     string            `json:"content"`
	StickerURL  string            `json:"stickerUrl"`
	BottomImage string            `json:"bottomImage"`
}

type ChatMessageListResponseActionData struct {
	ActionOnCard struct {
		ActionType       string      `json:"actionType"`  // REDIRECT
		FeatureCode      string      `json:"featureCode"` // transaction_history_detail
		ForwardingParams interface{} `json:"forwardingParams"`
	} `json:"actionOnCard"`
}

type ChatMessageListResponseTransactionTransferData struct {
	TranID string `json:"tranId"`
}
