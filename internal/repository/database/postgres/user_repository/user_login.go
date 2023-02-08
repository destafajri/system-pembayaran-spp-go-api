package user_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
)

func (user *userImplementation) Login(username string) (*entity.UserEntity, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()
	var resp entity.UserEntity

	query := `SELECT id, email, username, password, role FROM users WHERE username = $1 AND is_active is true`

	err := user.db.QueryRow(query, username).Scan(
		&resp.ID,
		&resp.Email,
		&resp.Username,
		&resp.Password,
		&resp.Role,
	)
	
	if err != nil {
		log.Println(err)
		return nil, errors.New("user not found")
	}

	return &resp, nil
}
