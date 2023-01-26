package spp

import "github.com/gofiber/fiber/v2"

func (controller *SppController) Route(api fiber.Router) {
	api.Post("/spp/admin/create", controller.CreateSpp)
}