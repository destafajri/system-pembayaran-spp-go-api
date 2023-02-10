package repository

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type SppRepository interface {
	CreateSpp(request *entity.SppEntity) (*model.CreateSppResponse, error)

	GetListSpp(kelas string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error)
	GetListSppBySiswa(siswa_id string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error)
	GetDetailSpp(spp_id string) (*model.GetDetailSppResponse, error)

	ActivateSpp(spp_id string, timestamp time.Time) error
	DeactivateSpp(spp_id string, timestamp time.Time) error

	GetSiswaIDByNIS(siswa_nis int) (string, error)
	GetSiswaIDByUserID(user_id string) (string, error)
	CekSppExistByID(spp_id string) (bool, error)
	CekSppByIDAndSiswaIDIsMatch(spp_id, siswa_id string) (bool, error)
}
