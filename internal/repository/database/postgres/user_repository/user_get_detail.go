package user_repository

import (
	"encoding/json"

	"github.com/nullism/bqb"
	"github.com/pkg/errors"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

func (user *userImplementation) GetDetailUser(id string) (*model.GetDetailUser, error) {
	var data model.GetDetailUser

	statement, params, err := user.getDetailQuery(id)
	if err != nil {
		return nil, errors.Wrap(err, "build statement query to get user detail from database")
	}

	rows, err := user.db.Query(statement, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			return nil, errors.Wrap(err, "scanning user from database")
		}
		if err := json.Unmarshal(bson, &data); err != nil {
			return nil, errors.Wrap(err, "unmarshalling user bson")
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
					AND
				is_active is true
		`, id)

	// build.Print()

	return build.ToPgsql()
}
