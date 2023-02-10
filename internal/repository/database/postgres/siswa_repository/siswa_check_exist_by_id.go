package siswa_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (siswa *siswaImplementation) CekSiswaExistByID(id string) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `SELECT id FROM siswa WHERE id = $1 LIMIT 1;`

	rows, err := siswa.db.Query(query, id)
	if err != nil {
		log.Println(err)
		return false, errors.New("siswa not exist")
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return rows.Next(), errors.New("siswa not exist")
}
