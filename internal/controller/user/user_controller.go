package user

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/middlewares"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService *service.UserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller *UserController) CreateAdmin(c *fiber.Ctx) error {
	var request model.CreateAdminRequest

	err := c.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
	}

	response, err := controller.UserService.CreateAdmin(&request, time.Now())
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.WebResponse{
		Code:    201,
		Status:  "SUCCESS",
		Message: "Create Admin Success",
		Data:    response,
	})
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var request model.LoginRequest

	err := c.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
	}

	response, err := controller.UserService.Login(&request)
	if err != nil {
		if err.Error() == "unauthorized" {
			log.Println(err)
			return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  "errors",
				Message: err.Error(),
			})
		}

		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Login Success",
		Data:    response,
	})
}

func (controller *UserController) GetListUser(c *fiber.Ctx) error {
	var (
		metadata = meta.MetadataFromURL(c)
		token, _ = middlewares.JWTAuthorizationHeader(c)
	)

	claim, _ := middlewares.GetClaims(token)

	metadata.Total = 100
	return c.Status(fiber.StatusCreated).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get List User Success",
		Meta:    metadata,
		Data:    claim,
	})
}
