package spp

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

func (controller *SppController) GetListSppAdmin(c *fiber.Ctx) error {
	var (
		metadata   = meta.MetadataFromURL(c)
		token, _   = jwts.JWTAuthorizationHeader(c)
		claim, _   = jwts.GetClaims(token)
		kelasparam = c.Query("kelas")
	)

	// claims
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	response, total, err := controller.sppService.GetListSpp(kelasparam, &metadata)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	metadata.Total = total
	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get List Spp Success",
		Meta:    metadata,
		Data:    response,
	})
}

func (controller *SppController) GetListSppBySiswaForAdmin(c *fiber.Ctx) error {
	var (
		metadata = meta.MetadataFromURL(c)
		token, _ = jwts.JWTAuthorizationHeader(c)
		claim, _ = jwts.GetClaims(token)
		siswa_id = c.Params("siswa_id")
	)

	// claims
	if claim.Role != "admin" && claim.Role != "guru" {
		return exception.ErrPermissionNotAllowed
	}

	response, total, err := controller.sppService.GetListSppBySiswaForAdmin(siswa_id, &metadata)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	metadata.Total = total
	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get List Spp By Siswa Success",
		Meta:    metadata,
		Data:    response,
	})
}

func (controller *SppController) GetListSppBySiswaForSiswa(c *fiber.Ctx) error {
	var (
		metadata = meta.MetadataFromURL(c)
		token, _ = jwts.JWTAuthorizationHeader(c)
		claim, _ = jwts.GetClaims(token)
	)

	// claims
	if claim.Role != "siswa" {
		return exception.ErrPermissionNotAllowed
	}

	response, total, err := controller.sppService.GetListSppBySiswaForSiswa(claim.ID, &metadata)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	metadata.Total = total
	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get List Spp By Siswa Success",
		Meta:    metadata,
		Data:    response,
	})
}
