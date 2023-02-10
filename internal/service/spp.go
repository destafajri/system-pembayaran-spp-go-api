package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type SppService interface {
	CreateSpp(request *model.CreateSppRequest, timestamp time.Time) (*model.CreateSppResponse, error)
	GetListSpp(kelasparam string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error)
	GetListSppBySiswaForAdmin(siswa_id string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error)
	GetListSppBySiswaForSiswa(user_id string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error)
	GetDetailSppForAdmin(spp_id string) (*model.GetDetailSppResponse, error)
	GetDetailSppForSiswa(spp_id, user_id string) (*model.GetDetailSppResponse, error)
	ActivateSpp(spp_id string, timestamp time.Time) error
	DeactivateSpp(spp_id string, timestamp time.Time) error
}