package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Model

	TagID         int    `json:"tag_id" gorm:"index"`
	Tag           Tag    `json:"tag"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

/* 通过 id 判定是否存在 */
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

/* 获取文章总数 */
func GetArticleTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

/* 获取所有文章 */
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, err
}

/*
获取单个文章

实体关联：https://gorm.io/zh_CN/docs/associations.html
*/
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	_ = db.Model(&article).Association("tag").Find(&article.Tag)
	return
}



func (article *Article) BeforeCreate() error {
	article.CreatedOn = time.Now().Unix()
	return nil
}

func (article *Article) BeforeUpdate() error {
	article.ModifiedOn = time.Now().Unix()
	return nil
}
