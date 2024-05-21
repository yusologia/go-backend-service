package Repository

import (
	logiares "github.com/yusologia/go-response"
	"gorm.io/gorm"
	"net/url"
	"strconv"
)

func Paginate(parameters url.Values, query *gorm.DB, model interface{}) (*gorm.DB, logiares.Pagination) {
	var count int64
	query.Model(&model).Count(&count)

	page, _ := strconv.Atoi(parameters.Get("page"))
	limit, _ := strconv.Atoi(parameters.Get("limit"))
	if limit == 0 {
		limit = 50
	}

	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	dataPagination := logiares.Pagination{
		Count:       int(count),
		PerPage:     limit,
		CurrentPage: page,
		TotalPage:   (int(count) / limit) + 1,
	}

	if page < dataPagination.TotalPage {
		dataPagination.NextPage = page + 1
	}

	if page > 1 {
		dataPagination.PrevPage = page - 1
	}

	return query, dataPagination
}
