package guru_repository

import (
	"encoding/json"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/nullism/bqb"
	"github.com/pkg/errors"
)

func (guru *guruImplementation) GetDetailGuruAdmin(guru_id string) (*model.GetDetailGuruResponse, error){
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var data model.GetDetailGuruResponse

	statement, params, err := guru.getDetailForAdminQuery(guru_id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("build statement query to get guru detail from database")
	}

	rows, err := guru.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, exception.ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, errors.New("scanning guru from database")
		}
		if err := json.Unmarshal(bson, &data); err != nil {
			log.Println(err)
			return nil, errors.New("unmarshalling guru bson")
		}
	}

	return &data, nil
}

func (repo *guruImplementation) getDetailForAdminQuery(id string) (string, []interface{}, error) {
	build := bqb.New(`
			SELECT 
			json_build_object(
						'id', guru.id,
						'user_id', users.id,
						'email', users.email,
						'username', users.username,
						'nama', guru.nama,
						'role', users.role,
						'is_active', guru.is_active,
						'created_at', users.created_at::timestamptz,
						'updated_at', users.updated_at::timestamptz,
						'deleted_at', users.deleted_at::timestamptz
			)
			FROM
				guru
			JOIN
				users
			ON
				guru.user_id = users.id
			WHERE 
				guru.id = ?
		`, id)

	// build.Print()

	return build.ToPgsql()
}
