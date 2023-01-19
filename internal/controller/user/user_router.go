package user

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *UserController) Route(app *fiber.App, api fiber.Router) {
	app.Post("/admin/create", controller.CreateAdmin)
	app.Post("/login", controller.Login)

	api.Get("/user/admin/list", controller.GetListUser)
	api.Get("/user/admin/detail/:user_id", controller.GetDetailUser)
}