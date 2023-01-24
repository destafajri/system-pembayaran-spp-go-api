package user_repository

import (
	"log"

	"github.com/pkg/errors"
)

func (guru *userImplementation) GetGuruID(user_id string) (string, error) {
	var guru_id string

	query := `SELECT id from guru WHERE user_id = $1 LIMIT 1`

	rows := guru.db.QueryRow(query, guru_id)
	if err := rows.Scan(&guru_id); err != nil {
		log.Println(err)
		return "", errors.New("guru not found")
	}

	return guru_id, nil
}
