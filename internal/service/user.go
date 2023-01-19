package service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type UserService interface {
	CreateAdmin(request *model.CreateAdminRequest, timestamp time.Time) (*model.CreateAdminResponse, error)
	Login(request *model.LoginRequest) (*model.LoginResponse, error)
	GetListUser(meta *meta.Metadata) ([]model.GetListUserResponse, int, error)
	GetDetailUser(user_id string) (*model.GetDetailUser, error)
}
