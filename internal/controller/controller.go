package controller

import (
	"fmt"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/bayar"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/guru"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/kelas"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/siswa"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/spp"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/user"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/middlewares"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/bayar_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/guru_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/kelas_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/siswa_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/spp_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/database/postgres/user_repository"
	bayar_service "github.com/destafajri/system-pembayaran-spp-go-api/internal/service/bayar_ervice"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service/guru_service"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service/kelas_service"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service/siswa_service"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service/spp_service"
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
	siswaRepository := siswa_repository.NewSiswaRepository(databasePostgre)
	sppRepository := spp_repository.NewSppRepository(databasePostgre)
	bayarRepository := bayar_repository.NewBayarRepository(databasePostgre)

	// Setup Service
	userService := user_service.NewUserService(&userRepository)
	guruService := guru_service.NewUserService(&guruRepository)
	kelasService := kelas_service.NewkelasService(&kelasRepository)
	siswaService := siswa_service.NewSiswaService(&siswaRepository)
	sppService := spp_service.NewSppService(&sppRepository)
	bayarService := bayar_service.NewBayarService(&bayarRepository)

	// Setup Controller
	userController := user.NewUserController(&userService)
	guruController := guru.NewGuruController(&guruService)
	kelasController := kelas.NewKelasController(&kelasService)
	siswaController := siswa.NewSiswaController(&siswaService)
	sppController := spp.NewSppController(&sppService)
	bayarController := bayar.NewBayarController(&bayarService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Versioning Route
	api := app.Group("/api", middlewares.New(middlewares.Config{SigningKey: middlewares.JWT_SECRET_KEY}))

	// Setup Routing
	userController.Route(app, api)
	guruController.Route(api)
	kelasController.Route(api)
	siswaController.Route(api)
	sppController.Route(api)
	bayarController.Route(api)

	// Start App
	port := fmt.Sprintf("0.0.0.0:%s", configuration.Get("PORT"))
	err := app.Listen(port)
	exception.PanicIfNeeded(err)
}
