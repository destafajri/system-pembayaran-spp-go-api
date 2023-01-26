package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type SppRepository interface {
	CreateSpp(request *entity.SppEntity) (*model.CreateSppResponse, error)

	GetSiswaID(siswa_nis int) (string, error)
}
