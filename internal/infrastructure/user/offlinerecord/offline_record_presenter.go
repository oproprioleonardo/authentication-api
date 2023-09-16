package offlinerecord

import (
	"github.com/skyepic/privateapi/internal/usecase/user/recordcases"
)

func PresentFromListOutput(user *recordcases.ListOfflineRecordOutput) Response {
	return Response{
		ID:         user.ID,
		UUID:       user.UUID,
		Name:       user.Name,
		OnlineMode: user.OnlineMode,
		Registered: user.Registered,
	}
}

func PresentFromOutput(user *recordcases.OfflineRecordOutput) Response {
	return Response{
		ID:         user.ID,
		UUID:       user.UUID,
		Name:       user.Name,
		OnlineMode: user.OnlineMode,
		Registered: user.Registered,
	}
}
