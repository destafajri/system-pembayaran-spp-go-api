package controller

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/user"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/user_repository"
	user_service "github.com/destafajri/system-pembayaran-spp-go-api/internal/service/user_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Controller(){
	// Setup Configuration
	configuration 	:= config.New()
	databasePostgre := config.NewPostgreDatabase(configuration)

	// Setup Repository
	userRepository := user_repository.NewUserRepository(databasePostgre)

	// Setup Service
	userService := user_service.NewUserService(&userRepository)

	// Setup Controller
	userController := user.NewUserController(&userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	userController.Route(app)

	// Start App
	err := app.Listen("0.0.0.0:9000")
	exception.PanicIfNeeded(err)
}