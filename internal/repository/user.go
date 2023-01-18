package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
)

type UserRepository interface {
	Login(username string) (*entity.UserEntity, error)
	CreateAdmin(users *entity.UserEntity) (*entity.UserEntity, error)
}