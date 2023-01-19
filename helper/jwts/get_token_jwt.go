package jwts

import (
	"strings"

	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

func JWTAuthorizationHeader(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse{
			Code:    fiber.StatusBadRequest,
			Status:  "Unauthorized",
			Message: "Missing JWT on the Header Request",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	return tokenString, nil
}
