package tag_service

import (
	"encoding/json"
	"gin-example/models"
	"gin-example/pkg/gredis"
	"gin-example/pkg/logging"
	"gin-example/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistById() (bool, error) {
	return models.ExistTagById(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State > 0 {
		data["state"] = t.State
	}
	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int64, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)

	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}

	key := cache.GetTagsKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Warn(err)
		} else {
			_ = json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	return tags, err
}

//func (t *Tag) Export() (string, error) {
//	tags, err := t.GetAll()
//	if err != nil {
//		return "", err
//	}
//
//	xlsFile:=
//}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if t.Name != "" {
		maps["name"] = t.Name
	}

	if t.State > 0 {
		maps["state"] = t.State
	}

	return maps
}
