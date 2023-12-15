package trans

import (
	"encoding/json"
	"errors"
	"momo-api/api/model"
	"momo-api/message"
	"momo-api/message/inbound"
	"momo-api/message/outbound"
	"momo-api/utils"
	"momo-api/utils/crypto"
	"strconv"
	"time"
)

type TransactionIterator struct {
	user     *model.User
	limit    int
	offset   int
	fromTime int64
	toTime   int64
	Buffer   []*model.P2PTransaction
	depleted bool
}

func (it *TransactionIterator) list() ([]*model.P2PTransaction, bool, error) {
	if !it.user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	req := message.CreateRequestBuilder()
	req.EnableBodyEncryption()
	if err := req.SetBodyJSON(&outbound.TransactionBrowseBody{
		RequestId:   strconv.FormatInt(time.Now().UnixMilli(), 10),
		StartDate:   it.fromTime,
		EndDate:     it.toTime,
		Offset:      it.offset,
		Limit:       it.limit,
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		DeviceOs:    it.user.Device.RegDeviceMsg.DeviceOs,
	}); err != nil {
		return nil, false, err
	}
	req.SetURL(utils.MomoApiEndpoint, "transhis/api/transhis/browse")
	res, err := req.Fetch(it.user)
	if err != nil {
		return nil, false, err
	}
	r := inbound.ParseTransactionBrowseResponse(res)
	if r == nil {
		return nil, false, errors.New("can not parse response")
	} else if r.Success() {
		data := make([]*model.P2PTransaction, 0)
		for _, v := range r.MomoMsg {
			if v.ServiceID == "transfer_masking" || v.ServiceID == "transfer_p2p" {
				if v.MoneySource == 1 {
					t := &model.P2PTransaction{
						Id:       strconv.FormatInt(v.TransID, 10),
						Amount:   float64(v.TotalAmount),
						Message:  "",
						Sender:   model.CreatePartner(v.SourceID, v.SourceName),
						Receiver: model.CreatePartner(v.TargetID, v.TargetName),
						Time:     v.CreatedAt,
					}
					if t.Sender.Name == t.Sender.Phone && t.Sender.Phone == it.user.Phone {
						t.Sender.Name = it.user.Name
					}
					if t.Receiver.Name == t.Receiver.Phone && t.Receiver.Phone == it.user.Phone {
						t.Receiver.Name = it.user.Name
					}
					data = append(data, t)
				}
			}
		}
		return data, r.ListOver, nil
	} else {
		return nil, false, errors.New(r.Message)
	}
}

func (it *TransactionIterator) HasNext() bool {
	return !it.depleted
}

func (it *TransactionIterator) Next() (bool, error) {
	if it.depleted {
		return false, nil
	}
	items, over, err := it.list()
	if err != nil {
		return false, err
	}
	it.Buffer = items
	if len(items) < it.limit || over {
		it.depleted = true
	} else {
		it.offset = len(items)
	}
	return it.depleted, nil
}

func BrowseTransactions(user *model.User, limit int, offset int, fromTime int64, toTime int64) *TransactionIterator {
	return &TransactionIterator{
		user:     user,
		limit:    limit,
		offset:   offset,
		fromTime: fromTime,
		toTime:   toTime,
		Buffer:   []*model.P2PTransaction{},
		depleted: false,
	}
}

func BrowseRecentTransactions(user *model.User, limit int, duration time.Duration) *TransactionIterator {
	return BrowseTransactions(user, limit, 0, time.Now().Add(-duration).UnixMilli(), time.Now().UnixMilli())
}

func BrowseDefaultRecentTransactions(user *model.User) *TransactionIterator {
	return BrowseTransactions(user, 20, 0, time.Now().AddDate(0, -3, 0).UnixMilli(), time.Now().UnixMilli())
}

func GetTransactionDetail(user *model.User, trans *model.P2PTransaction) error {
	req := message.CreateRequestBuilder()
	req.EnableBodyEncryption()
	if err := req.SetBodyJSON(&outbound.TransactionDetailBody{
		RequestId:   strconv.FormatInt(time.Now().UnixMilli(), 10),
		TransId:     trans.Id,
		ServiceId:   "transfer_p2p",
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		DeviceOs:    user.Device.RegDeviceMsg.DeviceOs,
	}); err != nil {
		return err
	}
	req.SetURL(utils.MomoApiEndpoint, "transhis/api/transhis/detail")
	res, err := req.Fetch(user)
	if err != nil {
		return err
	}
	r := inbound.ParseTransactionDetailResponse(res)
	if r == nil {
		return errors.New("can not parse response")
	} else if r.Success() {
		if r.MomoMsg != nil {
			v := r.MomoMsg
			if v.ServiceID == "transfer_masking" || v.ServiceID == "transfer_p2p" {
				if v.MoneySource == 1 {
					if trans.Amount == 0 {
						trans.Amount = float64(v.TotalAmount)
					}
					if trans.Sender == nil {
						trans.Sender = model.CreatePartner(v.SourceID, v.SourceName)
					}
					if trans.Receiver == nil {
						trans.Receiver = model.CreatePartner(v.TargetID, user.Name)
					}
					if trans.Time == 0 {
						trans.Time = v.CreatedAt
					}
				}

				if trans.Message == "" {
					data := &inbound.TransactionDetailP2PTransServiceData{}
					if err = json.Unmarshal([]byte(v.ServiceData), &data); err == nil {
						trans.Message = data.CommentValue
					}
				}
			}
		}
		return nil
	} else {
		return errors.New(r.Message)
	}
}

func InitTransfer(trans *model.P2PTransaction) (*inbound.MoneyTransferInitResponse, error) {
	picoTime := strconv.FormatInt(time.Now().UnixMilli()*1e6, 10)
	gigaTime := strconv.FormatFloat(float64(time.Now().UnixMilli())/1e12, 'f', -1, 64)
	checkSum, err := crypto.EncryptAes256CbcPKCS7ToBase64(
		[]byte(trans.Sender.Phone+picoTime+"M2MU_INIT"+gigaTime+"E12"),
		[]byte(trans.Sender.Profile.SetupKey[:32]),
	)
	if err != nil {
		return nil, err
	}
	req := message.CreateRequestBuilder()
	req.EnableBodyEncryption()
	if err = req.SetBodyJSON(&outbound.MoneyTransferInitBody{
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		CmdId:       utils.GenerateCmdId(),
		DeviceOs:    trans.Sender.Device.RegDeviceMsg.DeviceOs,
		Extra: outbound.MoneyTransferInitExtra{
			CheckSum: checkSum,
		},
		Lang: utils.LangCode,
		MomoMsg: outbound.MoneyTransferInitMsg{
			Class:       "mservice.backend.entity.msg.M2MUInitMsg",
			ServiceId:   "transfer_p2p",
			ServiceCode: "transfer_p2p",
			ClientTime:  time.Now().UnixMilli(),
			TranType:    2018,
			Comment:     trans.Message,
			Ref:         "",
			PartnerId:   trans.Receiver.Phone,
			Amount:      trans.Amount,
			Extras:      "{\"appSendChat\":false,\"themeP2P\":\"\",\"contactName\":\"\",\"stickers\":\"\"}",
			TranList: []outbound.MoneyTransferInitTrans{
				{
					PartnerName:    trans.Receiver.Name,
					PartnerId:      trans.Receiver.Phone,
					OriginalAmount: trans.Amount,
					TransferSource: "transfer_via_chat",
					MoneySource:    1,
				},
			},
			MoneySource: 1,
		},
		MsgType:        "M2MU_INIT",
		ConfirmMsgType: "M2MU_CONFIRM",
		ConfirmClass:   "mservice.backend.entity.msg.M2MUConfirmMsg",
		Time:           time.Now().UnixMilli(),
		User:           trans.Sender.Phone,
		PaymentInfo: []outbound.MoneyTransferInitPaymentInfo{
			{
				Title:  "Chuyển đến",
				Value:  trans.Receiver.Name,
				Format: "",
			},
			{
				Title:  "Số điện thoại",
				Value:  trans.Receiver.Phone,
				Format: "phone",
			},
			{
				Title:  "Số tiền",
				Value:  trans.Amount,
				Format: "currency",
			},
		},
	}); err != nil {
		return nil, err
	}
	req.SetURL(utils.MomoOwaEndpoint, "api/M2MU_INIT")
	req.SetHeader("MsgType", "M2MU_INIT")
	res, err := req.Fetch(trans.Sender)
	if err != nil {
		return nil, err
	}
	r := inbound.ParseMoneyTransferInitResponse(res)
	if r == nil {
		return nil, errors.New("can not parse response")
	} else if r.Success() {
		return r, nil
	} else {
		return nil, errors.New(r.ErrorDesc)
	}
}

func ConfirmTransfer(trans *model.P2PTransaction, rep *inbound.MoneyTransferInitResponse) error {
	picoTime := strconv.FormatInt(time.Now().UnixMilli()*1e6, 10)
	gigaTime := strconv.FormatFloat(float64(time.Now().UnixMilli())/1e12, 'f', -1, 64)
	checkSum, err := crypto.EncryptAes256CbcPKCS7ToBase64(
		[]byte(trans.Sender.Phone+picoTime+"M2MU_CONFIRM"+gigaTime+"E12"),
		[]byte(trans.Sender.Profile.SetupKey[:32]),
	)
	if err != nil {
		return err
	}
	req := message.CreateRequestBuilder()
	req.EnableBodyEncryption()
	if err = req.SetBodyJSON(&outbound.MoneyTransferConfirmBody{
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		CmdId:       utils.GenerateCmdId(),
		DeviceOs:    trans.Sender.Device.RegDeviceMsg.DeviceOs,
		Extra: outbound.MoneyTransferConfirmExtra{
			CheckSum: checkSum,
		},
		Lang: utils.LangCode,
		MomoMsg: outbound.MoneyTransferConfirmMsg{
			Class:       "mservice.backend.entity.msg.M2MUConfirmMsg",
			ServiceId:   "transfer_p2p",
			ServiceCode: "transfer_p2p",
			ClientTime:  time.Now().UnixMilli(),
			TranType:    2018,
			Comment:     trans.Message,
			Ref:         "",
			PartnerId:   trans.Receiver.Phone,
			Amount:      trans.Amount,
			Extras:      "{\"appSendChat\":false,\"themeP2P\":\"\",\"contactName\":\"\",\"stickers\":\"\"}",
			Id:          rep.MomoMsg.ReplyMsgs[0].ID,
			Ids:         []string{rep.MomoMsg.ReplyMsgs[0].ID},
			TranList: []outbound.MoneyTransferConfirmTrans{
				{
					PartnerName:    trans.Receiver.Name,
					PartnerId:      trans.Receiver.Phone,
					OriginalAmount: trans.Amount,
					TransferSource: "transfer_via_chat",
					MoneySource:    1,
				},
			},
			MoneySource:            1,
			TotalAmount:            rep.Extra.TotalAmount,
			OriginalAmount:         rep.Extra.OriginalAmount,
			OriginalClass:          rep.Extra.OriginalClass,
			OriginalPhone:          rep.Extra.OriginalPhone,
			TotalFee:               rep.Extra.TotalFee,
			CheckSum:               rep.Extra.CheckSum,
			GetUserInfoTaskRequest: rep.Extra.GetUserInfoTaskRequest,
		},
		MsgType:        "M2MU_CONFIRM",
		ConfirmMsgType: "M2MU_CONFIRM",
		ConfirmClass:   "mservice.backend.entity.msg.M2MUConfirmMsg",
		Time:           time.Now().UnixMilli(),
		User:           trans.Sender.Phone,
		PaymentInfo: []outbound.MoneyTransferConfirmPaymentInfo{
			{
				Title:  "Chuyển đến",
				Value:  trans.Receiver.Name,
				Format: "",
			},
			{
				Title:  "Số điện thoại",
				Value:  trans.Receiver.Phone,
				Format: "phone",
			},
			{
				Title:  "Số tiền",
				Value:  trans.Amount,
				Format: "currency",
			},
		},
	}); err != nil {
		return err
	}
	req.SetURL(utils.MomoOwaEndpoint, "api/M2MU_CONFIRM")
	req.SetHeader("MsgType", "M2MU_CONFIRM")
	res, err := req.Fetch(trans.Sender)
	if err != nil {
		return err
	}
	r := inbound.ParseMoneyTransferConfirmResponse(res)
	if r == nil {
		return errors.New("can not parse response")
	} else if r.Success() {
		return nil
	} else {
		return errors.New(r.ErrorDesc)
	}
}
