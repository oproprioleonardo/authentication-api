package authcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
)

type DisconnectUserUseCase struct {
	authUserGateway    gateway.AuthUserGateway
	userSessionGateway gateway.UserSessionGateway
}

func NewDisconnectUserUseCase(authUserGateway gateway.AuthUserGateway, userSessionGateway gateway.UserSessionGateway) DisconnectUserUseCase {
	return DisconnectUserUseCase{
		authUserGateway:    authUserGateway,
		userSessionGateway: userSessionGateway,
	}
}

func (uc DisconnectUserUseCase) Execute(ctx context.Context, id string) error {
	var (
		user    *entity.AuthUser
		session *entity.UserSession
		err     error
	)

	if user, err = uc.authUserGateway.FindById(ctx, id); err != nil {
		return err
	}
	if session, err = uc.userSessionGateway.FindById(ctx, user.LastSessionID); err != nil {
		return err
	}

	if err := session.Disconnect(); err != nil {
		return err
	}

	return uc.userSessionGateway.Update(ctx, session)
}
