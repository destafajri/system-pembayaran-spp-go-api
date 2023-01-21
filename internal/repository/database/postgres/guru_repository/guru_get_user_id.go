package guru_repository

import (
	"log"

	"github.com/pkg/errors"
)

func (guru *guruImplementation) GetUserID(guru_id string) (string, error) {
	var user_id string

	query := `SELECT user_id from guru WHERE id = $1 LIMIT 1`

	rows := guru.db.QueryRow(query, guru_id)
	if err := rows.Scan(&user_id); err != nil {
		log.Println(err)
		return "", errors.Wrap(err, "getting user id")
	}

	return user_id, nil
}
