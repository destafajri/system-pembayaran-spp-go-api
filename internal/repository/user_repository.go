package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type UserRepository interface {
	// Register(users *entity.UserEntity) error
	// GetData(phone string) (*entity.UserEntity, error)
	// Login(phone string) (*entity.UserEntity, error)
	CreateAdmin(users *entity.UserEntity) (*model.CreateAdminResponse, error)
}