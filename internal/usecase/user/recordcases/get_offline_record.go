package recordcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
)

type OfflineRecordOutput struct {
	ID         string
	UUID       string
	Name       string
	OnlineMode bool
	Registered bool
}

func OfflineRecordOutputFromRecord(record *entity.OfflineRecord) *OfflineRecordOutput {
	return &OfflineRecordOutput{
		ID:         record.ID,
		UUID:       record.UUID,
		Name:       record.Name,
		OnlineMode: record.OnlineMode,
		Registered: record.Registered,
	}
}

type GetOfflineRecordByIDUseCase struct {
	records gateway.OfflineRecordGateway
}

type GetOfflineRecordByNameUseCase struct {
	records gateway.OfflineRecordGateway
}

type GetOfflineRecordByUUIDUseCase struct {
	records gateway.OfflineRecordGateway
}

func NewGetOfflineRecordByNameUseCase(gat gateway.OfflineRecordGateway) *GetOfflineRecordByNameUseCase {
	return &GetOfflineRecordByNameUseCase{records: gat}
}

func NewGetOfflineRecordByUUIDUseCase(gat gateway.OfflineRecordGateway) *GetOfflineRecordByUUIDUseCase {
	return &GetOfflineRecordByUUIDUseCase{records: gat}
}

func NewGetOfflineRecordByIDUseCase(gat gateway.OfflineRecordGateway) *GetOfflineRecordByIDUseCase {
	return &GetOfflineRecordByIDUseCase{records: gat}
}

func (r GetOfflineRecordByNameUseCase) Execute(ctx context.Context, name string) (*OfflineRecordOutput, error) {
	record, err := r.records.FindByName(ctx, name)
	output := OfflineRecordOutputFromRecord(record)

	return output, err
}

func (r GetOfflineRecordByUUIDUseCase) Execute(ctx context.Context, uuid string) (*OfflineRecordOutput, error) {
	record, err := r.records.FindByUUID(ctx, uuid)
	output := OfflineRecordOutputFromRecord(record)

	return output, err
}

func (r GetOfflineRecordByIDUseCase) Execute(ctx context.Context, id string) (*OfflineRecordOutput, error) {
	record, err := r.records.FindById(ctx, id)
	output := OfflineRecordOutputFromRecord(record)

	return output, err
}
