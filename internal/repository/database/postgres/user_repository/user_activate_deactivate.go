package user_repository

import (
	"errors"
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (user *userImplementation) ActivateUser(id string, timestamp time.Time) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		UPDATE users SET
			is_active = COALESCE($1, is_active),
			updated_at = $2
		WHERE 
			id = $3
		`
	values := []interface{}{
		true,
		timestamp,
		id,
	}

	_, err := user.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error activate user")
	}

	return nil
}

func (user *userImplementation) DeactivateUser(id string, timestamp time.Time) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		UPDATE users SET
			is_active = COALESCE($1, is_active),
			updated_at = $2
		WHERE 
			id = $3
		`
	values := []interface{}{
		false,
		timestamp,
		id,
	}

	_, err := user.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error deactivate user")
	}

	return nil
}
