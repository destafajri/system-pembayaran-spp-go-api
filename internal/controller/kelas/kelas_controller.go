package kelas

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

type KelasController struct {
	kelasService service.KelasService
}

func NewKelasController(kelasService *service.KelasService) KelasController {
	return KelasController{kelasService: *kelasService}
}

func (controller *KelasController) CreateKelas(c *fiber.Ctx) error {
	var (
		request  model.CreateKelasRequest
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

	err := c.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "errors",
			Message: err.Error(),
		})
	}

	response, err := controller.kelasService.CreateKelas(&request, time.Now())
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
		Message: "Create Kelas Success",
		Data:    response,
	})
}

func (controller *KelasController) GetListKelas(c *fiber.Ctx) error {
	var (
		metadata = meta.MetadataFromURL(c)
		_, err   = jwts.JWTAuthorizationHeader(c)
	)

	if err != nil {
		log.Println(err)
		return err
	}

	response, total, err := controller.kelasService.GetListKelas(&metadata)
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
		Message: "Get List Kelas Success",
		Meta:    metadata,
		Data:    response,
	})
}

func (controller *KelasController) GetDetailKelas(c *fiber.Ctx) error {
	var (
		kelas_id  = c.Params("kelas_id")
		_, err = jwts.JWTAuthorizationHeader(c)
	)

	if err != nil {
		log.Println(err)
		return err
	}

	response, err := controller.kelasService.GetDetailKelas(kelas_id)
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
		Message: "Get Detail Kelas Success",
		Data:    response,
	})
}