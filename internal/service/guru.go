package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type GuruService interface {
	CreateGuru(request *model.CreateGuruRequest, timestamp time.Time) (*model.CreateGuruResponse, error)
	GetListGuru(role string, meta *meta.Metadata) ([]model.GetListGuruResponse, int, error)
	GetDetailGuru(role, guru_id string) (*model.GetDetailGuruResponse, error)
	ActivateGuru(guru_id string, timestamp time.Time) error
	DeactivateGuru(guru_id string, timestamp time.Time) error
}