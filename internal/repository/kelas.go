package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type KelasRepository interface {
	CreateKelas(kelas *entity.KelasEntity) (*model.CreateKelasResponse, error)
	GetListKelas(meta *meta.Metadata) ([]model.GetListKelasResponse, int, error)
	GetDetailKelas(kelas_id string) (*model.GetDetailKelasResponse, error)

	CekGuruExistByID(guru_id string) (bool, error)
	CekKelasExistByID(kelas_id string) (bool, error)
}
