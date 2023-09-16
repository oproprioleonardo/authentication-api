package session

import (
	"github.com/skyepic/privateapi/internal/usecase/user/sessioncases"
)

func PresentUserSession(session *sessioncases.UserSessionOutput) Response {
	return Response{
		ID:          session.ID,
		ProfileId:   session.ProfileId,
		Active:      session.Active,
		LastServer:  session.LastServer,
		Ip:          session.Ip,
		FinalizedAt: session.FinalizedAt,
		StartedAt:   session.StartedAt,
	}
}
