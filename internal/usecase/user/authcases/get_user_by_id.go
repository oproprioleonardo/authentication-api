package authcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
	"github.com/skyepic/privateapi/internal/usecase/user/sessioncases"
)

type UserOutput struct {
	ID                string
	ProfileId         string
	UUID              string
	Name              string
	OnlineMode        bool
	Email             string
	Phone             string
	CreatedAt         int64
	AuthByLastSession bool
	Secret            bool
	LastSession       *sessioncases.UserSessionOutput
}

func UserOutputFrom(profile *entity.UserProfile, authUser *entity.AuthUser, session *entity.UserSession) *UserOutput {
	return &UserOutput{
		ID:                authUser.ID,
		ProfileId:         profile.ID,
		UUID:              profile.UUID,
		Name:              profile.Name,
		OnlineMode:        profile.OnlineMode,
		Email:             profile.Email,
		Phone:             profile.Phone,
		CreatedAt:         profile.CreatedAt,
		AuthByLastSession: authUser.AuthByLastSession,
		Secret:            authUser.Secret != "",
		LastSession:       sessioncases.UserSessionOutputFromSession(session),
	}
}

type GetAuthUserByIDUseCase struct {
	authUserGateway    gateway.AuthUserGateway
	userSessionGateway gateway.UserSessionGateway
	userProfileGateway gateway.UserProfileGateway
}

func NewGetAuthUserByIDUseCase(authUserGateway gateway.AuthUserGateway, userSessionGateway gateway.UserSessionGateway, userProfileGateway gateway.UserProfileGateway) GetAuthUserByIDUseCase {
	return GetAuthUserByIDUseCase{
		authUserGateway:    authUserGateway,
		userSessionGateway: userSessionGateway,
		userProfileGateway: userProfileGateway,
	}
}

func (uc GetAuthUserByIDUseCase) Execute(ctx context.Context, id string) (*UserOutput, error) {
	var (
		authUser *entity.AuthUser
		profile  *entity.UserProfile
		session  *entity.UserSession
		err      error
	)

	if authUser, err = uc.authUserGateway.FindById(ctx, id); err != nil {
		return nil, err
	}
	if profile, err = uc.userProfileGateway.FindById(ctx, authUser.UserProfileID); err != nil {
		return nil, err
	}
	if session, err = uc.userSessionGateway.FindById(ctx, authUser.LastSessionID); err != nil && authUser.LastSessionID != "" {
		return nil, err
	}

	return UserOutputFrom(profile, authUser, session), nil
}
