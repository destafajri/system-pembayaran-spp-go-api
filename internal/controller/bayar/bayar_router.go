package bayar

import "github.com/gofiber/fiber/v2"

func (controller *BayarController) Route(api fiber.Router) {
	api.Post("/spp/status/callback/paid", controller.CallbackPaid)
	api.Post("/spp/status/rollback/unpaid", controller.RollbackUnPaid)
}