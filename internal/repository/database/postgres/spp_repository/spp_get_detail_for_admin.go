package spp_repository

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/nullism/bqb"
)

func (spp *sppImplementation) GetDetailSpp(spp_id string) (*model.GetDetailSppResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var data model.GetDetailSppResponse

	statement, params, err := spp.getDetailSppQuery(spp_id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("build statement query to get spp detail from database")
	}

	rows, err := spp.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, exception.ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, errors.New("scanning spp from database")
		}
		if err := json.Unmarshal(bson, &data); err != nil {
			log.Println(err)
			return nil, errors.New("unmarshalling spp bson")
		}
	}

	return &data, nil
}

func (repo *sppImplementation) getDetailSppQuery(spp_id string) (string, []interface{}, error) {
	build := bqb.New(`
			SELECT 
			json_build_object(
						'id', spp.id,
						'siswa_id', siswa.id,
						'nama', siswa.nama,
						'kelas', kelas.kelas,
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
															JOIN spp ON bayar.spp_id = spp.id
															LIMIT 1
														)
											ELSE null
											END)),
						'status', ((CASE
										WHEN ((
											SELECT
												bayar.spp_id
											FROM
												bayar
											WHERE spp_id = spp.id
											LIMIT 1
											) IS NOT NULL )
											THEN 'paid'
										ELSE 'unpaid'
									END)),
						'created_at', spp.created_at::timestamptz,
						'updated_at', spp.updated_at::timestamptz,
						'deleted_at', spp.deleted_at::timestamptz
					)
			FROM
				spp
			JOIN
				siswa
			ON
				spp.siswa_id = siswa.id
			JOIN
				kelas
			ON
				siswa.kelas_id = kelas.id
			WHERE 
				spp.id = ?
			LIMIT 1
			`, spp_id)

	// build.Print()

	return build.ToPgsql()
}
