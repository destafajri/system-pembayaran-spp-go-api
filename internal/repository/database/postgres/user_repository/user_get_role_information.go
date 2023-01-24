package user_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (user *userImplementation) GetRoleInformation(id string) (string, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var role string

	query := `SELECT role FROM users WHERE id = $1 AND deleted_at IS NULL LIMIT 1;`

	rows := user.db.QueryRow(query, id)
	if err := rows.Scan(&role); err != nil {
		log.Println(err)
		return "", errors.New("user not found")
	}

	return role, nil
}
