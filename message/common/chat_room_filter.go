package common

type ChatRoomFilter struct {
	RoomSources []string `json:"roomSources"`
	RoomTypes   []string `json:"roomTypes"`
	IsHide      bool     `json:"isHide"`
	IsBlock     bool     `json:"isBlock"`
}
