package user

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/admin/create", controller.CreateAdmin)
	app.Post("/login", controller.Login)

	// Setup Versioning Route
	api := app.Group("/api", middlewares.New(middlewares.Config{SigningKey: middlewares.JWT_SECRET_KEY}))

	api.Get("/user/admin/list", controller.GetListUser)
	// api.Get("/user/admin/detail/:user_id", controller)
}