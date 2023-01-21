package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type SiswaService interface {
	CreateSiswa(request *model.CreateSiswaRequest, timestamp time.Time) (*model.CreateSiswaResponse, error)
	GetListSiswa(role string, meta *meta.Metadata) ([]model.GetListSiswaResponse, int, error)
}