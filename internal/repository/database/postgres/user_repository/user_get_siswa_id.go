package user_repository

import (
	"log"

	"github.com/pkg/errors"
)

func (user *userImplementation) GetSiswaID(user_id string) (string, error) {
	var siswa_id string

	query := `SELECT id FROM siswa WHERE user_id = $1`

	rows := user.db.QueryRow(query, user_id)
	if err := rows.Scan(&siswa_id); err != nil {
		log.Println(err)
		return "", errors.New("siswa not found")
	}

	return siswa_id, nil
}
