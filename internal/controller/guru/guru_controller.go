package guru

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/helper/jwts"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

type GuruController struct {
	guruService service.GuruService
}

func NewGuruController(guruService *service.GuruService) GuruController {
	return GuruController{guruService: *guruService}
}

func (controller *GuruController) CreateGuru(c *fiber.Ctx) error {
	var (
		request  model.CreateGuruRequest
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims to check role
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := c.BodyParser(&request)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, err := controller.guruService.CreateGuru(&request, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(responses.WebResponse{
		Code:    201,
		Status:  "SUCCESS",
		Message: "Create Guru Success",
		Data:    response,
	})
}

func (controller *GuruController) GetListGuru(c *fiber.Ctx) error {
	var (
		metadata = meta.MetadataFromURL(c)
		token, _ = jwts.JWTAuthorizationHeader(c)
		claim, _ = jwts.GetClaims(token)
	)

	response, total, err := controller.guruService.GetListGuru(claim.Role, &metadata)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	metadata.Total = total
	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get List Guru Success",
		Meta:    metadata,
		Data:    response,
	})
}

func (controller *GuruController) GetDetailGuru(c *fiber.Ctx) error {
	var (
		guru_id  = c.Params("guru_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
		claim, _ = jwts.GetClaims(token)
	)

	response, err := controller.guruService.GetDetailGuru(claim.Role, guru_id)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get Detail Guru Success",
		Data:    response,
	})
}

func (controller *GuruController) ActivateGuru(c *fiber.Ctx) error {
	var (
		guru_id  = c.Params("guru_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := controller.guruService.ActivateGuru(guru_id, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Activate Guru Success",
	})
}

func (controller *GuruController) DeactivateGuru(c *fiber.Ctx) error {
	var (
		guru_id  = c.Params("guru_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := controller.guruService.DeactivateGuru(guru_id, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Deactivate Guru Success",
	})
}
