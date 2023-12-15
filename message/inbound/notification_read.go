package inbound

import "encoding/json"

type NotificationReadResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ParseNotificationReadResponse(body []byte) *NotificationReadResponse {
	data := &NotificationReadResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	return data
}
