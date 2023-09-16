package entity

import "errors"

type AuthUser struct {
	ID                string
	UserProfileID     string
	LastSessionID     string
	AuthByLastSession bool
	Password          string
	Secret            string
}

func NewAuthUser(id string, profileId string, lastSessionId string, authByLastSession bool, password string, secret string) (*AuthUser, error) {
	user := &AuthUser{
		ID:                id,
		UserProfileID:     profileId,
		LastSessionID:     lastSessionId,
		AuthByLastSession: authByLastSession,
		Password:          password,
		Secret:            secret,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u AuthUser) Validate() error {
	if u.ID == "" {
		return errors.New("id is empty")
	}
	if u.UserProfileID == "" {
		return errors.New("profile is empty")
	}
	if u.LastSessionID == "" && u.AuthByLastSession {
		return errors.New("last session is empty")
	}

	return nil
}
