package gateway

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
)

type AuthUserGateway interface {
	Create(ctx context.Context, authUser *entity.AuthUser) error
	FindBy(ctx context.Context, filter interface{}) (*entity.AuthUser, error)
	FindById(ctx context.Context, id string) (*entity.AuthUser, error)
	FindByProfileId(ctx context.Context, id string) (*entity.AuthUser, error)
	FindAll(ctx context.Context, filter interface{}, page int64, limit int64) ([]*entity.AuthUser, error)
	Update(ctx context.Context, authUser *entity.AuthUser) error
	DeleteBy(ctx context.Context, filter interface{}) error
	DeleteById(ctx context.Context, id string) error
	DeleteByUUID(ctx context.Context, uuid string) error
}
