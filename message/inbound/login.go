package inbound

import "encoding/json"

type LoginResponse struct {
	ErrorCode int                 `json:"errorCode"`
	ErrorDesc string              `json:"errorDesc"`
	Extra     *LoginExtraResponse `json:"extra"`
	MomoMsg   *LoginMsgResponse   `json:"momoMsg"`
}

func (r *LoginResponse) Success() bool {
	return r.ErrorCode == 0
}

func ParseLoginResponse(body []byte) *LoginResponse {
	data := &LoginResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}

type LoginExtraResponse struct {
	AccumulatedPoint    string  `json:"ACCUMULATED_POINT"`
	AuthToken           string  `json:"AUTH_TOKEN"`
	Balance             string  `json:"BALANCE"`
	Cashback            string  `json:"CASHBACK"`
	Email               string  `json:"EMAIL"`
	ExpireDesc          string  `json:"EXPIRE_DESC"`
	FirstTimeLogin      string  `json:"FIRST_TIME_LOGIN"`
	FullName            string  `json:"FULL_NAME"`
	GoldenPig           string  `json:"GOLDEN_PIG"`
	InvestmentBalance   string  `json:"INVESTMENT_BALANCE"`
	LastYearLevel       string  `json:"LAST_YEAR_LEVEL"`
	Level               string  `json:"LEVEL"`
	LevelPercent        string  `json:"LEVEL_PERCENT"`
	LockLevel           string  `json:"LOCK_LEVEL"`
	ModelId             string  `json:"MODELID"`
	NextLevel           string  `json:"NEXT_LEVEL"`
	OneSignalToken      string  `json:"ONESIGNAL_TOKEN"`
	PayLater            string  `json:"PAY_LAYTER"`
	Point               string  `json:"POINT"`
	PointToCurrentLevel string  `json:"POINT_TO_CURRENT_LEVEL"`
	PointToNextLevel    string  `json:"POINT_TO_NEXT_LEVEL"`
	PreLogin            string  `json:"PRE_LOGIN"`
	RefreshToken        string  `json:"REFRESH_TOKEN"`
	RequestEncryptKey   string  `json:"REQUEST_ENCRYPT_KEY"`
	SessionKey          string  `json:"SESSION_KEY"`
	Simulator           string  `json:"SIMULATOR"`
	Token               string  `json:"TOKEN"`
	VisaNewFlow         string  `json:"VISA_NEW_FLOW"`
	CheckSum            string  `json:"checkSum"`
	LockedUntilTime     float64 `json:"lockedUntillTime"`
	OriginalPhone       string  `json:"originalPhone"`
	PHash               string  `json:"pHash"`
}

type LoginMsgResponse struct {
	Class                   string  `json:"_class"`
	Address                 string  `json:"address"`
	AddressKyc              string  `json:"addressKyc"`
	AgentId                 float64 `json:"agentId"`
	AppCode                 string  `json:"appCode"`
	AppVer                  int     `json:"appVer"`
	AvatarUploadTime        float64 `json:"avatarUploadTime"`
	AvatarUrl               string  `json:"avatarUrl"`
	BankCardId              string  `json:"bankCardId"`
	BankCode                string  `json:"bankCode"`
	BankName                string  `json:"bankName"`
	BankVerifyDob           string  `json:"bankVerifyDob"`
	BankVerifyName          string  `json:"bankVerifyName"`
	BankVerifyPersonalId    string  `json:"bankVerifyPersonalid"`
	Capset                  float64 `json:"capset"`
	CreateAccountBankStatus string  `json:"createAccountBankStatus"`
	CreateAccountFaceStatus string  `json:"createAccountFaceStatus"`
	CreateAccountOcrStatus  string  `json:"createAccountOcrStatus"`
	CreateDate              string  `json:"createDate"`
	DateOfBirth             string  `json:"dateOfBirth"`
	DeviceName              string  `json:"deviceName"`
	Email                   string  `json:"email"`
	EmailBank               string  `json:"emailBank"`
	FaceMatching            float64 `json:"faceMatching"`
	FastLogin               bool    `json:"fastLogin"`
	Firmware                string  `json:"firmware"`
	Gender                  float64 `json:"gender"`
	GroupId                 string  `json:"groupId"`
	Hardware                string  `json:"hardware"`
	Identify                string  `json:"identify"`
	InvestCode              float64 `json:"investCode"`
	IsActived               bool    `json:"isActived"`
	IsEu                    bool    `json:"isEu"`
	IsNamed                 bool    `json:"isNamed"`
	IsOtpSuccessFirst       bool    `json:"isOtpSuccessFirst"`
	IsReged                 bool    `json:"isReged"`
	KeyPayLater             float64 `json:"keyPaylater"`
	KycDob                  string  `json:"kycDob"`
	KycGender               string  `json:"kycGender"`
	KycIssueDate            string  `json:"kycIssueDate"`
	KycIssuePlace           string  `json:"kycIssuePlace"`
	KycNationality          string  `json:"kycNationality"`
	LangCode                string  `json:"langCode"`
	LastImei                string  `json:"lastImei"`
	LastLogin               float64 `json:"lastLogin"`
	LastSessionTime         float64 `json:"lastSessionTime"`
	Manufacture             string  `json:"manufacture"`
	MmtAgent                string  `json:"mmtAgent"`
	Name                    string  `json:"name"`
	NameKyc                 string  `json:"nameKyc"`
	Nationality             string  `json:"nationality"`
	PassportKyc             string  `json:"passportKyc"`
	PersonalId              string  `json:"personalId"`
	PersonalIdKyc           string  `json:"personalIdKyc"`
	PhoneOs                 string  `json:"phoneOs"`
	Photo                   string  `json:"photo"`
	PushToken               string  `json:"pushToken"`
	RegisterDate            string  `json:"registerDate"`
	SwIsActive              float64 `json:"swIsActive"`
	UserId                  string  `json:"userId"`
	UserType                float64 `json:"userType"`
	ValidateEmail           string  `json:"validateEmail"`
	VerifyInfo              string  `json:"verifyInfo"`
	WalletName              string  `json:"walletName"`
	WalletStatus            string  `json:"walletStatus"`
}
