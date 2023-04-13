package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

/**
	Create a order by scope 
**/
func WithOrderBy(ob map[string]string) []func(db *gorm.DB) *gorm.DB {
	var clauses []func(*gorm.DB) *gorm.DB
	for k, v := range ob {
		order := fmt.Sprintf("%s %s", k, v)
		clauses = append(clauses, func(db *gorm.DB) *gorm.DB { return db.Order(order) })
	}
	return clauses
}

/**
	Create a pagination scope 
**/
func WithPagination(pageSize int, offSet int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offSet).Limit(pageSize)
	}
}
