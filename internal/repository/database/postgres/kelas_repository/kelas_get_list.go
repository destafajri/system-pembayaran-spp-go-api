package kelas_repository

import (
	"encoding/json"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta/param"
	"github.com/nullism/bqb"
	"github.com/pkg/errors"
)

func (kelas *kelasImplementation) GetListKelas(meta *meta.Metadata) ([]model.GetListKelasResponse, int, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	q, err := param.FromMetadata(meta, kelas)
	if err != nil {
		return nil, 0, errors.Wrap(err, "parsing metadata into query")
	}

	var (
		total    int
		count    = true
		notCount = !count
		data     []model.GetListKelasResponse
	)

	statement, params, err := kelas.getlistQuery(notCount, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.Wrap(err, "build statement query to get kelas from database")
	}

	rows, err := kelas.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var row model.GetListKelasResponse
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, 0, errors.New("scanning kelas from database")
		}

		if err := json.Unmarshal(bson, &row); err != nil {
			log.Println(err)
			return nil, 0, errors.Wrap(err, "unmarshalling kelas bson")
		}

		data = append(data, row)
	}

	// count total data
	statement, params, err = kelas.getlistQuery(count, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.Wrap(err, "build statement query to get kelas from database")
	}

	row := kelas.db.QueryRow(statement, params...)
	if err := row.Scan(&total); err != nil {
		log.Println(err)
		return nil, 0, errors.Wrap(err, "getting count kelas")
	}

	return data, total, nil
}

func (kelas *kelasImplementation) getlistQuery(is_count bool, q *param.Query) (string, []interface{}, error) {
	var selectx *bqb.Query

	if is_count {
		selectx = bqb.New(`
			SELECT 
				COUNT(kelas.id)
		`)
	} else {
		selectx = bqb.New(`
		SELECT 
			json_build_object(
						'id', kelas.id,
						'guru_id', kelas.guru_id,
						'walikelas', guru.nama,
						'kelas', kelas.kelas
			)
		`)
	}

	from := bqb.New(`
		FROM
			kelas
		JOIN
			guru
		ON
			kelas.guru_id = guru.id
	`)

	if is_count {
		return bqb.New("? ?", selectx, from).ToPgsql()
	}

	order := bqb.Optional("ORDER BY")
	if q.OrderBy != "" && q.OrderDirection != "" {
		order.Space("created_at").Space(q.OrderDirection)
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
