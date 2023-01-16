package controller

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	product "github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/product"
	user "github.com/destafajri/system-pembayaran-spp-go-api/internal/controller/user"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/product_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/user_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service/product_service"
	user_service "github.com/destafajri/system-pembayaran-spp-go-api/internal/service/user_service_impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Controller(){
	// Setup Configuration
	configuration 	:= config.New()
	databaseMongo 	:= config.NewMongoDatabase(configuration)
	databasePostgre := config.NewPostgreDatabase(configuration)

	// Setup Repository
	productRepository := product_repository.NewProductRepository(databaseMongo)
	userRepository := user_repository.NewUserRepository(databasePostgre)

	// Setup Service
	productService := product_service.NewProductService(&productRepository)
	userService := user_service.NewUserService(&userRepository)

	// Setup Controller
	productController := product.NewProductController(&productService)
	userController := user.NewUserController(&userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	productController.Route(app)
	userController.Route(app)

	// Start App
	err := app.Listen("0.0.0.0:9000")
	exception.PanicIfNeeded(err)
}