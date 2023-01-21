package kelas

import "github.com/gofiber/fiber/v2"

func (controller *KelasController) Route(api fiber.Router) {
	api.Post("/kelas/admin/create", controller.CreateKelas)
	api.Get("/kelas/list", controller.GetListKelas)
	api.Get("/kelas/detail/:kelas_id", controller.GetDetailKelas)
}