package user

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/helper/jwts"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{userService: *userService}
}

func (controller *UserController) CreateAdmin(c *fiber.Ctx) error {
	var (
		request       model.CreateAdminRequest
		apiKey        = c.Get("API-KEY")
		configuration = config.New()
	)

	if apiKey != configuration.Get("API-KEY") {
		return exception.ErrUnauthorized
	}

	err := c.BodyParser(&request)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, err := controller.userService.CreateAdmin(&request, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
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
		return exception.ErrorHandler(c, err)
	}

	response, err := controller.userService.Login(&request)
	if err != nil {
		if err.Error() == exception.ErrUnauthorized.Error() {
			return exception.ErrUnauthorized
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Login Success",
		Data:    response,
	})
}

func (controller *UserController) GetListUser(c *fiber.Ctx) error {
	var (
		metadata = meta.MetadataFromURL(c)
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	response, total, err := controller.userService.GetListUser(&metadata)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	metadata.Total = total
	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get List User Success",
		Meta:    metadata,
		Data:    response,
	})
}

func (controller *UserController) GetDetailUser(c *fiber.Ctx) error {
	var (
		user_id  = c.Params("user_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	response, err := controller.userService.GetDetailUser(user_id)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get Detail User Success",
		Data:    response,
	})
}

func (controller *UserController) ActivateUser(c *fiber.Ctx) error {
	var (
		user_id  = c.Params("user_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := controller.userService.ActivateUser(user_id, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Activate User Success",
	})
}

func (controller *UserController) DeactivateUser(c *fiber.Ctx) error {
	var (
		user_id  = c.Params("user_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := controller.userService.DeactivateUser(user_id, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Deactivate User Success",
	})
}
