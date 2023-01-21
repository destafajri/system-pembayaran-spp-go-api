package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type SiswaService interface {
	CreateSiswa(request *model.CreateSiswaRequest, timestamp time.Time) (*model.CreateSiswaResponse, error)
}