package guru

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/helper/jwts"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
)

type GuruController struct {
	GuruService service.GuruService
}

func NewGuruController(GuruService *service.GuruService) GuruController {
	return GuruController{GuruService: *GuruService}
}

func (controller *GuruController) CreateGuru(c *fiber.Ctx) error {
	var (
		request  model.CreateGuruRequest
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims to check role
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return c.Status(fiber.StatusCreated).JSON(responses.WebResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  "Error",
			Message: "unauthorized",
			Error:   "unauthorized as admin",
		})
	}

	err := c.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
	}

	response, err := controller.GuruService.CreateGuru(&request, time.Now())
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
		Message: "Create Guru Success",
		Data:    response,
	})
}
