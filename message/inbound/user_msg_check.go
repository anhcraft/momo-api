package inbound

import "encoding/json"

type UserMsgCheckResponse struct {
	ErrorCode int                        `json:"errorCode"`
	ErrorDesc string                     `json:"errorDesc"`
	Extra     *UserMsgCheckResponseExtra `json:"extra"`
}

type UserMsgCheckResponseExtra struct {
	UserType       string `json:"userType"`
	IsVoice        string `json:"isVoice"`
	IsChangeDevice string `json:"isChangeDevice"`
	IsEu           string `json:"IS_EU"`
	IdentityKey    string `json:"IDENTITY_KEY"`
}

func (r *UserMsgCheckResponse) Success() bool {
	return r.ErrorCode == 0
}

func ParseUserMsgCheckResponse(body []byte) *UserMsgCheckResponse {
	data := &UserMsgCheckResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}
