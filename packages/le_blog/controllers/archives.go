package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Archives(c *gin.Context) {
	auth := Auth{}.GetAuth(c)
	var archives []modules.Archive
	err := driver.DB.Order("created_at desc").Find(&archives).Error

	if err != nil {
		log.Fatalln(err)
	}

	type articleItems []modules.Article
	Archives := make(map[string]articleItems)
	if len(archives) > 0 {
		for _, archive := range archives {
			// 查找文章
			var ids []int
			for _, id := range strings.Split(archive.ArticleIds, ",") {
				id, _ := strconv.Atoi(id)
				ids = append(ids, id)
			}
			var Items articleItems
			driver.DB.Table("articles").Where("id in (?)", ids).Find(&Items)
			Archives[archive.ArchiveDate] = Items
		}
	}

	for _, archive := range Archives {
		for _, item := range archive {
			fmt.Printf("%#v", item)
		}
	}

	header := Header{"文件归档"}

	data := struct {
		Auth
		Archives map[string]articleItems
		Header   Header
	}{auth, Archives, header}

	c.HTML(http.StatusOK, "archives", data)
}
