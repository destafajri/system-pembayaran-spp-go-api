package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type SiswaRepository interface {
	CreateSiswa(user *entity.UserEntity, siswaReq *entity.SiswaEntity) (*model.CreateSiswaResponse, error)

	GetListSiswaForAdmin(meta *meta.Metadata) ([]model.GetListSiswaResponse, int, error)
	GetListSiswaNonAdmin(meta *meta.Metadata) ([]model.GetListSiswaResponse, int, error)

	GetKelasID(kelas string) (string, error)
	CekNIS(nim int) (string, error)
}