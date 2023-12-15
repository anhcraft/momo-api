package inbound

import "encoding/json"

type GenericResponse struct {
	ErrorCode int    `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}

func (r *GenericResponse) Success() bool {
	return r.ErrorCode == 0
}

func ParseGenericResponse(body []byte) *GenericResponse {
	data := &GenericResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}
