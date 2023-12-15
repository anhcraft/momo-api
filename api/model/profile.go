package model

type Profile struct {
	SessionKeyTracking string `json:"sessionKeyTracking,omitempty"`

	// Get from verify OTP
	SetupKey string `json:"setupKey,omitempty"`

	// Get from login response
	// extra
	AuthToken         string `json:"authToken,omitempty"`
	SessionKey        string `json:"sessionKey,omitempty"`
	RequestEncryptKey string `json:"requestEncryptKey,omitempty"`
	RefreshToken      string `json:"refreshToken,omitempty"`
	Balance           uint64 `json:"balance,omitempty"`
	FullName          string `json:"fullName,omitempty"`

	// momoMsg
	AgentId            string `json:"agentId,omitempty"`
	Identified         bool   `json:"identified,omitempty"`         // Id card verified
	DailyTransferLimit uint64 `json:"dailyTransferLimit,omitempty"` // daily maximum deposit/withdraw
	RegisterDate       int64  `json:"registerDate,omitempty"`
	BankCardOwner      string `json:"bankCardOwner,omitempty"`
	BankCardId         string `json:"bankCardId,omitempty"`
	BankName           string `json:"bankName,omitempty"`
	BankCode           string `json:"bankCode,omitempty"`
	EmailVerified      bool   `json:"emailVerified,omitempty"`
}
