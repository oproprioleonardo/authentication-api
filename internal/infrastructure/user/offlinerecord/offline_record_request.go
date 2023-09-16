package offlinerecord

type RegisterRequest struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	OnlineMode bool   `json:"onlineMode"`
}

type EditRequest struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	OnlineMode bool   `json:"onlineMode"`
	Registered bool   `json:"registered"`
}
