package siswa_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (siswa *siswaImplementation) DeleteSiswa(user_id, siswa_id string) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query1 := `
		DELETE FROM siswa WHERE id = $1
		`
	values := []interface{}{
		siswa_id,
	}

	_, err := siswa.db.Exec(query1, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error delete siswa")
	}

	query2 := `
		DELETE FROM users WHERE id = $1
		`
	value := []interface{}{
		user_id,
	}

	_, err = siswa.db.Exec(query2, value...)
	if err != nil {
		log.Println(err)
		return errors.New("error delete user")
	}

	return nil
}