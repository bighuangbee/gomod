package storage

import (
	"github.com/jinzhu/gorm"
)

type Page struct {
	Start  	uint `form:"start"`
	End 	uint `form:"end"`
}

type PageResult struct {
	List 	interface{} `json:"list"`
	Total 	uint	`json:"total"`
}

/**
	分页接口
 */
type Pagination interface {
	Pagination(Page)
}

/*
	根据分页接口的实现对象，动态创建DB分页模型
	@param pagination 接收实现了分页接口的模型（对象）作为参数
*/
func PageDB(pagination Pagination, page Page) (*gorm.DB){

	if page.Start <= 0 {
		page.Start = 1
	}
	if page.End <= 0 {
		page.End = 10
	}
	if page.End < page.Start {
		page.End = page.Start
	}

	limit := (page.End - page.Start) + 1

	return DB.Model(pagination).Offset(page.Start).Limit(limit)
}