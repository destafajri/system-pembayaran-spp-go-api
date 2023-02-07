package spp

import "github.com/gofiber/fiber/v2"

func (controller *SppController) Route(api fiber.Router) {
	api.Post("/spp/admin/create", controller.CreateSpp)
	api.Get("/spp/admin/list", controller.GetListSppAdmin)
	api.Get("/spp/admin/list/:siswa_id", controller.GetListSppBySiswaForAdmin)
	api.Get("/spp/list/siswa", controller.GetListSppBySiswaForSiswa)
}