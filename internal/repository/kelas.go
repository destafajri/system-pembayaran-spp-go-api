package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type KelasRepository interface {
	CreateKelas(kelas *entity.KelasEntity) (*model.CreateKelasResponse, error)

	GetListKelas(meta *meta.Metadata) ([]model.GetListKelasResponse, int, error)
	GetDetailKelas(kelas_id string) (*model.GetDetailKelasResponse, error)

	UpdateDetailKelas(kelasReq *entity.KelasEntity) (*model.UpdateDetailKelasResponse, error)
	DeleteKelas(kelas_id string) error

	CekGuruExistByID(guru_id string) (bool, error)
	CekKelasExistByID(kelas_id string) (bool, error)
}
