package guru

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *GuruController) Route(api fiber.Router) {
	api.Post("/guru/admin/create", controller.CreateGuru)

	api.Get("/guru/list", controller.GetListGuru)
}