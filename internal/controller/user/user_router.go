package user

import "github.com/gofiber/fiber/v2"

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Get("/userinfo", controller.GetData)
	app.Post("/login", controller.Login)
}