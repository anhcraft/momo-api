package chat

import (
	"errors"
	"momo-api/api/model"
	"momo-api/message"
	"momo-api/message/common"
	"momo-api/message/inbound"
	"momo-api/message/outbound"
	"momo-api/utils"
	"time"
)

type RoomIterator struct {
	user                *model.User
	limit               int
	filter              *common.ChatRoomFilter
	Buffer              []*model.ChatRoom
	targetLastMessageAt int64
	depleted            bool
}

func (it *RoomIterator) list() ([]*model.ChatRoom, error) {
	if !it.user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	req := message.CreateRequestBuilder()
	if err := req.SetBodyJSON(&outbound.ChatRoomListBody{
		TargetLastMessageAt: it.targetLastMessageAt,
		Limit:               it.limit,
		Filter:              it.filter,
	}); err != nil {
		return nil, err
	}
	req.SetURL(utils.CloudEndpoint, "coreroom/v1/room/load")
	res, err := req.Fetch(it.user)
	if err != nil {
		return nil, err
	}
	r := inbound.ParseChatRoomListResponse(res)
	if r == nil {
		return nil, errors.New("can not parse response")
	} else if r.Success {
		if r.JSON == nil {
			return make([]*model.ChatRoom, 0), nil
		}
		list := make([]*model.ChatRoom, len(r.JSON.Data))
		for i, v := range r.JSON.Data {
			list[i] = &model.ChatRoom{
				Source:        v.RoomSource,
				Type:          v.RoomType,
				Id:            v.RoomID,
				IsGroup:       v.IsGroup,
				IsBlock:       v.IsBlock,
				IsHide:        v.IsHide,
				Members:       v.Members,
				Creator:       v.Creator,
				Unread:        v.Unread,
				LastMessageAt: v.LastMessageAt,
				LastMessageID: v.LastMessageID,
				IsMute:        v.IsMute,
				LastMessage:   v.LastMessage,
			}
		}
		return list, nil
	} else {
		return nil, errors.New("request failed")
	}
}

func (it *RoomIterator) HasNext() bool {
	return !it.depleted
}

func (it *RoomIterator) Next() (bool, error) {
	if it.depleted {
		return false, nil
	}
	items, err := it.list()
	if err != nil {
		return false, err
	}
	it.Buffer = items
	if len(items) < it.limit {
		it.depleted = true
	} else {
		it.targetLastMessageAt = items[len(items)-1].LastMessageAt
	}
	return it.depleted, nil
}

func ListRooms(user *model.User, limit int, filter *common.ChatRoomFilter, targetLastMessageAt int64) *RoomIterator {
	return &RoomIterator{
		user:                user,
		limit:               limit,
		filter:              filter,
		Buffer:              []*model.ChatRoom{},
		targetLastMessageAt: targetLastMessageAt,
		depleted:            false,
	}
}

func ListRecentRooms(user *model.User, limit int, filter *common.ChatRoomFilter) *RoomIterator {
	return ListRooms(user, limit, filter, time.Now().UnixMilli())
}

func ListDefaultRecentRooms(user *model.User) *RoomIterator {
	return ListRooms(user, 20, &common.ChatRoomFilter{
		RoomSources: []string{"SOCIAL", "ESCROW"},
		RoomTypes:   []string{"NORMAL_ROOM", "REQUEST_ROOM", "MONEY_POOL"},
		IsHide:      false,
		IsBlock:     false,
	}, time.Now().UnixMilli())
}
