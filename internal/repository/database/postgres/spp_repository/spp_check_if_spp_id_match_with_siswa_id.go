package spp_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (spp *sppImplementation) CekSppByIDAndSiswaIDIsMatch(spp_id, siswa_id string) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `SELECT id FROM spp WHERE id = $1 AND siswa_id = $2 LIMIT 1;`

	rows, err := spp.db.Query(query, spp_id, siswa_id)
	if err != nil {
		log.Println(err)
		return false, errors.New("spp not exist for this siswa")
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return rows.Next(), errors.New("spp not exist for this siswa")
}
