package user

import (
	"log"
	"time"

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

	response, err := controller.userService.CreateAdmin(&request, time.Now())
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

	response, err := controller.userService.Login(&request)
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
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  "Error",
			Message: "unauthorized",
			Error:   "unauthorized as admin",
		})
	}

	response, total, err := controller.userService.GetListUser(&metadata)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
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
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  "Error",
			Message: "unauthorized",
			Error:   "unauthorized as admin",
		})
	}

	response, err := controller.userService.GetDetailUser(user_id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
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
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  "Error",
			Message: "unauthorized",
			Error:   "unauthorized as admin",
		})
	}

	err := controller.userService.ActivateUser(user_id, time.Now())
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
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
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  "Error",
			Message: "unauthorized",
			Error:   "unauthorized as admin",
		})
	}

	err := controller.userService.DeactivateUser(user_id, time.Now())
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Deactivate User Success",
	})
}