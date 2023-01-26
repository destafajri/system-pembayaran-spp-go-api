package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type SppService interface {
	CreateSpp(request *model.CreateSppRequest, timestamp time.Time) (*model.CreateSppResponse, error)
}