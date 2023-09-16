package recordcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
)

type UpdateOfflineRecordInput struct {
	ID         string
	UUID       string
	Name       string
	OnlineMode bool
	Registered bool
}

type UpdateOfflineRecordUseCase struct {
	records gateway.OfflineRecordGateway
}

func NewUpdateOfflineRecordUseCase(gat gateway.OfflineRecordGateway) *UpdateOfflineRecordUseCase {
	return &UpdateOfflineRecordUseCase{gat}
}

func (r UpdateOfflineRecordUseCase) Execute(ctx context.Context, input UpdateOfflineRecordInput) (*OfflineRecordOutput, error) {
	record, err := r.records.FindById(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	record.UUID = input.UUID
	record.Name = input.Name
	record.OnlineMode = input.OnlineMode
	record.Registered = input.Registered

	if err := record.Validate(); err != nil {
		return nil, err
	}

	err = r.records.Update(ctx, record)
	if err != nil {
		return nil, err
	}

	return &OfflineRecordOutput{
		ID:         record.ID,
		UUID:       record.UUID,
		Name:       record.Name,
		OnlineMode: record.OnlineMode,
		Registered: record.Registered,
	}, nil
}
