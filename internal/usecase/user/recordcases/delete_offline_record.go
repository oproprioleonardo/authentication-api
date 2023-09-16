package recordcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
)

type DeleteOfflineRecordByIdUseCase struct {
	records gateway.OfflineRecordGateway
}

func NewDeleteOfflineRecordByIdUseCase(gat gateway.OfflineRecordGateway) *DeleteOfflineRecordByIdUseCase {
	return &DeleteOfflineRecordByIdUseCase{records: gat}
}

func (r DeleteOfflineRecordByIdUseCase) Execute(ctx context.Context, id string) error {
	return r.records.DeleteById(ctx, id)
}
