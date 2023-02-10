package spp

import "github.com/gofiber/fiber/v2"

func (controller *SppController) Route(api fiber.Router) {
	api.Post("/spp/admin/create", controller.CreateSpp)
	api.Get("/spp/admin/list", controller.GetListSppAdmin)
	api.Get("/spp/admin/list/:siswa_id", controller.GetListSppBySiswaForAdmin)
	api.Get("/spp/list/siswa", controller.GetListSppBySiswaForSiswa)
	api.Get("/spp/admin/detail/:spp_id", controller.GetDetailSppForAdmin)
	api.Get("/spp/siswa/detail/:spp_id", controller.GetDetailSppForSiswa)
	api.Put("/spp/admin/activate/:spp_id", controller.ActivateSpp)
	api.Put("/spp/admin/deactivate/:spp_id", controller.DeactivateSpp)
	api.Delete("/spp/admin/delete/:spp_id", controller.DeleteSpp)
}