package authcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
	"github.com/skyepic/privateapi/internal/usecase/shared/pages"
	"github.com/skyepic/privateapi/internal/usecase/user/sessioncases"
)

type ListUserOutput struct {
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

func ListUserOutputFrom(profile *entity.UserProfile, authUser *entity.AuthUser, session *entity.UserSession) *ListUserOutput {
	return &ListUserOutput{
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

type ListAuthUsersUseCase struct {
	authUserGateway    gateway.AuthUserGateway
	userSessionGateway gateway.UserSessionGateway
	userProfileGateway gateway.UserProfileGateway
}

func NewListAuthUsersUseCase(authUserGateway gateway.AuthUserGateway, userSessionGateway gateway.UserSessionGateway, userProfileGateway gateway.UserProfileGateway) ListAuthUsersUseCase {
	return ListAuthUsersUseCase{
		authUserGateway:    authUserGateway,
		userSessionGateway: userSessionGateway,
		userProfileGateway: userProfileGateway,
	}
}

func (r ListAuthUsersUseCase) Execute(ctx context.Context, terms interface{}, query pages.Query) (*pages.Pagination[*ListUserOutput], error) {
	users, err := r.authUserGateway.FindAll(ctx, terms, query.Page, query.PerPage)
	output := make([]*ListUserOutput, len(users))

	for i, user := range users {
		var (
			profile *entity.UserProfile
			session *entity.UserSession
		)

		if profile, err = r.userProfileGateway.FindById(ctx, user.UserProfileID); err != nil {
			return nil, err
		}

		if session, err = r.userSessionGateway.FindById(ctx, user.LastSessionID); err != nil && user.LastSessionID != "" {
			return nil, err
		}

		output[i] = ListUserOutputFrom(profile, user, session)
	}

	return &pages.Pagination[*ListUserOutput]{
		CurrentPage: query.Page,
		PerPage:     query.PerPage,
		Total:       int64(len(output)),
		Items:       output,
	}, err
}
