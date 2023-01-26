package spp

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/helper/jwts"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

type SppController struct {
	sppService service.SppService
}

func NewSppController(sppService *service.SppService) SppController {
	return SppController{sppService: *sppService}
}

func (controller *SppController) CreateSpp(c *fiber.Ctx) error {
	var (
		request  model.CreateSppRequest
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

	response, err := controller.sppService.CreateSpp(&request, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(responses.WebResponse{
		Code:    201,
		Status:  "SUCCESS",
		Message: "Create Spp Success",
		Data:    response,
	})
}
