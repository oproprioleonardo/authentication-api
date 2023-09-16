package gateway

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
)

type UserSessionGateway interface {
	Create(ctx context.Context, session *entity.UserSession) error
	FindBy(ctx context.Context, filter interface{}) (*entity.UserSession, error)
	FindById(ctx context.Context, id string) (*entity.UserSession, error)
	FindAll(ctx context.Context, filter interface{}, page int64, limit int64) ([]*entity.UserSession, error)
	Update(ctx context.Context, session *entity.UserSession) error
	DeleteBy(ctx context.Context, filter interface{}) error
	DeleteById(ctx context.Context, id string) error
}
