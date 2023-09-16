package auth

import "github.com/skyepic/privateapi/internal/infrastructure/user/session"

type Response struct {
	ID                string           `json:"id"`
	ProfileId         string           `json:"profileId"`
	UUID              string           `json:"uuid"`
	Name              string           `json:"name"`
	OnlineMode        bool             `json:"onlineMode"`
	Email             string           `json:"email"`
	Phone             string           `json:"phone"`
	CreatedAt         int64            `json:"createdAt"`
	AuthByLastSession bool             `json:"authByLastSession"`
	Secret            bool             `json:"secret"`
	LastSession       session.Response `json:"lastSession"`
}
