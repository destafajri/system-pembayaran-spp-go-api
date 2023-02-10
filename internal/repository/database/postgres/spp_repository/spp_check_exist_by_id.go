package spp_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (spp *sppImplementation) CekSppExistByID(spp_id string) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `SELECT id FROM spp WHERE id = $1 LIMIT 1;`

	rows, err := spp.db.Query(query, spp_id)
	if err != nil {
		log.Println(err)
		return false, errors.New("spp not exist")
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return rows.Next(), errors.New("spp not exist")
}
