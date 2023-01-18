package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type UserService interface {
	CreateAdmin(request *model.CreateAdminRequest, timestamp time.Time) (*model.CreateAdminResponse, error)
	Login(request *model.LoginRequest) (*model.LoginResponse, error)
}
