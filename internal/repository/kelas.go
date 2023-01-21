package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type KelasRepository interface {
	CreateKelas(kelas *entity.KelasEntity) (*model.CreateKelasResponse, error)

	CekGuruExistByID(guru_id string) (bool, error)
}
