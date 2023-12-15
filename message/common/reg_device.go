package common

type RegDeviceMsg struct {
	Class       string `json:"_class"`
	CCode       string `json:"ccode"`
	CName       string `json:"cname"`
	Csp         string `json:"csp"`
	Device      string `json:"device"`
	DeviceOs    string `json:"device_os"`
	Firmware    string `json:"firmware"`
	Hardware    string `json:"hardware"`
	Icc         string `json:"icc"`
	Imei        string `json:"imei"`
	Manufacture string `json:"manufacture"`
	Mcc         string `json:"mcc"`
	Mnc         string `json:"mnc"`
	Number      string `json:"number"`
	SecureId    string `json:"secure_id"`
}
