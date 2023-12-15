package model

import (
	"encoding/json"
	"momo-api/message/common"
	"momo-api/utils"
	"strings"
)

type User struct {
	Phone   string   `json:"phone,omitempty"`
	Name    string   `json:"name,omitempty"`
	Profile *Profile `json:"profile,omitempty"`
	Device  *Device  `json:"device,omitempty"`
}

func ParseUser(str string) (*User, error) {
	data := &User{}
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *User) ToJson() string {
	b, err := json.Marshal(u)
	if err != nil {
		return "{}"
	}
	return string(b)
}

func (u *User) Standard() bool {
	return u.Profile != nil && u.Device != nil
}

func (u *User) Equals(another *User) bool {
	return another != nil && u.Phone == another.Phone
}

func CreatePartner(phone string, name string) *User {
	return &User{
		Phone: strings.TrimSpace(phone),
		Name:  name,
	}
}

func CreateStandardUser(phone string, name string) *User {
	return &User{
		Phone: strings.TrimSpace(phone),
		Name:  name,
		Profile: &Profile{
			SessionKeyTracking: utils.GenerateSessionKeyTracking(),
		},
		Device: &Device{
			RegDeviceMsg: &common.RegDeviceMsg{
				Class:       "mservice.backend.entity.msg.RegDeviceMsg",
				CCode:       utils.CountryCode,
				CName:       utils.CountryName,
				Csp:         utils.ContentServiceProvider,
				Device:      utils.Device,
				DeviceOs:    utils.DeviceOS,
				Firmware:    utils.FirmwareVersion,
				Hardware:    utils.Hardware,
				Icc:         "",
				Imei:        utils.GenerateImei(),
				Manufacture: utils.Manufacture,
				Mcc:         utils.MobileCountryCode,
				Mnc:         utils.MobileNetworkCode,
				Number:      phone,
				SecureId:    "",
			},
		},
	}
}
