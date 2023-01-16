package user

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService *service.UserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.RegisterUserPayload
	
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "errors",
			Data:   err.Error(),
		})
	}

	response, err := controller.UserService.Register(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "errors",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) GetData(c *fiber.Ctx) error {
	var request model.GetUserPayload
	
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "errors",
			Data:   err.Error(),
		})
	}

	response, err := controller.UserService.GetData(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "errors",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var request model.LoginPayload
	
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "errors",
			Data:   err.Error(),
		})
	}

	response, err := controller.UserService.Login(&request)
	if err != nil {
		if err.Error() == "You're Unauthorized" {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse{
				Code:   fiber.StatusUnauthorized,
				Status: "errors",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "errors",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}