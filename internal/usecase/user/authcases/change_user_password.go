package authcases

import (
	"context"
	"errors"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
	"github.com/skyepic/privateapi/pkg/security"
)

type ChangePasswordInput struct {
	ID          string
	OldPassword string
	NewPassword string
}

type ChangePasswordUseCase struct {
	authUserGateway gateway.AuthUserGateway
}

func NewChangePasswordUseCase(authUserGateway gateway.AuthUserGateway) ChangePasswordUseCase {
	return ChangePasswordUseCase{
		authUserGateway: authUserGateway,
	}
}

func (uc ChangePasswordUseCase) Execute(ctx context.Context, input *ChangePasswordInput) error {
	var (
		user           *entity.AuthUser
		hashedPassword string
		err            error
	)

	if user, err = uc.authUserGateway.FindById(ctx, input.ID); err != nil {
		return err
	}

	if security.Verify(input.OldPassword, user.Password) {
		if hashedPassword, err = security.Hash(input.NewPassword); err != nil {
			return err
		}

		user.Password = hashedPassword

		return uc.authUserGateway.Update(ctx, user)
	}

	return errors.New("O usuário " + input.ID + " inseriu a senha incorreta.")
}
