package logQuery

import (
	"blog_server/global"

	"gorm.io/gorm"
)

type PageInfo struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
	Order string `form:"order"`
}

func (p PageInfo) GetLimit() int {
	if p.Limit <= 0 || p.Limit > 20 {
		return 10
	}
	return p.Limit
}

func (p PageInfo) GetOffset() int {
	return p.GetLimit() * (p.GetPage() - 1)
}

func (p PageInfo) GetPage() int {
	if p.Page > 20 || p.Page <= 0 {
		return 1
	}
	return p.Page
}

type Options struct {
	PageInfo
	Likes        []string
	Preloads     []string
	Where        *gorm.DB
	Debug        bool
	DefaultOrder string
}

func ListQuery[T any](model any, options Options) (list []T, total int, err error) {
	//正常查询
	query := global.DB.Where(model)

	//调试模式
	if options.Debug == true {
		query = query.Debug()
	}

	//排序
	if options.Order != "" {
		query = query.Order(options.Order)
	} else if options.DefaultOrder != "" {
		query = query.Order(options.DefaultOrder) 
	}

	// 模糊查询
	if len(options.Likes) > 0 && options.Key != ""{
		var likes *gorm.DB
		for _, v := range options.Likes {
			likes = global.DB.Or("? like ?", gorm.Expr(v), "%"+options.Key+"%")
		}
		query = query.Debug().Where(likes)
	}

	//预加载
	if len(options.Preloads) > 0 {
		for _, v := range options.Preloads {
			query = query.Preload(v)
		}
	}

	//定制化查询
	if options.Where != nil {
		query = query.Where(options.Where)
	}

	// 获取日志列表
	limit := options.GetLimit()
	offst := options.GetOffset()
	err = query.Offset(offst).Limit(limit).Find(&list).Error
	return list, len(list), err
}
