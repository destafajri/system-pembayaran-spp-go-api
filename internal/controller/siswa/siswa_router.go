package siswa

import "github.com/gofiber/fiber/v2"

func (controller *SiswaController) Route(api fiber.Router) {
	api.Post("/siswa/admin/create", controller.CreateSiswa)
}