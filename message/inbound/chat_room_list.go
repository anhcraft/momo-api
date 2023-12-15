package inbound

import "encoding/json"

type ChatRoomListResponse struct {
	Success bool                      `json:"success"`
	JSON    *ChatRoomListResponseJSON `json:"json"`
}

func ParseChatRoomListResponse(body []byte) *ChatRoomListResponse {
	data := &ChatRoomListResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}

type ChatRoomListResponseJSON struct {
	Data   []*ChatRoomListResponseRoom `json:"data"`
	IsLast bool                        `json:"isLast"`
}

type ChatRoomListResponseRoom struct {
	RoomSource      string                                  `json:"roomSource"`
	RoomType        string                                  `json:"roomType"`
	RoomID          string                                  `json:"roomId"`
	IsGroup         bool                                    `json:"isGroup"`
	IsBlock         bool                                    `json:"isBlock"`
	IsHide          bool                                    `json:"isHide"`
	Name            string                                  `json:"name"`
	Avatar          string                                  `json:"avatar"`
	Members         []string                                `json:"members"` // user-Id
	Creator         string                                  `json:"creator"` // user-Id
	Unread          bool                                    `json:"unread"`
	LastMessageAt   int64                                   `json:"lastMessageAt"`
	LastMessageID   string                                  `json:"lastMessageId"`
	Mute            []interface{}                           `json:"mute"`
	IsMute          bool                                    `json:"isMute"`
	IsPin           bool                                    `json:"isPin"`
	Roles           map[string]string                       `json:"roles"` // Map<user-Id, Role> (ADMIN, MEMBER)
	IsBrowseMembers bool                                    `json:"isBrowseMembers"`
	BeAdd           string                                  `json:"beAdd"`
	Description     string                                  `json:"description"`
	LastMessage     string                                  `json:"lastMessage"`
	Profiles        map[string]*ChatRoomListResponseProfile `json:"profiles"` // Map<user-Id, Profile>
	ConfigID        string                                  `json:"configId"`
	LastReadMessage map[string]string                       `json:"lastReadMessage"` // Map<user-Id, Message-Id>
	Pin             []interface{}                           `json:"pin"`
	Hide            []interface{}                           `json:"hide"`
	CustomData      string                                  `json:"customData"`
}

type ChatRoomListResponseProfile struct {
	Name       string `json:"name"`
	Gender     int    `json:"gender"`
	Avatar     string `json:"avatar"`
	AgentID    string `json:"agentId"`
	IsBirthday bool   `json:"isBirthday"`
}
