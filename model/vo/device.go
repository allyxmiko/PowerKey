package vo

type DeviceVo struct {
	ID       uint   `json:"id"`
	IP       string `json:"ip"`
	Delay    int    `json:"delay"`
	Name     string `json:"name"`
	Topic    string `json:"topic"`
	Password string `json:"password"`
	Username string `json:"username"`
	Mac      string `json:"mac"`
}
