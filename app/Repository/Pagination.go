package Repository

import (
	"gorm.io/gorm"
	"net/url"
	"strconv"
)

type Pagination struct {
	Count       int
	CurrentPage any
	PerPage     int
	TotalPage   int
}

func Paginate(parameters url.Values, query *gorm.DB, model interface{}) (*gorm.DB, Pagination) {
	var count int64
	query.Model(&model).Count(&count)

	page, _ := strconv.Atoi(parameters.Get("page"))
	limit, _ := strconv.Atoi(parameters.Get("limit"))
	if limit == 0 {
		limit = 50
	}

	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	dataPagination := Pagination{
		Count:       int(count),
		PerPage:     limit,
		CurrentPage: page,
		TotalPage:   (int(count) / limit) + 1,
	}

	return query, dataPagination
}

func (p Pagination) ParsePagination() map[string]interface{} {
	return map[string]interface{}{
		"count":       p.Count,
		"currentPage": p.CurrentPage,
		"perPage":     p.PerPage,
		"totalPage":   p.TotalPage,
	}
}
