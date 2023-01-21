package guru

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *GuruController) Route(api fiber.Router) {
	api.Post("/guru/admin/create", controller.CreateGuru)

	api.Get("/guru/list", controller.GetListGuru)
	api.Get("/guru/detail/:guru_id", controller.GetDetailGuru)

	api.Put("/guru/admin/activate/:guru_id", controller.ActivateGuru)
	api.Put("/guru/admin/deactivate/:guru_id", controller.DeactivateGuru)
}