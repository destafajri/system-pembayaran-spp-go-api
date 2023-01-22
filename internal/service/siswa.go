package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type SiswaService interface {
	CreateSiswa(request *model.CreateSiswaRequest, timestamp time.Time) (*model.CreateSiswaResponse, error)
	GetListSiswa(role string, meta *meta.Metadata) ([]model.GetListSiswaResponse, int, error)
	GetListSiswaByKelas(kelas_id string, meta *meta.Metadata) ([]model.GetListSiswaByKelasResponse, int, error)
	GetDetailSiswa(role, siswa_id string) (*model.GetDetailSiswaResponse, error)
	ActivateSiswa(siswa_id string, timestamp time.Time) error
	DeactivateSiswa(siswa_id string, timestamp time.Time) error
}