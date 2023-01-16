package service

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

type UserService interface {
	Register(users *model.RegisterUserPayload) (*model.RegisterUserResponse, error)
	GetData(*model.GetUserPayload) (*model.GetUserResponse , error)
	Login(*model.LoginPayload) (*model.LoginResponse, error)
}

