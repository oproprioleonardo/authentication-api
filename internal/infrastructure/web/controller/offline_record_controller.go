package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/skyepic/privateapi/internal/infrastructure/dto/query"
	"github.com/skyepic/privateapi/internal/infrastructure/dto/response"
	"github.com/skyepic/privateapi/internal/infrastructure/user/offlinerecord"
	"github.com/skyepic/privateapi/internal/usecase/shared/pages"
	"github.com/skyepic/privateapi/internal/usecase/user/recordcases"
	"math"
)

type (
	OfflineRecordController struct {
		CreateOfflineRecordUseCase     *recordcases.CreateOfflineRecordUseCase
		GetOfflineRecordByIDUseCase    *recordcases.GetOfflineRecordByIDUseCase
		GetOfflineRecordByUUIDUseCase  *recordcases.GetOfflineRecordByUUIDUseCase
		GetOfflineRecordByNameUseCase  *recordcases.GetOfflineRecordByNameUseCase
		ListOfflineRecordsUseCase      *recordcases.ListOfflineRecordsUseCase
		UpdateOfflineRecordUseCase     *recordcases.UpdateOfflineRecordUseCase
		DeleteOfflineRecordByIdUseCase *recordcases.DeleteOfflineRecordByIdUseCase
	}
)

func NewOfflineRecordController(
	CreateOfflineRecordUseCase *recordcases.CreateOfflineRecordUseCase,
	GetOfflineRecordByIDUseCase *recordcases.GetOfflineRecordByIDUseCase,
	GetOfflineRecordByUUIDUseCase *recordcases.GetOfflineRecordByUUIDUseCase,
	GetOfflineRecordByNameUseCase *recordcases.GetOfflineRecordByNameUseCase,
	ListOfflineRecordsUseCase *recordcases.ListOfflineRecordsUseCase,
	UpdateOfflineRecordUseCase *recordcases.UpdateOfflineRecordUseCase,
	DeleteOfflineRecordByIdUseCase *recordcases.DeleteOfflineRecordByIdUseCase,
) OfflineRecordController {
	return OfflineRecordController{
		CreateOfflineRecordUseCase:     CreateOfflineRecordUseCase,
		GetOfflineRecordByIDUseCase:    GetOfflineRecordByIDUseCase,
		GetOfflineRecordByUUIDUseCase:  GetOfflineRecordByUUIDUseCase,
		GetOfflineRecordByNameUseCase:  GetOfflineRecordByNameUseCase,
		ListOfflineRecordsUseCase:      ListOfflineRecordsUseCase,
		UpdateOfflineRecordUseCase:     UpdateOfflineRecordUseCase,
		DeleteOfflineRecordByIdUseCase: DeleteOfflineRecordByIdUseCase,
	}
}

func (ofc OfflineRecordController) Register(c *fiber.Ctx) error {
	var (
		body *offlinerecord.RegisterRequest
		err  error
	)

	if err = c.BodyParser(&body); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	input := recordcases.CreateOfflineRecordInput{
		UUID:       body.UUID,
		Name:       body.Name,
		OnlineMode: body.OnlineMode,
	}

	output, err := ofc.CreateOfflineRecordUseCase.Execute(c.Context(), input)

	if err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	return response.NewResponseCreatedWithURI(c, "/v1/offline/mojang/record/"+output.ID)
}

func (ofc OfflineRecordController) GetOfflineUsers(c *fiber.Ctx) error {
	var (
		queryParams       query.Params
		offlineUserParams offlinerecord.Search
		outputs           *pages.Pagination[*recordcases.ListOfflineRecordOutput]
		resp              pages.Pagination[offlinerecord.Response]
		err               error
	)

	if err = c.QueryParser(&queryParams); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	queryParams.PerPage = int64(math.Max(float64(queryParams.PerPage), 1))
	queryParams.Page = int64(math.Max(float64(queryParams.Page), 1))

	if err = c.QueryParser(&offlineUserParams); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	outputs, err = ofc.ListOfflineRecordsUseCase.Execute(c.Context(), offlineUserParams, queryParams.ToDomain())
	if err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	resp = pages.Pagination[offlinerecord.Response]{
		CurrentPage: outputs.CurrentPage,
		PerPage:     outputs.PerPage,
		Total:       outputs.Total,
		Items:       make([]offlinerecord.Response, len(outputs.Items)),
	}

	for i, output := range outputs.Items {
		resp.Items[i] = offlinerecord.PresentFromListOutput(output)
	}

	return response.NewResponseOKWithData(c, resp)
}

func (ofc OfflineRecordController) GetOfflineUser(c *fiber.Ctx) error {
	var (
		data *recordcases.OfflineRecordOutput
		resp offlinerecord.Response
		err  error
	)

	if data, err = ofc.GetOfflineRecordByUUIDUseCase.Execute(c.Context(), c.Params("uuid")); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	resp = offlinerecord.PresentFromOutput(data)

	return response.NewResponseOKWithData(c, resp)
}

func (ofc OfflineRecordController) GetOfflineUserByName(c *fiber.Ctx) error {
	var (
		data *recordcases.OfflineRecordOutput
		resp offlinerecord.Response
		err  error
	)

	if data, err = ofc.GetOfflineRecordByNameUseCase.Execute(c.Context(), c.Params("name")); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	resp = offlinerecord.PresentFromOutput(data)

	return response.NewResponseOKWithData(c, resp)
}

func (ofc OfflineRecordController) UpdateOfflineUser(c *fiber.Ctx) error {
	var (
		body   *offlinerecord.EditRequest
		input  recordcases.UpdateOfflineRecordInput
		output *recordcases.OfflineRecordOutput
		err    error
	)

	if err = c.BodyParser(&body); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	input = recordcases.UpdateOfflineRecordInput{
		ID:         c.Params("id"),
		UUID:       body.UUID,
		Name:       body.Name,
		OnlineMode: body.OnlineMode,
		Registered: body.Registered,
	}

	if output, err = ofc.UpdateOfflineRecordUseCase.Execute(c.Context(), input); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	} else {
		return response.NewResponseUpdatedWithData(c, offlinerecord.PresentFromOutput(output))
	}
}

func (ofc OfflineRecordController) DeleteOfflineUser(c *fiber.Ctx) error {
	if err := ofc.DeleteOfflineRecordByIdUseCase.Execute(c.Context(), c.Params("id")); err != nil {
		return response.NewResponseBadRequest(c, err)
	} else {
		return response.NewResponseNoContent(c)
	}
}
