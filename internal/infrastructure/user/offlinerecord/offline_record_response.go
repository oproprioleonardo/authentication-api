package offlinerecord

type Response struct {
	ID         string `json:"id"`
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	OnlineMode bool   `json:"onlineMode"`
	Registered bool   `json:"registered"`
}
