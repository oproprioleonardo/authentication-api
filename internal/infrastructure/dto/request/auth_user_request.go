package request

type RegisterAuthUserRequest struct {
	UUID       string `json:"uuid" validate:"required"`
	Name       string `json:"name" validate:"required"`
	OnlineMode bool   `json:"onlineMode"`
	Password   string `json:"password"`
	IP         string `json:"ip" validate:"required"`
	Server     string `json:"server" validate:"required"`
}

type AuthenticateUserRequest struct {
	UUID       string `json:"uuid" validate:"required"`
	Name       string `json:"name" validate:"required"`
	OnlineMode bool   `json:"onlineMode"`
	Password   string `json:"password"`
	IP         string `json:"ip" validate:"required"`
	Server     string `json:"server" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
