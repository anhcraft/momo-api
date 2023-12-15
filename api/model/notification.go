package model

import (
	"momo-api/api/enum"
)

type Notification struct {
	Id             string                `json:"id,omitempty"`
	Type           enum.NotificationType `json:"type,omitempty"`
	Caption        string                `json:"caption,omitempty"`
	Body           string                `json:"body,omitempty"`
	IsRead         bool                  `json:"isRead,omitempty"`
	Sender         string                `json:"sender,omitempty"` // Sender (a system service or a user)
	Time           int64                 `json:"time,omitempty"`
	RefId          string                `json:"refId,omitempty"`       // Reference Id
	P2PTransaction *P2PTransaction       `json:"transaction,omitempty"` // Only if related to P2PTransaction
}
