package authcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/shared"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
	"github.com/skyepic/privateapi/internal/usecase/user/sessioncases"
	"github.com/skyepic/privateapi/pkg/security"
)

type RegisterAuthUserInput struct {
	UUID       string
	Name       string
	OnlineMode bool
	Password   string
	IP         string
	Server     string
}

type RegisterAuthUserOutput struct {
	ID                string
	ProfileId         string
	UUID              string
	Name              string
	Email             string
	Phone             string
	OnlineMode        bool
	CreatedAt         int64
	AuthByLastSession bool
	Secret            bool
	LastSession       *sessioncases.UserSessionOutput
}

func RegisterAuthUserOutputFrom(profile *entity.UserProfile, authUser *entity.AuthUser, session *entity.UserSession) *RegisterAuthUserOutput {
	return &RegisterAuthUserOutput{
		ID:                authUser.ID,
		ProfileId:         profile.ID,
		UUID:              profile.UUID,
		Name:              profile.Name,
		Email:             profile.Email,
		Phone:             profile.Phone,
		OnlineMode:        profile.OnlineMode,
		CreatedAt:         profile.CreatedAt,
		AuthByLastSession: authUser.AuthByLastSession,
		Secret:            authUser.Secret != "",
		LastSession:       sessioncases.UserSessionOutputFromSession(session),
	}
}

type RegisterUserUseCase struct {
	authUserGateway    gateway.AuthUserGateway
	profileUserGateway gateway.UserProfileGateway
	userSessionGateway gateway.UserSessionGateway
	offlineUserGateway gateway.OfflineRecordGateway
	idGenerator        shared.IDGenerator
}

func NewRegisterUserUseCase(authUserGateway gateway.AuthUserGateway, profileUserGateway gateway.UserProfileGateway, userSessionGateway gateway.UserSessionGateway, offlineRecordGateway gateway.OfflineRecordGateway, idGenerator shared.IDGenerator) RegisterUserUseCase {
	return RegisterUserUseCase{
		authUserGateway:    authUserGateway,
		profileUserGateway: profileUserGateway,
		userSessionGateway: userSessionGateway,
		offlineUserGateway: offlineRecordGateway,
		idGenerator:        idGenerator,
	}
}

func (u RegisterUserUseCase) Execute(ctx context.Context, input *RegisterAuthUserInput) (*RegisterAuthUserOutput, error) {
	var (
		err            error
		profileUser    *entity.UserProfile
		record         *entity.OfflineRecord
		authUser       *entity.AuthUser
		session        *entity.UserSession
		hashedPassword string
	)

	authUserId := u.idGenerator.Generate()
	userProfileId := u.idGenerator.Generate()
	userSessionId := u.idGenerator.Generate()

	// create profile
	if profileUser, err = entity.NewUserProfile(userProfileId, input.UUID, input.Name, input.OnlineMode); err != nil {
		return nil, err
	} else {
		if err = u.profileUserGateway.Create(ctx, profileUser); err != nil {
			return nil, err
		}
	}

	// update offline record
	if record, err = u.offlineUserGateway.FindByUUID(ctx, input.UUID); err != nil {
		return nil, err
	} else {
		record.Registered = true

		if err = u.offlineUserGateway.Update(ctx, record); err != nil {
			return nil, err
		}
	}

	// create session
	if session, err = entity.NewUserSession(userSessionId, userProfileId, input.Server, input.IP); err != nil {
		return nil, err
	} else {
		if err = u.userSessionGateway.Create(ctx, session); err != nil {
			return nil, err
		}
	}

	// create auth user

	// hash password
	if !input.OnlineMode && input.Password != "" {
		if hashedPassword, err = security.Hash(input.Password); err != nil {
			return nil, err
		}
	}

	// save authUser
	if authUser, err = entity.NewAuthUser(authUserId, userProfileId, userSessionId, true, hashedPassword, ""); err != nil {
		return nil, err
	} else {
		if err = u.authUserGateway.Create(ctx, authUser); err != nil {
			return nil, err
		}
	}

	return RegisterAuthUserOutputFrom(profileUser, authUser, session), nil
}
