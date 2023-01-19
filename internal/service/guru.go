package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type GuruService interface {
	CreateGuru(request *model.CreateGuruRequest, timestamp time.Time) (*model.CreateGuruResponse, error)
}