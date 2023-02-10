package spp_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (spp *sppImplementation) DeleteSpp(spp_id string) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		DELETE FROM spp WHERE id = $1
		`
	values := []interface{}{
		spp_id,
	}

	_, err := spp.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error delete spp on repository")
	}

	return nil
}

func (spp *sppImplementation) DeleteStatus(spp_id string) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		DELETE FROM bayar WHERE spp_id = $1
		`
	values := []interface{}{
		spp_id,
	}

	_, err := spp.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return errors.New("siswa haven't paid that spp")
	}

	return nil
}
