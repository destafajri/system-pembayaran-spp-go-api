package meta

import (
	"errors"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ErrInvalidMetadata is an error when metadata is invalid.
// This error usually returned by the implementation of Filter interface.
var ErrInvalidMetadata = errors.New("invalid metadata")

// Metadata represents a metadata for HTTP API.
type Metadata struct {
	Pagination
	Filtering
	*DateRange `json:"date_range,omitempty"`
}

type MetadataPage struct {
	PaginationPage `json:"pagination"`
	Filtering
	*DateRange `json:"date_range,omitempty"`
}

// MetadataFromURL gets metadata from the given request url.
func MetadataFromURL(u url.Values) Metadata {
	return Metadata{
		Pagination: PaginationFromURL(u),
		Filtering:  FilterFromURL(u),
	}
}

// DefaultPerPage is a default value for per_page query params.
const DefaultPerPage = 10

// Pagination is a meta data for pagination.
type Pagination struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
	Total   int `json:"total"`
}

// PaginationPage is a meta data for pagination with total page.
type PaginationPage struct {
	PerPage   int `json:"per_page"`
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
}

// PaginationFromURL gets pagination meta from request URL.
func PaginationFromURL(u url.Values) Pagination {
	p := Pagination{
		PerPage: DefaultPerPage,
		Page:    1,
	}

	pps := u.Get("per_page")
	if v, err := strconv.Atoi(pps); err == nil {
		if v <= 0 {
			v = DefaultPerPage
		}

		p.PerPage = v
	}

	ps := u.Get("page")
	if v, err := strconv.Atoi(ps); err == nil {
		if v < 1 {
			v = 1
		}

		p.Page = v
	}

	return p
}

// SortXXX are default values for order_type query params.
const (
	SortAscending  = "asc"
	SortDescending = "desc"
)

// Filtering represents a filterable fields.
type Filtering struct {
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`
	Search    string `json:"search,omitempty"`
	SearchBy  string `json:"search_by,omitempty"`
	Status    string `json:"status,omitempty"`
}

// FilterFromURL gets filter values from query params.
func FilterFromURL(u url.Values) Filtering {
	f := Filtering{
		OrderBy:   "created_at",
		OrderType: SortAscending,
	}

	ob := u.Get("order_by")
	if len(ob) != 0 {
		f.OrderBy = strings.ToLower(strings.ToLower(ob))
	}

	ot := u.Get("order_type")
	if len(ot) != 0 {
		ot = strings.TrimSpace(strings.ToLower(ot))
		if ot == SortDescending {
			f.OrderType = SortDescending
		}
	}

	search := strings.TrimSpace(u.Get("search"))
	if len(search) == 0 {
		search = strings.TrimSpace(u.Get("keyword"))
	}

	if len(search) != 0 {
		f.Search = search
	}

	searchBy := strings.TrimSpace(u.Get("search_by"))
	if len(searchBy) != 0 {
		f.SearchBy = searchBy
	}

	status := strings.TrimSpace(u.Get("status"))
	if len(status) == 0 {
		status = strings.TrimSpace(u.Get("keyword"))
	}

	if len(status) != 0 {
		f.Status = status
	}

	return f
}

type DateRange struct {
	Field string    `json:"field"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func DateRangeFromURL(u url.Values, field string, startQuery, endQuery string) (*DateRange, error) {
	ts := u.Get(startQuery)
	te := u.Get(endQuery)
	if len(ts) == 0 || len(te) == 0 {
		return nil, nil
	}

	dr := DateRange{
		Field: "created_at",
		Start: time.Time{},
		End:   time.Time{},
	}

	if v := u.Get(field); len(v) != 0 {
		dr.Field = strings.TrimSpace(strings.ToLower(v))
	}

	t, err := time.Parse("2006-01-02", ts)
	if err != nil {
		return nil, ErrInvalidMetadata
	}

	dr.Start = t

	t, err = time.Parse("2006-01-02", te)
	if err != nil {
		return nil, ErrInvalidMetadata
	}

	dr.End = t

	return &dr, nil
}

// Filter knows how to validate filterable fields.
// This Filter usually implemented by Repository.
type Filter interface {
	// Sortable returns true if a given field is allowed for sorting.
	Sortable(field string) bool
}

// PageCalculate calculate total page from count
func PageCalculate(count int64, limit int64) int64 {
	if count <= limit {
		return 1
	}

	return int64(math.Ceil(float64(count) / float64((limit))))
}

// Convert Meta to MetaPage
func ConvertMetaPage(m Metadata) MetadataPage {
	meta := MetadataPage{
		PaginationPage: PaginationPage{
			PerPage:   m.PerPage,
			Page:      m.Page,
			TotalPage: int(PageCalculate(int64(m.Total), int64(m.PerPage))),
		},
		Filtering: Filtering{
			OrderBy:   m.OrderBy,
			OrderType: m.OrderType,
			Search:    m.Search,
			SearchBy:  m.SearchBy,
		},
	}
	meta.DateRange = m.DateRange
	return meta
}
