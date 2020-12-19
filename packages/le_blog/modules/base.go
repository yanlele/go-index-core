package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"net/url"
	"strconv"
)

// Paginator 接口
type Paginator interface {
	Paginate(c *gin.Context, out interface{}, query interface{}, args ...interface{}) *Pagination
}

// 定义每页的PageLink
type link struct {
	P    int
	Link string
}

//Paginate 分页器
type Pagination struct {
	Total         int    // 总条数
	TotalPage     int    // 总页数
	PerPage       int    // 每页的数量
	CurrentPage   int    // 当前第几页
	ShowLastPage  bool   // 是否最后一页
	ShowFirstPage bool   // 是否第一页
	LastPageUrl   string // 最后一页url
	FirstPageUrl  string // 第一页的url
	Path          string
	PageLink      []link      // 每页的链接
	Data          interface{} // 返回的数据
}

// Paginate
// table 查询数据库的名称
// c gin.context 上下文
// out interface{} 数据填充
// query gorm query string
// args  gorm query args
func (pg *Pagination) Paginate(c *gin.Context, db gorm.DB, out interface{}) (*Pagination, error) {
	// 当前请求的url
	requestURL := fmt.Sprintf("%s", c.Request.URL)
	size, _ := strconv.Atoi(c.Query("size"))
	if size == 0 {
		size = 10
	}
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * size

	// 计算总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = db.Offset(offset).Limit(size).Find(out).Error
	if err != nil {
		return nil, err
	}

	// 总页数
	totalPage := int(math.Ceil(float64(total) / float64(size)))

	// 第一页
	showFirstPage := true
	if page == 1 {
		showFirstPage = false
	}

	// 最一页
	showLastPage := true
	if page == totalPage {
		showLastPage = false
	}

	// template 输出控制
	pages := make([]int, totalPage)
	if totalPage <= 5 {
		for index := range pages {
			pages[index] = index + 1
		}
	} else {
		if page <= 3 {
			for index := range pages {
				pages[index] = index + 1
			}
		} else {
			if page >= (totalPage - 2) {
				page := totalPage - 4
				for index := range pages {
					page := page + index
					pages[index] = page
				}
			} else {
				page := page - 2
				for index := range pages {
					pages[index] = page + index
				}
			}
		}
	}

	pageLink := make([]link, totalPage)
	for index, page := range pages {
		pLink := link{
			P:    page,
			Link: getPaginateItemUrl(requestURL, page),
		}
		pageLink[index] = pLink
	}

	return &Pagination{
		Total:         int(total),
		TotalPage:     totalPage,
		PerPage:       size,
		CurrentPage:   page,
		ShowFirstPage: showFirstPage,
		ShowLastPage:  showLastPage,
		LastPageUrl:   getPaginateItemUrl(requestURL, totalPage),
		FirstPageUrl:  getPaginateItemUrl(requestURL, 1),
		Path:          "",
		PageLink:      pageLink,
		Data:          out,
	}, nil
}

func getPaginateItemUrl(requestUrl string, page int) string {
	uParse, _ := url.Parse(requestUrl)
	query := uParse.Query()
	query.Set("page", strconv.Itoa(page))
	uParse.RawQuery = query.Encode()
	return fmt.Sprintf("%s", uParse)
}
