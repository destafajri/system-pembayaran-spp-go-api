package bayar_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (bayar *bayarImplementation) GetSiswaIDByNIS(siswa_nis int) (string, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()
	
	var siswa_id string

	query := `SELECT id FROM siswa WHERE nis = $1 LIMIT 1`

	rows := bayar.db.QueryRow(query, siswa_nis)
	if err := rows.Scan(&siswa_id); err != nil {
		log.Println(err)
		return "", errors.New("siswa not found")
	}

	return siswa_id, nil
}