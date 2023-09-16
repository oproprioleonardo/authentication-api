package recordcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/shared"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
)

type CreateOfflineRecordInput struct {
	UUID       string
	Name       string
	OnlineMode bool
}

type CreateOfflineRecordOutput struct {
	ID string
}

type CreateOfflineRecordUseCase struct {
	records     gateway.OfflineRecordGateway
	idGenerator shared.IDGenerator
}

func NewCreateOfflineRecord(gat gateway.OfflineRecordGateway, generator shared.IDGenerator) *CreateOfflineRecordUseCase {
	return &CreateOfflineRecordUseCase{records: gat, idGenerator: generator}
}

func (r CreateOfflineRecordUseCase) Execute(ctx context.Context, input CreateOfflineRecordInput) (*CreateOfflineRecordOutput, error) {
	record, err := entity.NewOfflineRecord(r.idGenerator.Generate(), input.UUID, input.Name, input.OnlineMode)
	if err != nil {
		return nil, err
	}
	return &CreateOfflineRecordOutput{ID: record.ID}, r.records.Create(ctx, record)
}
