package services

import (
	"gorm.io/gorm"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"log"
	"strings"
)

// handleTags 处理文章的tags
func HandleTags(tag string) bool {
	var tagSlice []string
	tagSlice = strings.Split(tag, ",")
	var tags []modules.Tag
	err := driver.DB.Where("name in (?)", tagSlice).Find(&tag).Error
	if err != nil {
		log.Println(err)
		return false
	}

	tagMap := make(map[string]string)

	for _, ts := range tagSlice {
		tagMap[ts] = ts
	}

	var tagsStructSlice []modules.Tag

	for _, tm := range tagMap {
		var ts = modules.Tag{
			Name:   tm,
			UseNum: 1,
		}

		for _, tag := range tags {
			if tm == tag.Name {
				ts = modules.Tag{
					Model: gorm.Model{
						ID: tag.ID,
					},
					Name:   tag.Name,
					UseNum: tag.UseNum + 1,
				}
			}
		}
		tagsStructSlice = append(tagsStructSlice, ts)
	}

	for _, tag := range tagsStructSlice {
		var err error
		if tag.ID == 0 {
			err = driver.DB.Create(&tag).Error
			if err != nil {
				panic(err)
			}
		} else {
			err = driver.DB.Model(modules.Tag{}).Updates(tag).Error
		}

		if err != nil {
			return false
		}
	}
	return true
}
