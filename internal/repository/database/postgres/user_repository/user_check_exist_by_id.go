package user_repository

import (
	"errors"
	"log"
)

func (user *userImplementation) CekUserExistByID(id string) (bool, error){
	query := `SELECT id FROM users WHERE id = $1 AND deleted_at IS NULL AND is_active is true LIMIT 1;`

	rows, err := user.db.Query(query, id)
	if err != nil {
		log.Println(err)
		return false, errors.New("user not exist")
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return rows.Next(), errors.New("user not exist")
}