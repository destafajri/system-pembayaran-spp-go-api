package kelas

import "github.com/gofiber/fiber/v2"

func (controller *KelasController) Route(api fiber.Router) {
	api.Post("/kelas/admin/create", controller.CreateKelas)
	api.Get("/kelas/list", controller.GetListKelas)
	api.Get("/kelas/detail/:kelas_id", controller.GetDetailKelas)
	api.Put("/kelas/admin/update/:kelas_id", controller.UpdateDetailKelas)
	api.Delete("/kelas/admin/delete/:kelas_id", controller.DeleteKelas)
}