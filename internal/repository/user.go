package repository

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type UserRepository interface {
	Login(username string) (*entity.UserEntity, error)
	CreateAdmin(users *entity.UserEntity) (*entity.UserEntity, error)

	GetListUser(meta *meta.Metadata) ([]model.GetListUserResponse, int, error)
	GetDetailUser(id string) (*model.GetDetailUser, error)

	ActivateUser(id string, timestamp time.Time) error
	DeactivateUser(id string, timestamp time.Time) error
	
	CekUserExistByID(id string) (bool, error)
}
