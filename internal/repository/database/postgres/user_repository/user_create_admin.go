package user_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
)

func (user *userImplementation) CreateAdmin(users *entity.UserEntity) (*entity.UserEntity, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `INSERT INTO users(
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
		users.ID,
		users.Email,
		users.Username,
		users.Password,
		users.Role,
		users.IsActive,
		users.CreatedAt,
		users.UpdatedAt,
	}

	_, err := user.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("user already exist")
	}

	resp := entity.UserEntity{
		ID:        users.ID,
		Email:     users.Email,
		Username:  users.Username,
		Role:      users.Role,
		IsActive:  users.IsActive,
		Timestamp: users.Timestamp,
	}

	return &resp, nil
}
