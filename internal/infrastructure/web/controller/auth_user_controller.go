package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/skyepic/privateapi/internal/infrastructure/dto/query"
	"github.com/skyepic/privateapi/internal/infrastructure/dto/request"
	"github.com/skyepic/privateapi/internal/infrastructure/dto/response"
	"github.com/skyepic/privateapi/internal/infrastructure/user/auth"
	"github.com/skyepic/privateapi/internal/usecase/shared/pages"
	"github.com/skyepic/privateapi/internal/usecase/user/authcases"
	"math"
)

type (
	AuthUserController struct {
		authUserUseCase       authcases.AuthenticateUserUseCase
		getUserByUUIDUseCase  authcases.GetAuthUserByUUIDUseCase
		getUserByIDUseCase    authcases.GetAuthUserByIDUseCase
		registerUserUseCase   authcases.RegisterUserUseCase
		changePasswordUseCase authcases.ChangePasswordUseCase
		listUsersUseCase      authcases.ListAuthUsersUseCase
		disconnectUserUseCase authcases.DisconnectUserUseCase
	}
)

func NewAuthUserController(
	authUserUseCase authcases.AuthenticateUserUseCase,
	getUserByUUIDUseCase authcases.GetAuthUserByUUIDUseCase,
	getUserByIDUseCase authcases.GetAuthUserByIDUseCase,
	registerUserUseCase authcases.RegisterUserUseCase,
	changePasswordUseCase authcases.ChangePasswordUseCase,
	listUsersUseCase authcases.ListAuthUsersUseCase,
	disconnectUserUseCase authcases.DisconnectUserUseCase,
) AuthUserController {
	return AuthUserController{
		authUserUseCase:       authUserUseCase,
		getUserByUUIDUseCase:  getUserByUUIDUseCase,
		getUserByIDUseCase:    getUserByIDUseCase,
		registerUserUseCase:   registerUserUseCase,
		changePasswordUseCase: changePasswordUseCase,
		listUsersUseCase:      listUsersUseCase,
		disconnectUserUseCase: disconnectUserUseCase,
	}
}

func (auc AuthUserController) Register(c *fiber.Ctx) error {
	var (
		body *request.RegisterAuthUserRequest
		resp auth.Response
	)

	if err := c.BodyParser(&body); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	input := &authcases.RegisterAuthUserInput{
		UUID:       body.UUID,
		Name:       body.Name,
		OnlineMode: body.OnlineMode,
		Password:   body.Password,
		IP:         body.IP,
		Server:     body.Server,
	}

	if userOutput, err := auc.registerUserUseCase.Execute(c.Context(), input); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	} else {
		resp = auth.PresentAuthUserFromRegister(userOutput)
		return response.NewResponseCreatedWithData(c, resp)
	}
}

func (auc AuthUserController) Authenticate(c *fiber.Ctx) error {
	var (
		body *request.AuthenticateUserRequest
		resp auth.Response
		err  error
	)

	if err = c.BodyParser(&body); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	input := &authcases.AuthenticateUserInput{
		UUID:       body.UUID,
		Password:   body.Password,
		OnlineMode: body.OnlineMode,
		IP:         body.IP,
		Server:     body.Server,
	}

	if userOutput, err := auc.authUserUseCase.Execute(c.Context(), input); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	} else {
		resp = auth.PresentAuthUserFromAuth(userOutput)
		return response.NewResponseOKWithData(c, resp)
	}
}

func (auc AuthUserController) ChangePassword(c *fiber.Ctx) error {
	var (
		body *request.ChangePasswordRequest
		err  error
	)

	if err = c.BodyParser(&body); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	input := &authcases.ChangePasswordInput{
		ID:          c.Params("id"),
		NewPassword: body.NewPassword,
		OldPassword: body.OldPassword,
	}

	if err = auc.changePasswordUseCase.Execute(c.Context(), input); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	} else {
		return response.NewResponseUpdated(c)
	}
}

func (auc AuthUserController) GetAuthUsers(c *fiber.Ctx) error {
	var (
		queryParams query.Params
		search      auth.Search
		err         error
	)
	if err = c.QueryParser(&queryParams); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	queryParams.PerPage = int64(math.Max(float64(queryParams.PerPage), 1))
	queryParams.Page = int64(math.Max(float64(queryParams.Page), 1))

	if err = c.QueryParser(&search); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	if paginationOutput, err := auc.listUsersUseCase.Execute(c.Context(), search, queryParams.ToDomain()); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	} else {
		resp := pages.Map(paginationOutput, auth.PresentListOutputToResponse)
		return response.NewResponseOKWithData(c, resp)
	}
}

func (auc AuthUserController) GetAuthUser(c *fiber.Ctx) error {

	if output, err := auc.getUserByIDUseCase.Execute(c.Context(), c.Params("id")); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	} else {
		resp := auth.PresentOutputToResponse(output)
		return response.NewResponseOKWithData(c, resp)
	}
}

func (auc AuthUserController) GetAuthUserByUUID(c *fiber.Ctx) error {
	if output, err := auc.getUserByUUIDUseCase.Execute(c.Context(), c.Params("uuid")); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	} else {
		resp := auth.PresentOutputToResponse(output)
		return response.NewResponseOKWithData(c, resp)
	}
}

func (auc AuthUserController) Logout(c *fiber.Ctx) error {
	if err := auc.disconnectUserUseCase.Execute(c.Context(), c.Params("id")); err != nil {
		return response.NewResponseBadRequest(c, err.Error())
	}

	return response.NewResponseOK(c, "session disconnected")
}
