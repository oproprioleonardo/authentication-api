package profile

type Response struct {
	ID         string `json:"id"`
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	OnlineMode bool   `json:"onlineMode"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	CreatedAt  int64  `json:"createdAt" `
}
