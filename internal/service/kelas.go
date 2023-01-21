package service

import "github.com/destafajri/system-pembayaran-spp-go-api/internal/model"

type KelasService interface {
	CreateKelas(request *model.CreateKelasRequest) (*model.CreateKelasResponse, error)
}