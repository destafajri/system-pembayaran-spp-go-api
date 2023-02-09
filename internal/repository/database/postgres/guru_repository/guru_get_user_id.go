package guru_repository

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/pkg/errors"
)

func (guru *guruImplementation) GetUserID(guru_id string) (string, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()
	
	var user_id string

	query := `SELECT user_id from guru WHERE id = $1 LIMIT 1`

	rows := guru.db.QueryRow(query, guru_id)
	if err := rows.Scan(&user_id); err != nil {
		log.Println(err)
		return "", errors.New("user not found")
	}

	return user_id, nil
}
