package user_repository

import (
	"encoding/json"
	"log"

	"github.com/nullism/bqb"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta/param"
	"github.com/pkg/errors"
)

func (user *userImplementation) GetListUser(meta *meta.Metadata) ([]model.GetListUserResponse, int, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	q, err := param.FromMetadata(meta, user)
	if err != nil {
		return nil, 0, errors.Wrap(err, "parsing metadata into query")
	}

	var (
		total    int
		count    = true
		notCount = !count
		data     []model.GetListUserResponse
	)

	statement, params, err := user.getlistQuery(notCount, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.Wrap(err, "build statement query to get user from database")
	}

	rows, err := user.db.Query(statement, params...)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var row model.GetListUserResponse
		var bson []byte

		if err := rows.Scan(&bson); err != nil {
			log.Println(err)
			return nil, 0, errors.New("scanning user from database")
		}

		if err := json.Unmarshal(bson, &row); err != nil {
			log.Println(err)
			return nil, 0, errors.Wrap(err, "unmarshalling user bson")
		}

		data = append(data, row)
	}

	// count total data
	statement, params, err = user.getlistQuery(count, q)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.Wrap(err, "build statement query to get user from database")
	}

	row := user.db.QueryRow(statement, params...)
	if err := row.Scan(&total); err != nil {
		log.Println(err)
		return nil, 0, errors.Wrap(err, "getting count user")
	}

	return data, total, nil
}

func (user *userImplementation) getlistQuery(is_count bool, q *param.Query) (string, []interface{}, error) {
	var selectx *bqb.Query

	if is_count {
		selectx = bqb.New(`
			SELECT 
				COUNT(id)
		`)
	} else {
		selectx = bqb.New(`
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
		`)
	}

	from := bqb.New(`
		FROM
			users
	`)

	if is_count {
		return bqb.New("? ? ", selectx, from).ToPgsql()
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
