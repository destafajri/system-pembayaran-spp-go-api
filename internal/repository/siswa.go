package repository

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type SiswaRepository interface {
	CreateSiswa(user *entity.UserEntity, siswaReq *entity.SiswaEntity) (*model.CreateSiswaResponse, error)

	GetListSiswaForAdmin(meta *meta.Metadata) ([]model.GetListSiswaResponse, int, error)
	GetListSiswaNonAdmin(meta *meta.Metadata) ([]model.GetListSiswaResponse, int, error)
	GetListSiswaByKelas(kelas_id string, meta *meta.Metadata) ([]model.GetListSiswaByKelasResponse, int, error)
	GetDetailSiswaForAdmin(siswa_id string) (*model.GetDetailSiswaResponse, error)
	GetDetailSiswaNonAdmin(siswa_id string) (*model.GetDetailSiswaResponse, error)

	ActivateSiswa(user_id, siswa_id string, timestamp time.Time) error
	DeactivateSiswa(user_id, siswa_id string, timestamp time.Time) error

	DeleteSiswa(user_id, siswa_id string) error
	UpdateKelasSiswa(siswa_id, kelas_id, angkatan string) error 

	GetUserID(siswa_id string) (string, error)
	GetKelasID(kelas string) (string, error)
	CekNIS(nim int) (string, error)
	CekSiswaExistByID(siswa_id string) (bool, error)
}
