package siswa_repository

import (
	"encoding/json"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta/param"
	"github.com/nullism/bqb"
	"github.com/pkg/errors"
)

func (siswa *siswaImplementation) GetListSiswaByKelas(kelas_id string, meta *meta.Metadata) ([]model.GetListSiswaByKelasResponse, int, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	q, err := param.FromMetadata(meta, siswa)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("parsing metadata into query")
	}

	var (
		total    int
		count    = true
		notCount = !count
		data     []model.GetListSiswaByKelasResponse
	)

	statement, params, err := siswa.getlistByKelasQuery(kelas_id, notCount, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("build statement query to get siswa by kelas from database")
	}

	rows, err := siswa.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, 0, exception.ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		var row model.GetListSiswaByKelasResponse
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, 0, errors.New("scanning siswa from database")
		}

		if err := json.Unmarshal(bson, &row); err != nil {
			log.Println(err)
			return nil, 0, errors.New("unmarshalling siswa bson")
		}

		data = append(data, row)
	}

	// count total data
	statement, params, err = siswa.getlistByKelasQuery(kelas_id, count, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("build statement query to get siswa by kelas from database")
	}

	row := siswa.db.QueryRow(statement, params...)
	if err := row.Scan(&total); err != nil {
		log.Println(err)
		return nil, 0, errors.New("getting count siswa")
	}

	return data, total, nil
}
func (siswa *siswaImplementation) getlistByKelasQuery(kelas_id string, is_count bool, q *param.Query) (string, []interface{}, error) {
	var selectx *bqb.Query

	if is_count {
		selectx = bqb.New(`
			SELECT 
				COUNT(siswa.id)
		`)
	} else {
		selectx = bqb.New(`
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
		`)
	}

	from := bqb.New(`
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
			siswa.is_active is true
		AND
			siswa.kelas_id = ?
	`, kelas_id)

	if is_count {
		return bqb.New("? ?", selectx, from).ToPgsql()
	}

	order := bqb.Optional("ORDER BY")
	if q.OrderBy != "" && q.OrderDirection != "" {
		order.Space("users.created_at").Space(q.OrderDirection)
	}

	limit := bqb.Optional("LIMIT")
	if q.Limit > 0 {
		limit.Space("?", q.Limit)
	}

	offset := bqb.Optional("OFFSET")
	if q.Offset > 0 {
		offset.Space("?", q.Offset)
	}

	buildx := bqb.New("? ? ? ? ?", selectx, from, order, limit, offset)
	// buildx.Print()

	return buildx.ToPgsql()
}
