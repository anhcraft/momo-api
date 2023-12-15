package inbound

import "encoding/json"

type VerifyOtpResponse struct {
	ErrorCode int                     `json:"errorCode"`
	ErrorDesc string                  `json:"errorDesc"`
	Extra     *VerifyOtpExtraResponse `json:"extra"`
}

type VerifyOtpExtraResponse struct {
	SetupKey string `json:"setupKey"`
}

func (r *VerifyOtpResponse) Success() bool {
	return r.ErrorCode == 0
}

func ParseVerifyOtpResponse(body []byte) *VerifyOtpResponse {
	data := &VerifyOtpResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}
