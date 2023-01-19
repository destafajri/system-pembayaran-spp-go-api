package param

import (
	"database/sql"
	"strings"

	"github.com/destafajri/system-pembayaran-spp-go-api/helper/timeutil"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func Filter(metadata *meta.Metadata) *Query {
	var form, end sql.NullTime
	if metadata.DateRange != nil {

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

	q := Query{
		OrderBy:        metadata.OrderBy,
		OrderDirection: metadata.OrderType,
		Search:         search,
		Limit:          limit,
		Offset:         offset,
		DateFrom:       form,
		DateEnd:        end,
	}

	return &q
}
