package dto

type DeviceDto struct {
	IP       string `validate:"ip"`
	Delay    int    `validate:"gte=0"`
	Name     string
	Topic    string `validate:"alphanum"`
	Password string
	Username string
	Mac      string `validate:"mac"`
}
