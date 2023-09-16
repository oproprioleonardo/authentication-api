package auth

import (
	"github.com/skyepic/privateapi/internal/infrastructure/user/session"
	"github.com/skyepic/privateapi/internal/usecase/user/authcases"
)

func PresentAuthUserFromRegister(output *authcases.RegisterAuthUserOutput) Response {
	return Response{
		ID:                output.ID,
		ProfileId:         output.ProfileId,
		UUID:              output.UUID,
		Name:              output.Name,
		OnlineMode:        output.OnlineMode,
		Email:             output.Email,
		Phone:             output.Phone,
		CreatedAt:         output.CreatedAt,
		AuthByLastSession: output.AuthByLastSession,
		Secret:            output.Secret,
		LastSession:       session.PresentUserSession(output.LastSession),
	}
}

func PresentAuthUserFromAuth(output *authcases.AuthenticateUserOutput) Response {
	return Response{
		ID:                output.ID,
		ProfileId:         output.ProfileId,
		UUID:              output.UUID,
		Name:              output.Name,
		OnlineMode:        output.OnlineMode,
		Email:             output.Email,
		Phone:             output.Phone,
		CreatedAt:         output.CreatedAt,
		AuthByLastSession: output.AuthByLastSession,
		Secret:            output.Secret,
		LastSession:       session.PresentUserSession(output.LastSession),
	}
}

func PresentListOutputToResponse(o *authcases.ListUserOutput) Response {
	return Response{
		ID:                o.ID,
		ProfileId:         o.ProfileId,
		UUID:              o.UUID,
		Name:              o.Name,
		OnlineMode:        o.OnlineMode,
		Email:             o.Email,
		Phone:             o.Phone,
		CreatedAt:         o.CreatedAt,
		AuthByLastSession: o.AuthByLastSession,
		Secret:            o.Secret,
		LastSession:       session.PresentUserSession(o.LastSession),
	}
}

func PresentOutputToResponse(o *authcases.UserOutput) Response {
	return Response{
		ID:                o.ID,
		ProfileId:         o.ProfileId,
		UUID:              o.UUID,
		Name:              o.Name,
		OnlineMode:        o.OnlineMode,
		Email:             o.Email,
		Phone:             o.Phone,
		CreatedAt:         o.CreatedAt,
		AuthByLastSession: o.AuthByLastSession,
		Secret:            o.Secret,
		LastSession:       session.PresentUserSession(o.LastSession),
	}
}
