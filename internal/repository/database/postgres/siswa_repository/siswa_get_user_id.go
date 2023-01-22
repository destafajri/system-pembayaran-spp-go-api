package siswa_repository

import (
	"log"

	"github.com/pkg/errors"
)

func (siswa *siswaImplementation) GetUserID(siswa_id string) (string, error) {
	var user_id string

	query := `SELECT user_id from siswa WHERE id = $1 LIMIT 1`

	rows := siswa.db.QueryRow(query, siswa_id)
	if err := rows.Scan(&user_id); err != nil {
		log.Println(err)
		return "", errors.New("user not found")
	}

	return user_id, nil
}
