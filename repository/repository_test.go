package repository

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestWithOrderBy(t *testing.T) {
	tests := []struct {
		ob   map[string]string
		want []func(db *gorm.DB) *gorm.DB
	}{
		{
			ob: map[string]string{
				"name": "",
				"age":  "desc",
			},
			want: []func(db *gorm.DB) *gorm.DB{
				func(db *gorm.DB) *gorm.DB {
					return db.Order("name ")
				},
				func(db *gorm.DB) *gorm.DB {
					return db.Order("age desc")
				},
			},
		},
		{
			ob: map[string]string{
				"name": "",
			},
			want: []func(db *gorm.DB) *gorm.DB{
				func(db *gorm.DB) *gorm.DB {
					return db.Order("name ")
				},
			},
		},
		{
			ob:   map[string]string{},
			want: []func(db *gorm.DB) *gorm.DB{},
		},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestWithOrderBy(%v)", input.ob), func(t *testing.T) {
			got := WithOrderBy(input.ob)
			if len(got) != len(input.want) {
				t.Errorf("WithOrderBy(%v) output array length = %#v, want length %#v", input.ob, got, input.want)
			}
		})
	}
}
