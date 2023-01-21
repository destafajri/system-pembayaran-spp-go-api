package siswa

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

type SiswaController struct {
	siswaService service.SiswaService
}

func NewSiswaController(siswaService *service.SiswaService) SiswaController {
	return SiswaController{siswaService: *siswaService}
}

func (controller *SiswaController) CreateSiswa(c *fiber.Ctx) error {
	var (
		request  model.CreateSiswaRequest
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

	response, err := controller.siswaService.CreateSiswa(&request, time.Now())
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
		Message: "Create Siswa Success",
		Data:    response,
	})
}

func (controller *SiswaController) GetListSiswa(c *fiber.Ctx) error {
	var (
		metadata = meta.MetadataFromURL(c)
		token, _ = jwts.JWTAuthorizationHeader(c)
		claim, _ = jwts.GetClaims(token)
	)

	response, total, err := controller.siswaService.GetListSiswa(claim.Role, &metadata)
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
		Message: "Get List Siswa Success",
		Meta:    metadata,
		Data:    response,
	})
}
