package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 获取标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// 获取标签总数
func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// 添加标签
func AddTag(name string, state int, createdBy string) bool {
	dbResult := db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	fmt.Println("error ", dbResult.Error)

	if dbResult.Error != nil {
		log.Fatalf("has error: %v", dbResult.Error)
		return false
	}
	return true
}

func (tag *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	tag.CreatedOn = time.Now().Unix()
	return
}

func (tag *Tag) BeforeUpdate(tx *gorm.DB) (err error) {
	tag.ModifiedOn = time.Now().Unix()
	return nil
}
