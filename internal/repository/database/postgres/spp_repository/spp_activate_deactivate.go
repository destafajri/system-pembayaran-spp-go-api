package spp_repository

import (
	"errors"
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (spp *sppImplementation) ActivateSpp(spp_id string, timestamp time.Time) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		UPDATE spp SET
			is_active = COALESCE($1, is_active),
			updated_at = $2
		WHERE 
			id = $3
		`
	value := []interface{}{
		true,
		timestamp,
		spp_id,
	}

	_, err := spp.db.Exec(query, value...)
	if err != nil {
		log.Println(err)
		return errors.New("error activate spp")
	}

	return nil
}

func (spp *sppImplementation) DeactivateSpp(spp_id string, timestamp time.Time) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		UPDATE spp SET
			is_active = COALESCE($1, is_active),
			updated_at = $2
		WHERE 
			id = $3
		`
	value := []interface{}{
		false,
		timestamp,
		spp_id,
	}

	_, err := spp.db.Exec(query, value...)
	if err != nil {
		log.Println(err)
		return errors.New("error deactivate spp")
	}

	return nil
}
