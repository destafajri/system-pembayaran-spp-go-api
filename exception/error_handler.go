package exception

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "BAD_REQUEST",
			Message:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
		Code:   fiber.StatusUnprocessableEntity,
		Status: "UnprocessableEntity",
		Message:   err.Error(),
	})
}