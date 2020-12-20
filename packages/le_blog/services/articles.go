package services

import (
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"log"
	"strings"
)

func HandleTags(tag string) bool {
	var tagSlice []string
	tagSlice = strings.Split(tag, ",")
	var tags []modules.Tag
	err := driver.DB.Where("name in (?)", tagSlice).Find(&tag).Error
	if err != nil {
		log.Println(err)
		return false
	}
}
