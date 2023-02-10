package guru_repository

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

func (guru *guruImplementation) GetListGuruAdmin(meta *meta.Metadata) ([]model.GetListGuruResponse, int, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	q, err := param.FromMetadata(meta, guru)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("parsing metadata into query")
	}

	var (
		total    int
		count    = true
		notCount = !count
		data     []model.GetListGuruResponse
	)

	statement, params, err := guru.getlistForAdminQuery(notCount, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("build statement query to get guru from database")
	}

	rows, err := guru.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, 0, exception.ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		var row model.GetListGuruResponse
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, 0, errors.New("scanning guru from database")
		}

		if err := json.Unmarshal(bson, &row); err != nil {
			log.Println(err)
			return nil, 0, errors.New("unmarshalling guru bson")
		}

		data = append(data, row)
	}

	// count total data
	statement, params, err = guru.getlistForAdminQuery(count, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.New("build statement query to get guru from database")
	}

	row := guru.db.QueryRow(statement, params...)
	if err := row.Scan(&total); err != nil {
		log.Println(err)
		return nil, 0, errors.New("getting count guru")
	}

	return data, total, nil
}

func (guru *guruImplementation) getlistForAdminQuery(is_count bool, q *param.Query) (string, []interface{}, error) {
	var selectx *bqb.Query

	if is_count {
		selectx = bqb.New(`
			SELECT 
				COUNT(guru.id)
		`)
	} else {
		selectx = bqb.New(`
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
		`)
	}

	from := bqb.New(`
		FROM
			guru
		JOIN
			users
		ON
			guru.user_id = users.id
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
