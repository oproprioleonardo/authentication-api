package query

import "github.com/skyepic/privateapi/internal/usecase/shared/pages"

type Params struct {
	Page    int64 `query:"page,omitempty"`
	PerPage int64 `query:"perPage,omitempty"`
}

func (q Params) ToDomain() pages.Query {
	return pages.Query{
		Page:    q.Page,
		PerPage: q.PerPage,
	}
}
