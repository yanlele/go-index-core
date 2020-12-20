package services

import (
	"fmt"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"log"
	"time"
)

// insertArchive 将发布的文章归档
func SetArticleArchive(article *modules.Article) {
	var archive modules.Archive
	layout := "2006-01-02 03:04:05"

	archiveDateParse, err := time.Parse(layout, article.CreatedAt.Format(layout))
	if err != nil {
		return
	}

	theArticleArchiveDate := archiveDateParse.Format("2020-01")
	driver.DB.Where("archive_date = ?", theArticleArchiveDate).First(&archive)
	if archive.ID == 0 {
		// 创建
		archive.ArchiveDate = archiveDateParse.Format("2020-01")
		archive.ArticleIds = fmt.Sprintf("%d", article.ID)
		err := driver.DB.Create(&archive).Error
		if err != nil {
			log.Println(err)
		}
		return
	}


}
