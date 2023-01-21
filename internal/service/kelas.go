package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type KelasService interface {
	CreateKelas(request *model.CreateKelasRequest, timestamp time.Time) (*model.CreateKelasResponse, error)
	GetListKelas(meta *meta.Metadata) ([]model.GetListKelasResponse, int, error)
}