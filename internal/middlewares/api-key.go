package middlewares

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

func APIKey(c *fiber.Ctx) error {
	var (
		apiKey        = c.Get("API-KEY")
		configuration = config.New()
	)

	if apiKey == "" {
		return c.Status(fiber.StatusForbidden).JSON(responses.WebResponse{
			Code:    fiber.StatusForbidden,
			Status:  "Permission",
			Message: "Missing API-KEY on the header",
		})
	}

	if apiKey != configuration.Get("API-KEY") {
		return exception.ErrPermissionNotAllowed
	}

	return nil
}
