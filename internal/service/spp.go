package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type SppService interface {
	CreateSpp(request *model.CreateSppRequest, timestamp time.Time) (*model.CreateSppResponse, error)
	GetListSpp(meta *meta.Metadata) ([]model.GetListSppResponse, int, error)
}