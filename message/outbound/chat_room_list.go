package outbound

import (
	"momo-api/message/common"
)

type ChatRoomListBody struct {
	TargetLastMessageAt int64                  `json:"targetLastMessageAt"`
	Limit               int                    `json:"limit"`
	Filter              *common.ChatRoomFilter `json:"filter"`
}
