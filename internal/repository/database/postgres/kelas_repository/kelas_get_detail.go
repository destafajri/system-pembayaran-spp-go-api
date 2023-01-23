package kelas_repository

import (
	"encoding/json"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/nullism/bqb"
	"github.com/pkg/errors"
)

func (kelas *kelasImplementation) GetDetailKelas(kelas_id string) (*model.GetDetailKelasResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var data model.GetDetailKelasResponse

	statement, params, err := kelas.getDetailQuery(kelas_id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("build statement query to get kelas detail from database")
	}

	rows, err := kelas.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, exception.ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, errors.New("scanning kelas from database")
		}
		if err := json.Unmarshal(bson, &data); err != nil {
			log.Println(err)
			return nil, errors.New("unmarshalling kelas bson")
		}
	}

	return &data, nil
}

func (repo *kelasImplementation) getDetailQuery(kelas_id string) (string, []interface{}, error) {
	build := bqb.New(`
			SELECT 
			json_build_object(
						'id', kelas.id,
						'guru_id', guru.id,
						'walikelas', guru.nama,
						'kelas', kelas.kelas,
						'created_at', created_at::timestamptz,
						'updated_at', updated_at::timestamptz,
						'deleted_at', deleted_at::timestamptz
					)
			FROM
				kelas
			JOIN
				guru
			ON
				kelas.guru_id = guru.id
			WHERE kelas.id = ?
		`, kelas_id)

	// build.Print()

	return build.ToPgsql()
}

