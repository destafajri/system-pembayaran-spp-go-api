package exception

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

type Error string

const (
	ErrPermissionNotAllowed = Error("permission not allowed")
	ErrUnauthorized         = Error("you're unauthorized")
	ErrNotFound             = Error("not found")
	ErrInternal             = Error("internal server error")
	ErrBadRequest           = Error("bad request")
)

func (e Error) Error() string {
	return string(e)
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Println(err)

	_, ok := err.(ValidationError)
	if ok {
		log.Println(err)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "Error Validation",
			Message: err.Error(),
		})
	}

	if err == ErrPermissionNotAllowed {
		log.Println(err)
		return ctx.Status(fiber.StatusForbidden).JSON(responses.WebResponse{
			Code:    fiber.StatusForbidden,
			Status:  "Error Permission",
			Message: err.Error(),
		})
	}

	if err == ErrUnauthorized {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  "Unauthorizhed",
			Message: err.Error(),
		})
	}

	if err == ErrInternal {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(responses.WebResponse{
			Code:    fiber.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
	}

	if err == ErrBadRequest {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.WebResponse{
			Code:    fiber.StatusBadRequest,
			Status:  "Error Bad Request",
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
		Code:    fiber.StatusUnprocessableEntity,
		Status:  "errors",
		Message: err.Error(),
	})
}
