package product

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository/product_repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service/product_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	productController.Route(app)
	return app
}

var configuration = config.New("../.env.test")

var database = config.NewMongoDatabase(configuration)
var productRepository = product_repository.NewProductRepository(database)
var productService = product_service.NewProductService(&productRepository)

var productController = NewProductController(&productService)

var app = createTestApp()
