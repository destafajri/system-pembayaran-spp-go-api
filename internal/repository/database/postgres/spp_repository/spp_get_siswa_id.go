package spp_repository

import (
	"errors"
	"log"
)

func (spp *sppImplementation) GetSiswaID(siswa_nis int) (string, error) {
	var siswa_id string

	query := `SELECT id FROM siswa WHERE nis = $1`

	rows := spp.db.QueryRow(query, siswa_nis)
	if err := rows.Scan(&siswa_id); err != nil {
		log.Println(err)
		return "", errors.New("siswa not found")
	}

	return siswa_id, nil
}
