package models

import (
	"errors"
	"fmt"
	"gin-example/pkg/logging"
	"gorm.io/gorm"
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
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
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
		logging.Fatal("has error: %v", dbResult.Error)
		return false
	}
	return true
}

// 通过 id 判定是否存在tag
func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// 通过 id 删除 tag
func DeleteTag(id int) bool {
	if db.Where("id = ?", id).Delete(&Tag{}).Error != nil {
		return true
	}
	return false
}

// 通过 id 编辑 tag
func EditTag(id int, data map[string]interface{}) bool {
	if db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error != nil {
		return true
	}
	return false
}

func FindOneTag(id int) (tag Tag, err error) {
	err = db.Where("id = ?", id).Find(&tag).Error
	if tag.ID == 0 {
		return tag, errors.New("tag is not exist")
	}
	if err != nil {
		return tag, err
	}
	return tag, nil
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{})
	return true
}

/*
gorm所支持的回调方法：
	创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
	更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
	删除：BeforeDelete、AfterDelete
	查询：AfterFind
*/
func (tag *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	tag.CreatedOn = time.Now().Unix()
	return
}

func (tag *Tag) BeforeUpdate(tx *gorm.DB) (err error) {
	tag.ModifiedOn = time.Now().Unix()
	return nil
}
