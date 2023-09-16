package authcases

import (
	"context"
	"errors"
	"github.com/skyepic/privateapi/internal/domain/shared"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
	"github.com/skyepic/privateapi/internal/usecase/user/sessioncases"
	"github.com/skyepic/privateapi/pkg/config"
	"github.com/skyepic/privateapi/pkg/security"
	"time"
)

type AuthenticateUserInput struct {
	UUID       string
	Name       string
	OnlineMode bool
	Password   string
	IP         string
	Server     string
}

type AuthenticateUserOutput struct {
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

func AuthenticateUserOutputFrom(profile *entity.UserProfile, authUser *entity.AuthUser, session *entity.UserSession) *AuthenticateUserOutput {
	return &AuthenticateUserOutput{
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

type AuthenticateUserUseCase struct {
	authUserGateway    gateway.AuthUserGateway
	profileUserGateway gateway.UserProfileGateway
	userSessionGateway gateway.UserSessionGateway
	offlineUserGateway gateway.OfflineRecordGateway
	idGenerator        shared.IDGenerator
}

func NewAuthenticateUserUseCase(authUserGateway gateway.AuthUserGateway, profileUserGateway gateway.UserProfileGateway, userSessionGateway gateway.UserSessionGateway, offlineRecordGateway gateway.OfflineRecordGateway, idGenerator shared.IDGenerator) AuthenticateUserUseCase {
	return AuthenticateUserUseCase{
		authUserGateway:    authUserGateway,
		profileUserGateway: profileUserGateway,
		userSessionGateway: userSessionGateway,
		offlineUserGateway: offlineRecordGateway,
		idGenerator:        idGenerator,
	}
}

func (auc AuthenticateUserUseCase) Execute(ctx context.Context, input *AuthenticateUserInput) (*AuthenticateUserOutput, error) {
	var (
		authUser    *entity.AuthUser
		profile     *entity.UserProfile
		lastSession *entity.UserSession
		err         error
	)

	if profile, err = auc.profileUserGateway.FindByUUID(ctx, input.UUID); err != nil {
		return nil, errors.New("profile fetched by uuid " + input.UUID + " not found")
	}

	if authUser, err = auc.authUserGateway.FindByProfileId(ctx, profile.ID); err != nil {
		return nil, errors.New("user fetched by profileId " + authUser.UserProfileID + " not found")
	}

	// recupera a sessão anterior ou desconecta caso não tenha desconectado da última vez que o jogador saiu
	if authUser.LastSessionID != "" {
		if lastSession, err = auc.userSessionGateway.FindById(ctx, authUser.LastSessionID); err != nil {
			return nil, errors.New("last session fetched by id " + authUser.LastSessionID + " not found")
		}

		if err = lastSession.Disconnect(); err != nil && lastSession.Ip == input.IP && time.Now().UnixMilli() <= lastSession.FinalizedAt+config.TimeToReopenSession {
			lastSession.Recovery(input.Server)
		}

		if err = auc.userSessionGateway.Update(ctx, lastSession); err != nil {
			return nil, errors.New("broken collection")
		}
	}

	reuseSession := authUser.LastSessionID != "" && lastSession.Active && authUser.AuthByLastSession

	// tenta validar a autenticação do usuário
	if input.OnlineMode || reuseSession || security.Verify(input.Password, authUser.Password) {
		if !reuseSession {
			lastSession = &entity.UserSession{
				ID:          auc.idGenerator.Generate(),
				ProfileId:   profile.ID,
				Active:      true,
				LastServer:  input.Server,
				Ip:          input.IP,
				FinalizedAt: 0,
				StartedAt:   time.Now().UnixMilli(),
			}

			if err = auc.userSessionGateway.Create(ctx, lastSession); err != nil {
				return nil, err
			}
			authUser.LastSessionID = lastSession.ID

			if err = auc.authUserGateway.Update(ctx, authUser); err != nil {
				return nil, err
			}

		}

		profile.Name = input.Name

		if err = auc.profileUserGateway.Update(ctx, profile); err != nil {
			return nil, err
		}

		return AuthenticateUserOutputFrom(profile, authUser, lastSession), err
	}

	return nil, errors.New("Não foi possível autenticar o usuário " + input.UUID + ":" + input.Name)
}
