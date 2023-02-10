package user_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (user *userServiceimpl) GetListUser(meta *meta.Metadata) ([]model.GetListUserResponse, int, error){
	users, total, err := user.userRepository.GetListUser(meta)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return users, total, nil
}