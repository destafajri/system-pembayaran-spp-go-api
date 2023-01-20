package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type GuruService interface {
	CreateGuru(request *model.CreateGuruRequest, timestamp time.Time) (*model.CreateGuruResponse, error)
	GetListGuru(meta *meta.Metadata) ([]model.GetListGuruResponse, int, error)
}