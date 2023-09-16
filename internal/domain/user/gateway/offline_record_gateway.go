package gateway

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
)

type OfflineRecordGateway interface {
	Create(ctx context.Context, record *entity.OfflineRecord) error
	FindBy(ctx context.Context, filter interface{}) (*entity.OfflineRecord, error)
	FindByName(ctx context.Context, name string) (*entity.OfflineRecord, error)
	FindById(ctx context.Context, id string) (*entity.OfflineRecord, error)
	FindByUUID(ctx context.Context, uuid string) (*entity.OfflineRecord, error)
	FindAll(ctx context.Context, filter interface{}, page int64, perPage int64) ([]*entity.OfflineRecord, error)
	Update(ctx context.Context, record *entity.OfflineRecord) error
	DeleteBy(ctx context.Context, filter interface{}) error
	DeleteById(ctx context.Context, id string) error
	DeleteByUUID(ctx context.Context, uuid string) error
}
