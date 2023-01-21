package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type KelasService interface {
	CreateKelas(request *model.CreateKelasRequest, timestamp time.Time) (*model.CreateKelasResponse, error)
	GetListKelas(meta *meta.Metadata) ([]model.GetListKelasResponse, int, error)
	GetDetailKelas(kelas_id string) (*model.GetDetailKelasResponse, error)
	UpdateDetailKelas(kelas_id string, kelasReq *model.UpdateDetailKelasRequest) (*model.UpdateDetailKelasResponse, error)
	DeleteKelas(kelas_id string) error
}