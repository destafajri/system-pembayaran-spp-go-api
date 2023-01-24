package user_repository

import (
	"log"

	"github.com/pkg/errors"
)

func (siswa *userImplementation) GetSiswaID(user_id string) (string, error) {
	var siswa_id string

	query := `SELECT id from siswa WHERE user_id = $1 LIMIT 1`

	rows := siswa.db.QueryRow(query, siswa_id)
	if err := rows.Scan(&siswa_id); err != nil {
		log.Println(err)
		return "", errors.New("siswa not found")
	}

	return siswa_id, nil
}
