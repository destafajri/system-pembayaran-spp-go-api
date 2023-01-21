package kelas_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (guru *kelasImplementation) CekKelasExistByID(kelas_id string) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `SELECT id FROM kelas WHERE id = $1 LIMIT 1;`

	rows, err := guru.db.Query(query, kelas_id)
	if err != nil {
		log.Println(err)
		return false, errors.New("kelas not exist")
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return rows.Next(), errors.New("kelas not exist")
}
