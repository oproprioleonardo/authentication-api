package request

type RegisterUserProfileRequest struct {
	UUID       string `json:"uuid,omitempty" validate:"required"`
	Name       string `json:"name,omitempty" validate:"required"`
	OnlineMode bool   `json:"onlineMode,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
}

type EditUserProfile struct {
	UUID       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	OnlineMode bool   `json:"onlineMode,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
}
