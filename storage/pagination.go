package storage

import (
	"github.com/jinzhu/gorm"
)

type PageQuery struct {
	Page  	uint `form:"page"`
	Length 	uint `form:"length"`
}

func (page *PageQuery)Offset()uint{
	if page.Page <= 0 {
		page.Page = 1
	}

	return (page.Page-1)*page.Limit()
}

func (page *PageQuery)Limit()uint{
	if page.Length <= 0 {
		return 10
	}
	if page.Length > 500 {
		return 500
	}
	return page.Length
}


type PageResult map[string]interface{}

/**
	分页接口
 */
type Pagination interface {
	Pagination(PageQuery)
}

/*
	根据分页接口的实现对象，动态创建DB分页模型
	@param pagination 接收实现了分页接口的模型（对象）作为参数
*/
func PageDB(pagination Pagination, page PageQuery) (*gorm.DB){

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
