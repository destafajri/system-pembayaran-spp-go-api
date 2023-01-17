package user

import "github.com/gofiber/fiber/v2"

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/admin/create", controller.CreateAdmin)
}
