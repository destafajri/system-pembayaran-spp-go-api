package spp_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (spp *sppImplementation) CekStatusBayar(spp_id string) (string, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var status string

	query := `SELECT ((CASE
				WHEN ((
					SELECT
						bayar.spp_id
					FROM
						bayar
					WHERE spp_id = $1
					LIMIT 1
					) IS NOT NULL )
					THEN 'paid'
				ELSE 'unpaid'
				END))
			FROM bayar`

	rows := spp.db.QueryRow(query, spp_id)
	if err := rows.Scan(&status); err != nil {
		log.Println(err)
		return status, errors.New("spp not found")
	}

	return status, nil
}
