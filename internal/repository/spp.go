package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type SppRepository interface {
	CreateSpp(request *entity.SppEntity) (*model.CreateSppResponse, error)

	GetListSppForAdmin(kelas string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error)
	GetListSppBySiswa(siswa_id string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error)

	GetSiswaIDByNIS(siswa_nis int) (string, error)
	GetSiswaIDByUserID(user_id string) (string, error)
}
