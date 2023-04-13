package repository

import (
	"fmt"
	"testing"
	"reflect"

	"github.com/jinzhu/gorm"
)

func TestWithOrderBy(t *testing.T) {
	tests := []struct {
		ob map[string]string
		want []func(db *gorm.DB) *gorm.DB
	}{
		{
			ob: map[string]string{
				"name": "",
				"age": "desc",
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
			ob: map[string]string{},
			want: []func(db *gorm.DB) *gorm.DB{},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("WithOrderBy(%v)",tt.ob), func(t *testing.T) {
			got := WithOrderBy(tt.ob)
			if len(got) != len(tt.want) {
				t.Errorf("WithOrderBy() = %v, want %v", got, tt.want)
			}
			db1, _ := gorm.Open("mysql","")
			db2, _ := gorm.Open("mysql","")
			for k,v := range got {
				db1 = v(db1)
				db2 = tt.want[k](db2)
			}
			if !reflect.DeepEqual(db1, db2) {
				t.Errorf("WithOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
