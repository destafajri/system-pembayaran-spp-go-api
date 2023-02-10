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

func (guru *userImplementation) ActivateGuru(user_id, guru_id string, timestamp time.Time) error {
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

func (guru *userImplementation) DeactivateGuru(user_id, guru_id string, timestamp time.Time) error {
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

func (siswa *userImplementation) ActivateSiswa(user_id, siswa_id string, timestamp time.Time) error {
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

	_, err := siswa.db.Exec(query1, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error activate user")
	}

	query2 := `
		UPDATE siswa SET
			is_active = COALESCE($1, is_active)
		WHERE 
			id = $2
		`
	value := []interface{}{
		true,
		siswa_id,
	}

	_, err = siswa.db.Exec(query2, value...)
	if err != nil {
		log.Println(err)
		return errors.New("error activate siswa")
	}

	return nil
}

func (siswa *userImplementation) DeactivateSiswa(user_id, siswa_id string, timestamp time.Time) error {
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

	_, err := siswa.db.Exec(query1, values...)
	if err != nil {
		log.Println(err)
		return errors.New("error deactivate user")
	}

	query2 := `
		UPDATE siswa SET
			is_active = COALESCE($1, is_active)
		WHERE 
			id = $2
		`
	value := []interface{}{
		false,
		siswa_id,
	}

	_, err = siswa.db.Exec(query2, value...)
	if err != nil {
		log.Println(err)
		return errors.New("error deactivate siswa")
	}

	return nil
}
