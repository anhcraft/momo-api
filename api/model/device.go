package model

import "momo-api/message/common"

type Device struct {
	RegDeviceMsg *common.RegDeviceMsg `json:"regDeviceMsg,omitempty"`

	// Generate on RequestOTP
	DeviceToken   string `json:"deviceToken,omitempty"`
	ModelId       string `json:"modelId,omitempty"`
	FirebaseToken string `json:"firebaseToken,omitempty"`
	RKey          string `json:"RKey,omitempty"`

	// Generate on VerifyOTP
	OHash string `json:"OHash,omitempty"` // OTP Hash
}
