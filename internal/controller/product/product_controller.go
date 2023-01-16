package product

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
	"github.com/destafajri/system-pembayaran-spp-go-api/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.Create(request)
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	respons := controller.ProductService.List()
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   respons,
	})
}
