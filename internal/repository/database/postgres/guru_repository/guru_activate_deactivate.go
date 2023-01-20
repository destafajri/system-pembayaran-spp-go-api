package guru_repository

import (
	"errors"
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (guru *guruImplementation) ActivateGuru(user_id, guru_id string, timestamp time.Time) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query1 := `
		UPDATE users SET
			is_active = COALESCE($1, is_active),
			updated_at = $2
		WHERE 
			id = $3
		`
	values := []interface{}{
		true,
		timestamp,
		user_id,
	}

	_, err := guru.db.Exec(query1, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error activate user")
	}

	query2 := `
		UPDATE guru SET
			is_active = COALESCE($1, is_active)
		WHERE 
			id = $2
		`
	value := []interface{}{
		true,
		guru_id,
	}

	_, err = guru.db.Exec(query2, value...)
	if err != nil {
		log.Println(err)
		return errors.New("error activate guru")
	}

	return nil
}

func (guru *guruImplementation) DeactivateGuru(user_id, guru_id string, timestamp time.Time) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query1 := `
		UPDATE users SET
			is_active = COALESCE($1, is_active),
			updated_at = $2
		WHERE 
			id = $3
		`
	values := []interface{}{
		false,
		timestamp,
		user_id,
	}

	_, err := guru.db.Exec(query1, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error deactivate user")
	}

	query2 := `
		UPDATE guru SET
			is_active = COALESCE($1, is_active)
		WHERE 
			id = $2
		`
	value := []interface{}{
		false,
		guru_id,
	}

	_, err = guru.db.Exec(query2, value...)
	if err != nil {
		log.Println(err)
		return errors.New("error deactivate guru")
	}

	return nil
}
