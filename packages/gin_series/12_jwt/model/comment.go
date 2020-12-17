package model

import (
	"github.com/jinzhu/gorm"
	"go-index/packages/gin_series/12_jwt/initDB"
)

type Comment struct {
	gorm.Model
	Content string
}

func init() {
	table := initDB.Db.HasTable(Comment{})
	if !table {
		initDB.Db.CreateTable(Comment{})
	}
}
