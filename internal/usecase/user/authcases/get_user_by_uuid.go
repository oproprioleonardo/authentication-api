package authcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
)

type GetAuthUserByUUIDUseCase struct {
	authUserGateway    gateway.AuthUserGateway
	userSessionGateway gateway.UserSessionGateway
	userProfileGateway gateway.UserProfileGateway
}

func NewGetAuthUserByUUIDUseCase(authUserGateway gateway.AuthUserGateway, userSessionGateway gateway.UserSessionGateway, userProfileGateway gateway.UserProfileGateway) GetAuthUserByUUIDUseCase {
	return GetAuthUserByUUIDUseCase{
		authUserGateway:    authUserGateway,
		userSessionGateway: userSessionGateway,
		userProfileGateway: userProfileGateway,
	}
}

func (uc GetAuthUserByUUIDUseCase) Execute(ctx context.Context, uuid string) (*UserOutput, error) {
	var (
		authUser *entity.AuthUser
		profile  *entity.UserProfile
		session  *entity.UserSession
		err      error
	)

	if profile, err = uc.userProfileGateway.FindByUUID(ctx, uuid); err != nil {
		return nil, err
	}
	if authUser, err = uc.authUserGateway.FindByProfileId(ctx, profile.ID); err != nil {
		return nil, err
	}
	if session, err = uc.userSessionGateway.FindById(ctx, authUser.LastSessionID); err != nil && authUser.LastSessionID != "" {
		return nil, err
	}

	return UserOutputFrom(profile, authUser, session), nil
}
