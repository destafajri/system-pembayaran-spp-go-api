package guru_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

func (guru *guruImplementation) CreateGuru(userReq *entity.UserEntity, guruReq *entity.GuruEntity) (*model.CreateGuruResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query1 := `INSERT INTO users(
			id,
			email,
			username,
			password,
			role,
			is_active,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	values := []interface{}{
		userReq.ID,
		userReq.Email,
		userReq.Username,
		userReq.Password,
		userReq.Role,
		userReq.IsActive,
		userReq.CreatedAt,
		userReq.UpdatedAt,
	}

	_, err := guru.db.Exec(query1, values...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("user already exist")
	}

	query2 := `INSERT INTO guru(
			id,
			user_id,
			nama,
			is_active
		)
		VALUES ($1, $2, $3, $4)
	`
	value := []interface{}{
		guruReq.ID,
		guruReq.UserID,
		guruReq.Nama,
		guruReq.IsActive,
	}

	_, err = guru.db.Exec(query2, value...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("guru already exist")
	}

	resp := model.CreateGuruResponse{
		ID:        guruReq.ID,
		UserID:    guruReq.UserID,
		Email:     userReq.Email,
		Username:  userReq.Username,
		Nama:      guruReq.Nama,
		Role:      userReq.Role,
		IsActive:  guruReq.IsActive,
		Timestamp: userReq.Timestamp,
	}

	return &resp, nil
}
