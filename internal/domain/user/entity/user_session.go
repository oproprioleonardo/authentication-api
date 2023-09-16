package entity

import (
	"errors"
	"github.com/skyepic/privateapi/pkg/config"
	"time"
)

type UserSession struct {
	ID          string
	ProfileId   string
	Active      bool
	LastServer  string
	Ip          string
	FinalizedAt int64
	StartedAt   int64
}

func NewUserSession(id string, profileId string, lastServer string, ip string) (*UserSession, error) {
	session := &UserSession{
		ID:          id,
		ProfileId:   profileId,
		Active:      true,
		LastServer:  lastServer,
		Ip:          ip,
		FinalizedAt: 0,
		StartedAt:   time.Now().UnixMilli(),
	}

	if err := session.Validate(); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *UserSession) Validate() error {
	if s.ID == "" {
		return errors.New("id is empty")
	}
	if s.ProfileId == "" {
		return errors.New("profileId is empty")
	}
	if s.LastServer == "" {
		return errors.New("lastServer is empty")
	}
	if s.Ip == "" {
		return errors.New("ip is empty")
	}
	return nil
}

func (s *UserSession) CanRecovery(ip string) bool {
	return s.Ip == ip && time.Now().UnixMilli() <= s.FinalizedAt+config.TimeToReopenSession
}

func (s *UserSession) Recovery(newServer string) {
	s.Active = true
	s.LastServer = newServer
	s.FinalizedAt = 0
}

func (s *UserSession) Disconnect() error {
	if s.Active {
		s.Active = false
		s.FinalizedAt = time.Now().UnixMilli()
		return nil
	}
	return errors.New("session is already disconnected")
}
