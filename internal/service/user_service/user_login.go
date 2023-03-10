package user_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/helper/jwts"
	"github.com/destafajri/system-pembayaran-spp-go-api/helper/password"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
)

func (user *userServiceimpl) Login(request *model.LoginRequest) (*model.LoginResponse, error) {
	userInfo, err := user.userRepository.Login(request.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = password.CompareHashPassword(request.Password, userInfo.Password)
	if err != nil {
		return nil, exception.ErrUnauthorized
	}

	//Generate JWT
	token, err := jwts.GenerateJwtToken(userInfo.ID, userInfo.Username, userInfo.Role)
	if err != nil {
		return nil, err
	}

	resp := model.LoginResponse{
		Token: token,
	}

	return &resp, nil
}
