package siswa_repository

import (
	"encoding/json"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/nullism/bqb"
	"github.com/pkg/errors"
)

func (siswa *siswaImplementation) GetDetailSiswaForAdmin(id string) (*model.GetDetailSiswaResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var data model.GetDetailSiswaResponse

	statement, params, err := siswa.getDetailAdminQuery(id)
	if err != nil {
		return nil, errors.Wrap(err, "build statement query to get siswa detail from database")
	}

	rows, err := siswa.db.Query(statement, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			return nil, errors.Wrap(err, "scanning siswa from database")
		}
		if err := json.Unmarshal(bson, &data); err != nil {
			return nil, errors.Wrap(err, "unmarshalling siswa bson")
		}
	}

	return &data, nil
}

func (repo *siswaImplementation) getDetailAdminQuery(id string) (string, []interface{}, error) {
	build := bqb.New(`
			SELECT 
			json_build_object(
						'id', siswa.id,
						'user_id', users.id,
						'email', users.email,
						'username', users.username,
						'role', users.role,
						'nis', siswa.nis,
						'nama', siswa.nama,
						'kelas', kelas.kelas,
						'angkatan', siswa.angkatan,
						'is_active', siswa.is_active,
						'created_at', users.created_at::timestamptz,
						'updated_at', users.updated_at::timestamptz,
						'deleted_at', users.deleted_at::timestamptz
					)
			FROM
				siswa
			JOIN
				users
			ON
				siswa.user_id = users.id
			JOIN
				kelas
			ON
				siswa.kelas_id = kelas.id
			WHERE 
				siswa.id = ?
			`, id)

	// build.Print()

	return build.ToPgsql()
}
