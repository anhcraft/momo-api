package auth

import (
	"errors"
	"github.com/spf13/cast"
	"momo-api/api/model"
	"momo-api/message"
	"momo-api/message/inbound"
	"momo-api/message/outbound"
	"momo-api/utils"
	"momo-api/utils/crypto"
	"strconv"
	"time"
)

func VerifyPhone(user *model.User) (error, bool) {
	if !user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	req := message.CreateRequestBuilder()
	if err := req.SetBodyJSON(&outbound.VerifyPhoneBody{
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		CmdId:       utils.GenerateCmdId(),
		DeviceOs:    user.Device.RegDeviceMsg.DeviceOs,
		Lang:        utils.LangCode,
		MomoMsg:     *user.Device.RegDeviceMsg,
		MsgType:     "CHECK_USER_BE_MSG",
		Time:        time.Now().UnixMilli(),
		User:        user.Phone,
	}); err != nil {
		return err, false
	}
	req.SetURL(utils.MomoApiEndpoint, "backend/auth-app/public/CHECK_USER_BE_MSG")
	req.SetHeader("MsgType", "CHECK_USER_BE_MSG")
	res, err := req.Fetch(user)
	if err != nil {
		return err, false
	}
	r := inbound.ParseUserMsgCheckResponse(res)
	if r == nil {
		return errors.New("can not parse response"), false
	} else if r.Success() {
		return nil, cast.ToBool(r.Extra.IsChangeDevice)
	} else {
		return errors.New(r.ErrorDesc), false
	}
}

func RequestOTP(user *model.User) error {
	if !user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	user.Device.DeviceToken = utils.GenerateDeviceToken()
	user.Device.FirebaseToken = utils.GenerateFirebaseToken()
	user.Device.ModelId = utils.GenerateModelId()
	user.Device.RKey = utils.GenerateRKey()

	req := message.CreateRequestBuilder()
	if err := req.SetBodyJSON(&outbound.RequestOTPBody{
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		CmdId:       utils.GenerateCmdId(),
		DeviceOs:    user.Device.RegDeviceMsg.DeviceOs,
		Extra: outbound.RequestOTPExtra{
			DeviceToken:          user.Device.DeviceToken,
			Idfa:                 "",
			ModelId:              user.Device.ModelId,
			OneSignalToken:       user.Device.FirebaseToken,
			RequireHashStringOtp: true,
			SecureId:             user.Device.RegDeviceMsg.SecureId,
			Simulator:            false,
			Token:                user.Device.FirebaseToken,
			Action:               "SEND",
			IsVoice:              true,
			RKey:                 user.Device.RKey,
		},
		Lang:    utils.LangCode,
		MomoMsg: *user.Device.RegDeviceMsg,
		MsgType: "SEND_OTP_MSG",
		Time:    time.Now().UnixMilli(),
		User:    user.Phone,
	}); err != nil {
		return err
	}
	req.SetURL(utils.MomoApiEndpoint, "backend/otp-app/public/")
	req.SetHeader("MsgType", "SEND_OTP_MSG")
	res, err := req.Fetch(user)
	if err != nil {
		return err
	}
	r := inbound.ParseGenericResponse(res)
	if r == nil {
		return errors.New("can not parse response")
	} else if r.Success() {
		return nil
	} else {
		return errors.New(r.ErrorDesc)
	}
}

func VerifyOTP(user *model.User, otp string) error {
	if !user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	user.Device.OHash = crypto.HashSha256([]byte(user.Phone + user.Device.RKey + otp))
	req := message.CreateRequestBuilder()
	if err := req.SetBodyJSON(&outbound.VerifyOtpBody{
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		CmdId:       utils.GenerateCmdId(),
		DeviceOs:    user.Device.RegDeviceMsg.DeviceOs,
		Extra: outbound.VerifyOTPExtra{
			DeviceToken:    user.Device.DeviceToken,
			Idfa:           "",
			ModelId:        user.Device.ModelId,
			OneSignalToken: user.Device.FirebaseToken,
			SecureId:       user.Device.RegDeviceMsg.SecureId,
			Simulator:      false,
			Token:          user.Device.FirebaseToken,
			OHash:          user.Device.OHash,
		},
		Lang:    utils.LangCode,
		MomoMsg: *user.Device.RegDeviceMsg,
		MsgType: "REG_DEVICE_MSG",
		Time:    time.Now().UnixMilli(),
		User:    user.Phone,
	}); err != nil {
		return err
	}
	req.SetURL(utils.MomoApiEndpoint, "backend/otp-app/public/")
	req.SetHeader("MsgType", "REG_DEVICE_MSG")
	res, err := req.Fetch(user)
	if err != nil {
		return err
	}
	r := inbound.ParseVerifyOtpResponse(res)
	if r == nil {
		return errors.New("can not parse response")
	} else if r.Success() {
		v, err1 := crypto.DecryptAes256CbcPKCS7FromBase64(r.Extra.SetupKey, []byte(user.Device.OHash[:32]))
		if err1 != nil {
			return err1
		}
		user.Profile.SetupKey = string(v)
		return nil
	} else {
		return errors.New(r.ErrorDesc)
	}
}

func Login(user *model.User, password string) error {
	if !user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	pHash, err := crypto.EncryptAes256CbcPKCS7ToBase64(
		[]byte(user.Device.RegDeviceMsg.Imei+"|"+password),
		[]byte(user.Profile.SetupKey[:32]),
	)
	if err != nil {
		return err
	}
	picoTime := strconv.FormatInt(time.Now().UnixMilli()*1e6, 10)
	gigaTime := strconv.FormatFloat(float64(time.Now().UnixMilli())/1e12, 'f', -1, 64)
	checkSum, err := crypto.EncryptAes256CbcPKCS7ToBase64(
		[]byte(user.Phone+picoTime+"USER_LOGIN_MSG"+gigaTime+"E12"),
		[]byte(user.Profile.SetupKey[:32]),
	)
	if err != nil {
		return err
	}
	req := message.CreateRequestBuilder()
	err = req.SetBodyJSON(&outbound.LoginBody{
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		CmdId:       utils.GenerateCmdId(),
		DeviceOs:    user.Device.RegDeviceMsg.DeviceOs,
		Extra: outbound.LoginExtra{
			DeviceToken:    user.Device.DeviceToken,
			Idfa:           "",
			ModelId:        user.Device.ModelId,
			OneSignalToken: user.Device.FirebaseToken,
			SecureId:       user.Device.RegDeviceMsg.SecureId,
			Simulator:      false,
			Token:          user.Device.FirebaseToken,
			CheckSum:       checkSum,
			PHash:          pHash,
		},
		Lang: utils.LangCode,
		MomoMsg: outbound.LoginMsg{
			Class:   "mservice.backend.entity.msg.LoginMsg",
			IsSetup: false,
		},
		MsgType: "USER_LOGIN_MSG",
		Pass:    password,
		Time:    time.Now().UnixMilli(),
		User:    user.Phone,
	})
	if err != nil {
		return err
	}
	req.SetURL(utils.MomoOwaEndpoint, "public/login")
	req.SetHeader("MsgType", "USER_LOGIN_MSG")
	res, err := req.Fetch(user)
	if err != nil {
		return err
	}
	r := inbound.ParseLoginResponse(res)
	if r == nil {
		return errors.New("can not parse response")
	} else if r.Success() {
		// extra
		user.Profile.AuthToken = r.Extra.AuthToken
		user.Profile.SessionKey = r.Extra.SessionKey
		user.Profile.RequestEncryptKey = r.Extra.RequestEncryptKey
		user.Profile.RefreshToken = r.Extra.RefreshToken
		if balance, err1 := strconv.ParseUint(r.Extra.Balance, 10, 64); err1 == nil {
			user.Profile.Balance = balance
		}
		user.Name = r.Extra.FullName // TODO should be username
		user.Profile.FullName = r.Extra.FullName
		// momoMsg
		user.Profile.AgentId = strconv.FormatInt(int64(r.MomoMsg.AgentId), 10)
		user.Profile.Identified = r.MomoMsg.Identify != "UNCONFIRM"
		user.Profile.DailyTransferLimit = uint64(r.MomoMsg.Capset)
		if registerDate, err1 := strconv.ParseInt(r.MomoMsg.RegisterDate, 10, 64); err1 == nil {
			user.Profile.RegisterDate = registerDate
		}
		user.Profile.BankCardOwner = r.MomoMsg.BankVerifyName
		user.Profile.BankCardId = r.MomoMsg.BankCardId
		user.Profile.BankName = r.MomoMsg.BankName
		user.Profile.BankCode = r.MomoMsg.BankCode
		user.Profile.EmailVerified = r.MomoMsg.ValidateEmail != "UNCONFIRM"
		return nil
	} else {
		return errors.New(r.ErrorDesc)
	}
}

func Relogin(user *model.User) error {
	if !user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	req := message.CreateRequestBuilder()
	err := req.SetBodyJSON(&outbound.ReloginBody{
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		CmdId:       utils.GenerateCmdId(),
		DeviceOs:    user.Device.RegDeviceMsg.DeviceOs,
		Extra: outbound.ReloginExtra{
			DeviceToken:    user.Device.DeviceToken,
			Idfa:           "",
			ModelId:        user.Device.ModelId,
			OneSignalToken: user.Device.FirebaseToken,
			SecureId:       user.Device.RegDeviceMsg.SecureId,
			Simulator:      false,
			Token:          user.Device.FirebaseToken,
			RKey:           user.Device.RKey,
		},
		Lang:    utils.LangCode,
		MomoMsg: *user.Device.RegDeviceMsg,
		MsgType: "RE_LOGIN",
		Time:    time.Now().UnixMilli(),
		User:    user.Phone,
	})
	if err != nil {
		return err
	}
	req.SetURL(utils.MomoApiEndpoint, "backend/otp-app/public/RE_LOGIN")
	req.SetHeader("MsgType", "RE_LOGIN")
	res, err := req.Fetch(user)
	if err != nil {
		return err
	}
	// Relogin use the same VerifyOTP response
	r := inbound.ParseVerifyOtpResponse(res)
	if r == nil {
		return errors.New("can not parse response")
	} else if r.Success() {
		user.Profile.SetupKey = r.Extra.SetupKey
		return nil
	} else {
		return errors.New(r.ErrorDesc)
	}
}

func Logout(user *model.User) error {
	if !user.Standard() {
		panic("Non-standard user is unsupported for this operation")
	}
	req := message.CreateRequestBuilder()
	err := req.SetBodyJSON(&outbound.LogoutBody{
		AppCode:     utils.AppCode,
		AppId:       utils.AppId,
		AppVer:      utils.AppVersion,
		BuildNumber: utils.BuildNumber,
		Channel:     utils.Channel,
		CmdId:       utils.GenerateCmdId(),
		DeviceOs:    user.Device.RegDeviceMsg.DeviceOs,
		Lang:        utils.LangCode,
		MomoMsg: outbound.LogoutMsg{
			Class: "mservice.backend.entity.msg.ForwardMsg",
		},
		MsgType: "USER_LOGOUT_MSG",
		Time:    time.Now().UnixMilli(),
		User:    user.Phone,
	})
	if err != nil {
		return err
	}
	req.SetURL(utils.MomoApiEndpoint, "backend/auth-app/api/USER_LOGOUT_MSG")
	req.SetHeader("MsgType", "USER_LOGOUT_MSG")
	req.EnableBodyEncryption()
	res, err := req.Fetch(user)
	if err != nil {
		return err
	}
	r := inbound.ParseGenericResponse(res)
	if r == nil {
		return errors.New("can not parse response")
	} else if r.Success() {
		return nil
	} else {
		return errors.New(r.ErrorDesc)
	}
}
