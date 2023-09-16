package entity

import (
	"errors"
	"time"
)

type UserProfile struct {
	ID         string
	UUID       string
	Name       string
	OnlineMode bool
	Email      string
	Phone      string
	CreatedAt  int64
}

func NewUserProfile(id string, uuid string, name string, onlineMode bool) (*UserProfile, error) {
	profile := &UserProfile{
		ID:         id,
		UUID:       uuid,
		Name:       name,
		OnlineMode: onlineMode,
		CreatedAt:  time.Now().UnixMilli(),
	}

	if err := profile.Validate(); err != nil {
		return nil, err
	}

	return profile, nil
}

func (p UserProfile) Validate() error {
	if p.ID == "" {
		return errors.New("id is empty")
	}
	if p.UUID == "" {
		return errors.New("uuid is empty")
	}
	if p.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}
