package model

import (
	"github.com/jinzhu/gorm"
	"go-index/packages/gin_series/12_jwt/initDB"
)

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	if !initDB.Db.HasTable(Role{}) {
		initDB.Db.CreateTable(Role{})
	}
}
