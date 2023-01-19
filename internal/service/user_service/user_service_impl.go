package user_service

import (
	_ "strconv"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
)

type userServiceimpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) service.UserService {
	return &userServiceimpl{
		userRepository: *userRepository,
	}
}
