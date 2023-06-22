package controllers

import (
	"fmt"
	"math"
)

type PageLink struct {
	Page          int
	URL           string
	IsCurrentPage bool
}

type PaginationLink struct {
	CurrentPage string
	NextPage    string
	PrevPage    string
	TotalRows   int
	TotalPage   int
	Links       []PageLink
}

type PaginationParam struct {
	Path        string
	TotalRows   int
	PerPage     int
	CurrentPage int
}

func GetPagintionLink(param PaginationParam) (PaginationLink, error) {
	var Links []PageLink
	totalpage := int(math.Ceil(float64(param.TotalRows) / float64(param.PerPage)))
	if totalpage > 0 {
		for i := 1; i <= totalpage; i++ {
			Links = append(Links, PageLink{
				Page:          i,
				URL:           fmt.Sprintf("http://localhost:1323/products?page=%d", i),
				IsCurrentPage: i == param.CurrentPage,
			})
		}
	}
	prev := param.CurrentPage
	next := param.CurrentPage
	if totalpage > 1 {
		prev = param.CurrentPage - 1
	}
	if param.CurrentPage < totalpage {
		next = param.CurrentPage + 1
	}
	return PaginationLink{
		CurrentPage: fmt.Sprintf("http://localhost:1323/products?page=%d", param.CurrentPage),
		NextPage:    fmt.Sprintf("http://localhost:1323/products?page=%d", next),
		PrevPage:    fmt.Sprintf("http://localhost:1323/products?page=%d", prev),
		TotalRows:   param.TotalRows,
		TotalPage:   totalpage,
		Links:       Links,
	}, nil
}
