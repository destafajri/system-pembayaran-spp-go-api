package param

import (
	"database/sql"
	"strings"

	"github.com/destafajri/system-pembayaran-spp-go-api/helper/timeutil"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type Query struct {
	OrderBy        string
	OrderDirection string
	Search         string
	SearchBy       string
	Status         string
	Limit          int
	Offset         int
	DateFrom       sql.NullTime
	DateEnd        sql.NullTime
}

func FromMetadata(metadata *meta.Metadata, filter meta.Filter) (*Query, error) {
	if !filter.Sortable(metadata.OrderBy) {
		return nil, meta.ErrInvalidMetadata
	}

	var form, end sql.NullTime
	if metadata.DateRange != nil {
		if !filter.Sortable(metadata.DateRange.Field) {
			return nil, meta.ErrInvalidMetadata
		}

		form = sql.NullTime{
			Time:  timeutil.BeginOfDay(metadata.DateRange.Start),
			Valid: !metadata.DateRange.Start.IsZero(),
		}

		end = sql.NullTime{
			Time:  timeutil.BeginOfNextDay(metadata.DateRange.End),
			Valid: !metadata.DateRange.End.IsZero(),
		}
	}

	limit := metadata.PerPage
	offset := (metadata.Page - 1) * limit
	search := "%" + strings.ToLower(metadata.Search) + "%"
	status := "%" + strings.ToLower(metadata.Status) + "%"

	q := Query{
		OrderBy:        metadata.OrderBy,
		OrderDirection: metadata.OrderType,
		Search:         search,
		Limit:          limit,
		Offset:         offset,
		DateFrom:       form,
		DateEnd:        end,
		SearchBy:       strings.ToLower(metadata.SearchBy),
		Status:         status,
	}

	return &q, nil
}
