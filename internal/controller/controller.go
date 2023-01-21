package controller

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/guru"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/kelas"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/user"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/middlewares"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/guru_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/kelas_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/user_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service/guru_service"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service/kelas_service"
	user_service "github.com/destafajri/system-pembayaran-spp-go-api/internal/service/user_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Controller() {
	// Setup Configuration
	configuration := config.New()
	databasePostgre := config.NewPostgreDatabase(configuration)

	// Setup Repository
	userRepository := user_repository.NewUserRepository(databasePostgre)
	guruRepository := guru_repository.NewGuruRepository(databasePostgre)
	kelasRepository := kelas_repository.NewkelasRepository(databasePostgre)

	// Setup Service
	userService := user_service.NewUserService(&userRepository)
	guruService := guru_service.NewUserService(&guruRepository)
	kelasService := kelas_service.NewkelasService(&kelasRepository)

	// Setup Controller
	userController := user.NewUserController(&userService)
	guruController := guru.NewGuruController(&guruService)
	kelasController := kelas.NewKelasController(&kelasService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Versioning Route
	api := app.Group("/api", middlewares.New(middlewares.Config{SigningKey: middlewares.JWT_SECRET_KEY}))

	// Setup Routing
	userController.Route(app, api)
	guruController.Route(api)
	kelasController.Route(api)

	// Start App
	err := app.Listen("0.0.0.0:9000")
	exception.PanicIfNeeded(err)
}
