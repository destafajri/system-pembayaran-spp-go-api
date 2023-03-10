package siswa

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
		return exception.ErrPermissionNotAllowed
	}

	err := c.BodyParser(&request)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, err := controller.siswaService.CreateSiswa(&request, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
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
		return exception.ErrorHandler(c, err)
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

func (controller *SiswaController) GetListSiswaByKelas(c *fiber.Ctx) error {
	var (
		metadata = meta.MetadataFromURL(c)
		kelas_id = c.Params("kelas_id")
		_, err   = jwts.JWTAuthorizationHeader(c)
	)

	if err != nil {
		return err
	}

	response, total, err := controller.siswaService.GetListSiswaByKelas(kelas_id, &metadata)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	metadata.Total = total
	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get List Siswa By Kelas Success",
		Meta:    metadata,
		Data:    response,
	})
}

func (controller *SiswaController) GetDetailSiswa(c *fiber.Ctx) error {
	var (
		siswa_id = c.Params("siswa_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
		claim, _ = jwts.GetClaims(token)
	)

	response, err := controller.siswaService.GetDetailSiswa(claim.Role, siswa_id)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Get Detail Siswa Success",
		Data:    response,
	})
}

func (controller *SiswaController) ActivateSiswa(c *fiber.Ctx) error {
	var (
		siswa_id = c.Params("siswa_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := controller.siswaService.ActivateSiswa(siswa_id, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Activate Siswa Success",
	})
}

func (controller *SiswaController) DeactivateSiswa(c *fiber.Ctx) error {
	var (
		siswa_id = c.Params("siswa_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := controller.siswaService.DeactivateSiswa(siswa_id, time.Now())
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Deactivate Siswa Success",
	})
}

func (controller *SiswaController) DeleteSiswa(c *fiber.Ctx) error {
	var (
		siswa_id = c.Params("siswa_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := controller.siswaService.DeleteSiswa(siswa_id)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Delete Siswa Success",
	})
}

func (controller *SiswaController) UpdateKelasSiswa(c *fiber.Ctx) error {
	var (
		input    model.UpdateKelasSiswaRequest
		siswa_id = c.Params("siswa_id")
		token, _ = jwts.JWTAuthorizationHeader(c)
	)

	// claims
	claim, _ := jwts.GetClaims(token)
	if claim.Role != "admin" {
		return exception.ErrPermissionNotAllowed
	}

	err := c.BodyParser(&input)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	resp, err := controller.siswaService.UpdateKelasSiswa(siswa_id, &input)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Update Kelas Siswa Success",
		Data:    resp,
	})
}
