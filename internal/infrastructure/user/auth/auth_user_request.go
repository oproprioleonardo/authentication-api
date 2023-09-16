package auth

type RegisterAuthUserRequest struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	OnlineMode bool   `json:"onlineMode"`
	Password   string `json:"password"`
	IP         string `json:"ip"`
	Server     string `json:"server"`
}

type AuthenticateUserRequest struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	OnlineMode bool   `json:"onlineMode"`
	Password   string `json:"password"`
	IP         string `json:"ip"`
	Server     string `json:"server"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
