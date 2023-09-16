package webserver

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/skyepic/privateapi/internal/infrastructure/db"
	"github.com/skyepic/privateapi/internal/infrastructure/dto/response"
	"github.com/skyepic/privateapi/internal/infrastructure/user/auth"
	"github.com/skyepic/privateapi/internal/infrastructure/user/offlinerecord"
	"github.com/skyepic/privateapi/internal/infrastructure/user/profile"
	"github.com/skyepic/privateapi/internal/infrastructure/user/session"
	"github.com/skyepic/privateapi/internal/infrastructure/web/controller"
	"github.com/skyepic/privateapi/internal/usecase/user/authcases"
	"github.com/skyepic/privateapi/internal/usecase/user/recordcases"
	"github.com/skyepic/privateapi/pkg/database"
	"log"
	"os"
	"os/signal"
)

func Setup() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return ctx.Status(code).JSON(response.NewResponseError(ctx, err, code))
		},
	})

	app.Use(pprof.New(), logger.New())
	group := "/v1"
	route := app.Group(group)

	idGenerator := db.NewIdGenerator()
	recordMongoGateway := offlinerecord.NewOfflineRecordMongoGateway(database.DB.Collection("recordsUsers"))
	userSessionMongoGateway := session.NewUserSessionMongoGateway(database.DB.Collection("usersSessions"))
	userProfileMongoGateway := profile.NewUserProfileMongoGateway(database.DB.Collection("usersProfiles"))
	authUserMongoGateway := auth.NewAuthUserMongoGateway(database.DB.Collection("users"))

	createOfflineRecordUseCase := recordcases.NewCreateOfflineRecord(recordMongoGateway, idGenerator)
	getOfflineRecordByNameUseCase := recordcases.NewGetOfflineRecordByNameUseCase(recordMongoGateway)
	getOfflineRecordByUUIDUseCase := recordcases.NewGetOfflineRecordByUUIDUseCase(recordMongoGateway)
	getOfflineRecordByIDUseCase := recordcases.NewGetOfflineRecordByIDUseCase(recordMongoGateway)
	deleteOfflineRecordByIdUseCase := recordcases.NewDeleteOfflineRecordByIdUseCase(recordMongoGateway)
	updateOfflineRecordUseCase := recordcases.NewUpdateOfflineRecordUseCase(recordMongoGateway)
	listOfflineRecordsUseCase := recordcases.NewListOfflineRecordsUseCase(recordMongoGateway)

	registerUserUseCase := authcases.NewRegisterUserUseCase(authUserMongoGateway, userProfileMongoGateway, userSessionMongoGateway, recordMongoGateway, idGenerator)
	listAuthUsersUseCase := authcases.NewListAuthUsersUseCase(authUserMongoGateway, userSessionMongoGateway, userProfileMongoGateway)
	authenticateUserUseCase := authcases.NewAuthenticateUserUseCase(authUserMongoGateway, userProfileMongoGateway, userSessionMongoGateway, recordMongoGateway, idGenerator)
	changePasswordUseCase := authcases.NewChangePasswordUseCase(authUserMongoGateway)
	disconnectUserUseCase := authcases.NewDisconnectUserUseCase(authUserMongoGateway, userSessionMongoGateway)
	getAuthUserByIDUseCase := authcases.NewGetAuthUserByIDUseCase(authUserMongoGateway, userSessionMongoGateway, userProfileMongoGateway)
	getAuthUserByUUIDUseCase := authcases.NewGetAuthUserByUUIDUseCase(authUserMongoGateway, userSessionMongoGateway, userProfileMongoGateway)

	offlineRecordController := controller.NewOfflineRecordController(
		createOfflineRecordUseCase,
		getOfflineRecordByIDUseCase,
		getOfflineRecordByUUIDUseCase,
		getOfflineRecordByNameUseCase,
		listOfflineRecordsUseCase,
		updateOfflineRecordUseCase,
		deleteOfflineRecordByIdUseCase,
	)

	offlineMojangRecordGroup := route.Group("/offline/mojang")

	offlineMojangRecordGroup.Post("/record", offlineRecordController.Register)
	offlineMojangRecordGroup.Get("/record", offlineRecordController.GetOfflineUsers)
	offlineMojangRecordGroup.Get("/record/:name", offlineRecordController.GetOfflineUserByName)
	offlineMojangRecordGroup.Get("/record/:uuid/uuid", offlineRecordController.GetOfflineUser)
	offlineMojangRecordGroup.Put("/record/:id", offlineRecordController.UpdateOfflineUser)
	offlineMojangRecordGroup.Delete("/record/:id", offlineRecordController.DeleteOfflineUser)

	userController := controller.NewAuthUserController(
		authenticateUserUseCase,
		getAuthUserByUUIDUseCase,
		getAuthUserByIDUseCase,
		registerUserUseCase,
		changePasswordUseCase,
		listAuthUsersUseCase,
		disconnectUserUseCase,
	)

	authGroup := route.Group("/auth")

	authGroup.Post("/signup", userController.Register)
	authGroup.Post("/login", userController.Authenticate)
	authGroup.Post("/logout/:id", userController.Logout)
	authGroup.Post("/changePassword", userController.ChangePassword)
	authGroup.Get("/users", userController.GetAuthUsers)
	authGroup.Get("/users/:id", userController.GetAuthUser)
	authGroup.Get("/users/:uuid/uuid", userController.GetAuthUserByUUID)

	// not found routes
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"status":  fiber.StatusNotFound,
				"message": "endpoint is not found",
			})
		},
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()
	return app
}
