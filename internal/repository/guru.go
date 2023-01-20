package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type GuruRepository interface {
	CreateGuru(req1 *entity.UserEntity, req2 *entity.GuruEntity) (*model.CreateGuruResponse, error)

	GetListGuru(meta *meta.Metadata) ([]model.GetListGuruResponse, int, error)
}
