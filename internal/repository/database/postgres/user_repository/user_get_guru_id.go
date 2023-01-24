package user_repository

import (
	"log"

	"github.com/pkg/errors"
)

func (user *userImplementation) GetGuruID(user_id string) (string, error) {
	var guru_id string

	query := `SELECT id FROM guru WHERE user_id = $1`

	rows := user.db.QueryRow(query, user_id)
	if err := rows.Scan(&guru_id); err != nil {
		log.Println(err)
		return "", errors.New("guru not found")
	}

	return guru_id, nil
}
