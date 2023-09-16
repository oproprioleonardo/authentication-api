package sessioncases

import "github.com/skyepic/privateapi/internal/domain/user/entity"

type UserSessionOutput struct {
	ID          string
	ProfileId   string
	Active      bool
	LastServer  string
	Ip          string
	FinalizedAt int64
	StartedAt   int64
}

func UserSessionOutputFromSession(session *entity.UserSession) *UserSessionOutput {
	return &UserSessionOutput{
		ID:          session.ID,
		ProfileId:   session.ProfileId,
		Active:      session.Active,
		LastServer:  session.LastServer,
		Ip:          session.Ip,
		FinalizedAt: session.FinalizedAt,
		StartedAt:   session.StartedAt,
	}
}
