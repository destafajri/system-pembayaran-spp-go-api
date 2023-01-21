package repository

import (
	"time"

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

	ActivateGuru(user_id, guru_id string, timestamp time.Time) error
	DeactivateGuru(user_id, guru_id string, timestamp time.Time) error

	CekGuruExistByID(guru_id string) (bool, error)
	GetUserID(guru_id string) (string, error)
}
