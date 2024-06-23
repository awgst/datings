// Package pagination provides pagination functions
package pagination

import (
	"fmt"
	"math"
)

type Paginator struct {
	Limit       int
	Offset      int
	Total       int
	CurrentPage int
	PathUrl     string
	Count       int64
}

func (p *Paginator) GetLinks() map[string]any {
	if p.CurrentPage > p.Total {
		p.CurrentPage = p.Total
	}

	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}

	nextUrl := fmt.Sprintf("%s?page=%d", p.PathUrl, p.CurrentPage+1)
	if p.CurrentPage >= p.Total {
		nextUrl = ""
	}

	prevUrl := fmt.Sprintf("%s?page=%d", p.PathUrl, p.CurrentPage-1)
	if p.CurrentPage <= 1 {
		prevUrl = ""
	}

	return map[string]any{
		"current_page": p.CurrentPage,
		"total_page":   p.Total,
		"limit":        p.Limit,
		"next_url":     nextUrl,
		"previous_url": prevUrl,
		"total_data":   p.Count,
	}
}

func (p *Paginator) New(limit, page int, path string) {
	p.Limit = 10
	if limit != 0 {
		p.Limit = limit
	}

	p.CurrentPage = 1
	if page != 0 {
		p.CurrentPage = page
	}

	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}

	p.Offset = (p.CurrentPage - 1) * p.Limit

	p.PathUrl = path
}

func (p *Paginator) SetTotal(total int64) {
	if total == 0 {
		return
	}
	p.Count = total
	p.Total = int(math.Ceil(float64(total) / float64(p.Limit)))
}
