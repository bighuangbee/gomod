package storage

import (
	"github.com/jinzhu/gorm"
)

type Page struct {
	Page  	uint `form:"page"`
	Length 	uint `form:"length"`
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

	if page.Page <= 0 {
		page.Page = 1
	}
	if page.Length <= 0 {
		page.Length = 10
	}
	if page.Length > 200 {
		page.Length = 200
	}

	return DB.Model(pagination).Offset(int((page.Page-1)*page.Length)).Limit(page.Length)
}