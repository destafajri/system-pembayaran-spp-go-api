package service

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type ProductService interface {
	Create(request model.CreateProductRequest) (response model.CreateProductResponse)
	List() (responses []model.GetProductResponse)
}
