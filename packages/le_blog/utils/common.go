package utils

import (
	"fmt"
	"html/template"
	"le-blog/config"
	"strings"
)

// Html 非转义输出html
func Html(str string) interface{} {
	return template.HTML(str)
}

// TagString2Map 将文章tags转换为[]string
func TagString2Map(tagString string) []string {
	var tagSlice []string
	tagSlice = strings.Split(tagString, ",")

	for index, value := range tagSlice {
		tagSlice[index] = strings.Trim(value, " ")
	}

	return tagSlice
}

// SetLinkTitle 设置a链接的title
func SetLinkTitle(title string) string {
	section, err := config.Config.GetSection("env")
	if err != nil {
		return ""
	}

	appName, _ := section.GetKey("AppName")
	return fmt.Sprintf("%s - %s", title, appName)
}

// AppUrl 获取站点配置的url
func AppUrl(path string) string {
	section, err := config.Config.GetSection("env")
	if err != nil {
		return fmt.Sprintf("%s", path)
	}
	appUrl, err := section.GetKey("AppUrl")
	if err != nil {
		return fmt.Sprintf("%s", path)
	}

	return fmt.Sprintf("%s%s", appUrl, path)
}

// SocialHtml 返回社交html
func SocialHtml() string {
	var socialHtml string
	socialSection, err := config.Config.GetSection("social")
	if err != nil {
		return socialHtml
	}

	// 定义映射关系
	socialMap := make(map[string]string)
	socialMap["github"] = `<i class="fab fa-github" aria-hidden="true"></i>`
	socialMap["twitter"] = `<i class="fab fa-twitter" aria-hidden="true"></i>`
	socialMap["linkedin"] = `<i class="fab fa-linkedin" aria-hidden="true"></i>`
	socialMap["stack"] = `<i class="fab fa-stack-overflow" aria-hidden="true"></i>`
	socialMap["codepen"] = `<i class="fab fa-codepen" aria-hidden="true"></i>`

	// <li class="list-inline-item"><a href="#"><i class="fab fa-twitter fa-fw"></i></a></li>

	for _, key := range socialSection.KeyStrings() {
		val, _ := socialSection.GetKey(key)
		if keyHtml,ok := socialMap[key]; ok {
			socialHtml = fmt.Sprintf(`%s<li class="list-inline-item"><a href="%s">%s</a></li>`, socialHtml, val, keyHtml)
		} else {
			socialHtml = fmt.Sprintf(`%s<li class="list-inline-item"><a href="%s">%s</a></li>`, socialHtml, val, key)
		}
	}

	return socialHtml
}


