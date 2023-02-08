package user_repository

import (
	"encoding/json"
	"log"

	"github.com/nullism/bqb"
	"github.com/pkg/errors"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
)

func (user *userImplementation) GetDetailUser(id string) (*model.GetDetailUser, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var data model.GetDetailUser

	statement, params, err := user.getDetailQuery(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("build statement query to get user detail from database")
	}

	rows, err := user.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, exception.ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, errors.New("scanning user from database")
		}
		if err := json.Unmarshal(bson, &data); err != nil {
			log.Println(err)
			return nil, errors.New("unmarshalling user bson")
		}
	}

	return &data, nil
}

func (repo *userImplementation) getDetailQuery(id string) (string, []interface{}, error) {
	build := bqb.New(`
			SELECT 
			json_build_object(
						'id', id,
						'email', email,
						'username', username,
						'role', role,
						'is_active', is_active,
						'created_at', created_at::timestamptz,
						'updated_at', updated_at::timestamptz,
						'deleted_at', deleted_at::timestamptz
					)
			FROM
				users
			WHERE id = ?
		`, id)

	// build.Print()

	return build.ToPgsql()
}
