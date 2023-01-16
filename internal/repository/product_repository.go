package repository

import "github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"

type ProductRepository interface {
	Insert(product entity.Product)

	FindAll() (products []entity.Product)

	DeleteAll()
}
