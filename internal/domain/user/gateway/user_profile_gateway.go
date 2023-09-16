package gateway

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
)

type UserProfileGateway interface {
	Create(ctx context.Context, userProfile *entity.UserProfile) error
	FindBy(ctx context.Context, filter interface{}) (*entity.UserProfile, error)
	FindById(ctx context.Context, id string) (*entity.UserProfile, error)
	FindByName(ctx context.Context, name string) (*entity.UserProfile, error)
	FindByUUID(ctx context.Context, uuid string) (*entity.UserProfile, error)
	FindAll(ctx context.Context, filter interface{}, page int64, limit int64) ([]*entity.UserProfile, error)
	Update(ctx context.Context, userProfile *entity.UserProfile) error
	DeleteBy(ctx context.Context, filter interface{}) error
	DeleteById(ctx context.Context, id string) error
}
