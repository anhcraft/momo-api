package model

type P2PTransaction struct {
	Id       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Message  string  `json:"message"`
	Sender   *User   `json:"sender"`
	Receiver *User   `json:"receiver"`
	Time     int64   `json:"time"`
}
