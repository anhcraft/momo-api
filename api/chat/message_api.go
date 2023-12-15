package chat

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"momo-api/api/model"
	"momo-api/message"
	"momo-api/message/inbound"
	"momo-api/message/outbound"
	"momo-api/utils"
)

type MessageIterator struct {
	user                *model.User
	limit               int
	roomId              string
	roomSource          string
	Buffer              []*model.ChatMessage
	targetLastMessageId string
	depleted            bool
}

func (it *MessageIterator) list() ([]*model.ChatMessage, error) {
	if !it.user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	req := message.CreateRequestBuilder()
	if err := req.SetBodyJSON(&outbound.ChatMessageListBody{
		RoomSource:      it.roomSource,
		RoomID:          it.roomId,
		UserID:          it.user.Phone,
		TargetMessageID: it.targetLastMessageId,
		Limit:           it.limit,
		Action:          1,
	}); err != nil {
		return nil, err
	}
	req.SetURL(utils.CloudEndpoint, "coremessage/load")
	res, err := req.Fetch(it.user)
	if err != nil {
		return nil, err
	}
	r := inbound.ParseChatMessageListResponse(res)
	if r == nil {
		return nil, errors.New("can not parse response")
	} else if r.Success {
		if r.JSON == nil {
			return make([]*model.ChatMessage, 0), nil
		}
		list := make([]*model.ChatMessage, len(r.JSON.Messages))
		for i, v := range r.JSON.Messages {
			var infoCard *model.ChatMessageInfoCard
			if v.MessageType == "INFORMATIONAL_CARD" {
				var body *inbound.ChatMessageListResponseCardBody
				if err1 := mapstructure.Decode(v.TemplateData.Items.Default.Body, &body); err1 != nil {
					return nil, err1
				}
				infoCard = &model.ChatMessageInfoCard{
					Title:       body.Title,
					Status:      body.Status,
					Amount:      cast.ToFloat64(body.Amount),
					Content:     body.Content,
					StickerURL:  body.StickerURL,
					BottomImage: body.BottomImage,
				}
			}

			var trans *model.ChatMessageTransaction
			if v.MessageType == "INFORMATIONAL_CARD" && v.CampaignID == "P2P#TRANSFER_MONEY" {
				var data *inbound.ChatMessageListResponseTransactionTransferData
				if err1 := mapstructure.Decode(v.ActionData.ActionOnCard.ForwardingParams, &data); err1 != nil {
					return nil, err1
				}
				trans = &model.ChatMessageTransaction{
					Id:      data.TranID,
					Amount:  infoCard.Amount,
					Message: infoCard.Content,
					Sender:  v.SenderID,
				}
			}

			list[i] = &model.ChatMessage{
				Id:          v.MessageID,
				Sender:      v.SenderID,
				Room:        v.RoomID,
				Source:      v.MessageSource,
				RequestId:   v.RequestID,
				MsgType:     v.MessageType,
				CampaignID:  v.CampaignID,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				IsDelete:    v.IsDelete,
				Text:        v.Text,
				InfoCard:    infoCard,
				Transaction: trans,
			}
		}
		return list, nil
	} else {
		return nil, errors.New("request failed")
	}
}

func (it *MessageIterator) HasNext() bool {
	return !it.depleted
}

func (it *MessageIterator) Next() (bool, error) {
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
		it.targetLastMessageId = items[len(items)-1].Id
	}
	return it.depleted, nil
}

func ListMessages(user *model.User, limit int, roomId string, roomSource string, targetLastMessageId string) *MessageIterator {
	return &MessageIterator{
		user:                user,
		limit:               limit,
		roomId:              roomId,
		roomSource:          roomSource,
		Buffer:              []*model.ChatMessage{},
		targetLastMessageId: targetLastMessageId,
		depleted:            false,
	}
}

func ListRecentMessages(user *model.User, limit int, roomId string, roomSource string) *MessageIterator {
	return ListMessages(user, limit, roomId, roomSource, "")
}

func ListDefaultRecentMessages(user *model.User, room *model.ChatRoom) *MessageIterator {
	return ListMessages(user, 20, room.Id, room.Source, "")
}
