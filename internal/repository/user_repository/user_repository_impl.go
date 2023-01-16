package user_repository

import (
	"database/sql"
	"errors"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
)

type userImplementation struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository{
	return &userImplementation{
		db: db,
	}
}

func (user *userImplementation) Register(users *entity.UserEntity) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `INSERT INTO users(
			id,
			name, 
			phone,
			role,
			password
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`
	err := user.db.QueryRow(query, users.ID, users.Name, users.Phone, users.Role, users.Password).Scan(&users.ID)
	if err != nil {
		exception.PanicIfNeeded(err)
		return err
	}

	return nil
}

func (user userImplementation) GetData(phone string) (*entity.UserEntity, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var users entity.UserEntity
	query := `SELECT id, name, phone, role, password FROM users where phone=$1`

	err := user.db.QueryRow(query, phone).Scan(
		&users.ID,
		&users.Name,
		&users.Phone,
		&users.Role,
		&users.Password,
	)
	if err != nil {
		return nil, errors.New("User Not Found")
	}

	return &users, nil
}

func (user userImplementation) Login(phone string) (*entity.UserEntity, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var users entity.UserEntity

	query := `SELECT
				id,
				name,
				phone,
				role,
				password
			FROM users
			WHERE
				phone=$1
			`
	err := user.db.QueryRow(query, phone).Scan(
		&users.ID,
		&users.Name,
		&users.Phone,
		&users.Role,
		&users.Password,
	)

	if err != nil {
		exception.PanicIfNeeded(err)
		return nil, err
	}

	return &users, nil
}
