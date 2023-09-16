package recordcases

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/domain/user/gateway"
	"github.com/skyepic/privateapi/internal/usecase/shared/pages"
)

type ListOfflineRecordOutput struct {
	ID         string
	UUID       string
	Name       string
	OnlineMode bool
	Registered bool
}

func ListOfflineRecordOutputFromRecord(record *entity.OfflineRecord) *ListOfflineRecordOutput {
	return &ListOfflineRecordOutput{
		ID:         record.ID,
		UUID:       record.UUID,
		Name:       record.Name,
		OnlineMode: record.OnlineMode,
		Registered: record.Registered,
	}
}

type ListOfflineRecordsUseCase struct {
	records gateway.OfflineRecordGateway
}

func NewListOfflineRecordsUseCase(gat gateway.OfflineRecordGateway) *ListOfflineRecordsUseCase {
	return &ListOfflineRecordsUseCase{records: gat}
}

func (r ListOfflineRecordsUseCase) Execute(ctx context.Context, terms interface{}, query pages.Query) (*pages.Pagination[*ListOfflineRecordOutput], error) {
	offRecords, err := r.records.FindAll(ctx, terms, query.Page, query.PerPage)
	output := make([]*ListOfflineRecordOutput, len(offRecords))

	for i, rec := range offRecords {
		output[i] = ListOfflineRecordOutputFromRecord(rec)
	}

	return &pages.Pagination[*ListOfflineRecordOutput]{
		CurrentPage: query.Page,
		PerPage:     query.PerPage,
		Total:       int64(len(output)),
		Items:       output,
	}, err
}
