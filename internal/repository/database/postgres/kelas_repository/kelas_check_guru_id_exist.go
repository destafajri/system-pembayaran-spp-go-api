package kelas_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (guru *kelasImplementation) CekGuruExistByID(guru_id string) (bool, error){
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `SELECT id FROM guru WHERE id = $1 LIMIT 1;`

	rows, err := guru.db.Query(query, guru_id)
	if err != nil {
		log.Println(err)
		return false, errors.New("guru not exist")
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return rows.Next(), errors.New("guru not exist")
}