package exception

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Println(err)

	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "Error Validation",
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(responses.WebResponse{
		Code:    fiber.StatusBadRequest,
		Status:  "Error Bad Request",
		Message: err.Error(),
	})
}
