package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type GuruRepository interface {
	CreateGuru(req1 *entity.UserEntity, req2 *entity.GuruEntity) (*model.CreateGuruResponse, error)

	GetListGuruAdmin(meta *meta.Metadata) ([]model.GetListGuruResponse, int, error)
	GetDetailGuruAdmin(guru_id string) (*model.GetDetailGuruResponse, error)
	GetListGuruNonAdmin(meta *meta.Metadata) ([]model.GetListGuruResponse, int, error)
	GetDetailGuruNonAdmin(guru_id string) (*model.GetDetailGuruResponse, error)

	CekGuruExistByID(guru_id string) (bool, error)
}
