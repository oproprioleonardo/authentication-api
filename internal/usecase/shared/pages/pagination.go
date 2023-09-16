package pages

type Pagination[T interface{}] struct {
	CurrentPage int64 `json:"currentPage,omitempty"`
	PerPage     int64 `json:"perPage,omitempty"`
	Total       int64 `json:"total,omitempty"`
	Items       []T   `json:"items,omitempty"`
}

func Map[T, K interface{}](p *Pagination[*T], f func(*T) K) Pagination[K] {
	result := Pagination[K]{
		CurrentPage: p.CurrentPage,
		PerPage:     p.PerPage,
		Total:       p.Total,
		Items:       make([]K, len(p.Items)),
	}

	for i, item := range p.Items {
		result.Items[i] = f(item)
	}

	return result
}
