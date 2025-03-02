package dto

type DeviceDto struct {
	IP       string `json:"ip" validate:"ip,omitempty"`
	Delay    int    `json:"delay" validate:"gte=0,omitempty"`
	Name     string `json:"name"`
	Topic    string `json:"topic" validate:"alphanum,omitempty"`
	Password string `json:"password"`
	Username string `json:"username"`
	Mac      string `json:"mac" validate:"mac,omitempty"`
}
