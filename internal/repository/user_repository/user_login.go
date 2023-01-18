package user_repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
)

func (user *userImplementation) Login(phone string) (*entity.UserEntity, error){
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var resp entity.UserEntity



	return &resp, nil
}