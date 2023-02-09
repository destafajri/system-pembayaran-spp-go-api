package bayar

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/helper/jwts"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

type BayarController struct {
	bayarService service.BayarService
}

func NewBayarController(bayarService *service.BayarService) BayarController {
	return BayarController{bayarService: *bayarService}
}

func (controller *BayarController) CallbackPaid(c *fiber.Ctx) error {
	var (
		request  model.BayarSppRequest
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

	response, err := controller.bayarService.PaidSpp(&request)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(responses.WebResponse{
		Code:    201,
		Status:  "SUCCESS",
		Message: "Paid Spp Success",
		Data:    response,
	})
}

func (controller *BayarController) RollbackUnPaid(c *fiber.Ctx) error {
	var (
		request  model.BayarSppRequest
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

	response, err := controller.bayarService.UnpaidSpp(&request)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Rollback Unpaid Spp Success",
		Data:    response,
	})
}
