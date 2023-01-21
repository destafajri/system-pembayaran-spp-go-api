package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type SiswaRepository interface {
	CreateSiswa(user *entity.UserEntity, siswaReq *entity.SiswaEntity) (*model.CreateSiswaResponse, error)

	GetKelasID(kelas string) (string, error)
}