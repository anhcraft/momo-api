package noti

import (
	"encoding/json"
	"errors"
	"momo-api/api/enum"
	"momo-api/api/model"
	"momo-api/message"
	"momo-api/message/common"
	"momo-api/message/inbound"
	"momo-api/message/outbound"
	"momo-api/utils"
	"time"
)

type NotificationIterator struct {
	user     *model.User
	limit    int
	fromTime int64
	toTime   int64
	Buffer   []*model.Notification
	cursor   string
	depleted bool
}

func (it *NotificationIterator) list() ([]*model.Notification, string, error) {
	if !it.user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	req := message.CreateRequestBuilder()
	if err := req.SetBodyJSON(&outbound.NotificationListBody{
		Cursor:   it.cursor,
		FromTime: it.fromTime,
		Limit:    it.limit,
		ToTime:   it.toTime,
		UserId:   it.user.Phone,
	}); err != nil {
		return nil, "", err
	}
	req.SetURL(utils.CloudEndpoint, "hydra/v2/user/noti")
	res, err := req.Fetch(it.user)
	if err != nil {
		return nil, "", err
	}
	r := inbound.ParseNotificationListResponse(res)
	if r == nil {
		return nil, "", errors.New("can not parse response")
	} else if r.Success() {
		if r.Message == nil || r.Message.Data == nil || r.Message.Data.Notifications == nil {
			return make([]*model.Notification, 0), "", nil
		}
		list := make([]*model.Notification, len(r.Message.Data.Notifications))
		for i, v := range r.Message.Data.Notifications {
			notiType := enum.NotificationType(int(v.Type))

			var trans *model.P2PTransaction
			if notiType == enum.NOTI_RECEIVE_MONEY_P2P || notiType == enum.NOTI_RECEIVE_MONEY_P2P_TO_CHAT {
				extra := &common.NotificationExtraTransaction{}
				if err = json.Unmarshal([]byte(v.Extra), &extra); err == nil {
					if extra.ReceiverType == 1 {
						trans = &model.P2PTransaction{
							Id:      extra.TranId,
							Amount:  extra.Amount,
							Message: extra.Comment,
							Sender:  model.CreatePartner(extra.PartnerId, extra.PartnerName),
							// we create partner version (so serialization will not copy the entire userdata)
							Receiver: model.CreatePartner(it.user.Phone, it.user.Name),
							Time:     int64(v.Time),
						}
					} else {
						// TODO implement case ReceiverType == 0
						panic("case ReceiverType == 0 need to be implemented")
					}
				}
			}

			list[i] = &model.Notification{
				Id:             v.Id,
				Type:           notiType,
				Caption:        v.Caption,
				Body:           v.Body,
				IsRead:         false,
				Sender:         v.Sender,
				Time:           int64(v.Time),
				RefId:          v.RefId,
				P2PTransaction: trans,
			}
		}
		return list, r.Message.Data.Cursor, nil
	} else {
		return nil, "", errors.New(r.Message.ResponseInfo.ErrorDesc)
	}
}

func (it *NotificationIterator) HasNext() bool {
	return !it.depleted
}

func (it *NotificationIterator) Next() (bool, error) {
	if it.depleted {
		return false, nil
	}
	items, cursor, err := it.list()
	if err != nil {
		return false, err
	}
	it.Buffer = items
	if len(items) < it.limit || cursor == "" {
		it.depleted = true
	} else {
		it.cursor = cursor
	}
	return it.depleted, nil
}

func ListNotification(user *model.User, limit int, fromTime int64, toTime int64) *NotificationIterator {
	return &NotificationIterator{
		user:     user,
		limit:    limit,
		cursor:   "",
		fromTime: fromTime,
		toTime:   toTime,
		Buffer:   []*model.Notification{},
		depleted: false,
	}
}

func ListRecentNotification(user *model.User, limit int, duration time.Duration) *NotificationIterator {
	return ListNotification(user, limit, time.Now().Add(-duration).UnixMilli(), time.Now().UnixMilli())
}

func ListDefaultRecentNotification(user *model.User) *NotificationIterator {
	return ListNotification(user, 20, time.Now().AddDate(0, -3, 0).UnixMilli(), time.Now().UnixMilli())
}

func MarkReadNotification(user *model.User, id string) (bool, error) {
	if !user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	req := message.CreateRequestBuilder()
	if err := req.SetBodyJSON(&outbound.NotificationReadBody{
		ID:     id,
		IsRead: true,
	}); err != nil {
		return false, err
	}
	req.SetURL(utils.CloudEndpoint, "hydra/v2/user/read")
	res, err := req.Fetch(user)
	if err != nil {
		return false, err
	}
	r := inbound.ParseNotificationReadResponse(res)
	if r == nil {
		return false, errors.New("can not parse response")
	} else if r.Success {
		return true, nil
	} else {
		return false, errors.New(r.Message)
	}
}
