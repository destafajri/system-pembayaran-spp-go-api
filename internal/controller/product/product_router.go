package product

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func (controller *ProductController, ) Route(app *fiber.App) {
	// Setup Versioning Route
	api := app.Group("/api", middlewares.New(middlewares.Config{SigningKey: middlewares.JWT_SECRET_KEY}))

	api.Post("/products", controller.Create)
	api.Get("/products", controller.List)
}