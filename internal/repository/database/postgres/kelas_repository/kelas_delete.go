package kelas_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (user *kelasImplementation) DeleteKelas(kelas_id string) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		DELETE FROM kelas WHERE id = $1
		`
	values := []interface{}{
		kelas_id,
	}

	_, err := user.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error delete kelas")
	}

	return nil
}