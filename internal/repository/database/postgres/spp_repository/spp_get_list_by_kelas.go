package spp_repository

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta/param"
	"github.com/nullism/bqb"
)

func (spp *sppImplementation) GetListSppByKelas(kelas_id string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	q, err := param.FromMetadata(meta, spp)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("parsing metadata into query")
	}

	var (
		total    int
		count    = true
		notCount = !count
		data     []model.GetListSppResponse
	)

	statement, params, err := spp.getlistSppBykelasQuery(kelas_id, notCount, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("build statement query to get spp by kelas from database")
	}

	rows, err := spp.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, 0, exception.ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		var row model.GetListSppResponse
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, 0, errors.New("scanning spp by kelas from database")
		}

		if err := json.Unmarshal(bson, &row); err != nil {
			log.Println(err)
			return nil, 0, errors.New("unmarshalling spp by kelas bson")
		}

		data = append(data, row)
	}

	// count total data
	statement, params, err = spp.getlistSppBykelasQuery(kelas_id, count, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("build statement query to get spp by kelas from database")
	}

	row := spp.db.QueryRow(statement, params...)
	if err := row.Scan(&total); err != nil {
		log.Println(err)
		return nil, 0, errors.New("getting count spp")
	}

	return data, total, nil
}

func (spp *sppImplementation) getlistSppBykelasQuery(kelas_id string, is_count bool, q *param.Query) (string, []interface{}, error) {
	var selectx *bqb.Query

	if is_count {
		selectx = bqb.New(`
			SELECT 
				COUNT(spp.id)
		`)
	} else {
		selectx = bqb.New(`
		SELECT 
			json_build_object(
						'id', spp.id,
						'siswa_id', siswa.id,
						'nis', siswa.nis,
						'no_spp', spp.no_spp,
						'jatuh_tempo', spp.jatuh_tempo,
						'jumlah', spp.jumlah,
						'tanggal_bayar', ((CASE
											WHEN ((
												SELECT
													bayar.spp_id
												FROM
													bayar
												WHERE spp_id = spp.id
													) IS NOT NULL )
														 THEN (
															SELECT
																tanggal_bayar
															FROM
																bayar
															JOIN spp ON bayar.spp_id = spp.id)
											ELSE null
											END)),
						  'status', ((CASE
										WHEN ((
											SELECT
												bayar.spp_id
											FROM
												bayar
											WHERE spp_id = spp.id
											) IS NOT NULL )
											THEN 'paid'
										ELSE 'unpaid'
									END)),
						'created_at', created_at::timestamptz,
						'updated_at', updated_at::timestamptz,
						'deleted_at', deleted_at::timestamptz
			)
		`)
	}

	from := bqb.New(`
		FROM
			spp
		JOIN
			siswa
		ON
			spp.siswa_id = siswa.id
		WHERE siswa.kelas_id = ?
	`, kelas_id)


	and := bqb.Optional("AND")

	if q.Status != "%%" {
		and.Space(`(CASE
			WHEN ((
				SELECT
					bayar.spp_id
				FROM
					bayar
				WHERE spp_id = spp.id
				) IS NOT NULL )
				THEN 'paid'
			ELSE 'unpaid'
		END) = ?`, strings.ReplaceAll(q.Status, "%", ""))
	}

	if is_count {
		return bqb.New("? ? ?", selectx, from, and).ToPgsql()
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

	buildx := bqb.New("? ? ? ? ? ?", selectx, from, and, order, limit, offset)
	// buildx.Print()

	return buildx.ToPgsql()
}
