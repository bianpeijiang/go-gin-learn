package models

import (
	"github.com/bytedance/gopkg/util/logger"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (a *Article) BeforeCreate(db *gorm.DB) error {
	now := time.Now().Unix()
	db.Statement.SetColumn("CreatedOn", now)
	db.Statement.SetColumn("UpdatedOn", now)
	return nil
}

func (a *Article) BeforeUpdate(db *gorm.DB) error {
	db.Statement.SetColumn("UpdatedOn", time.Now().Unix())
	return nil
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (int64, error) {
	var count int64
	result := db.Model(&Article{}).Where(maps).Count(&count)
	return count, result.Error
}

func GetArticles(pageNum int, pageSize int, maps interface{}) ([]Article, error) {
	var articles []Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil {
		logger.Errorf("GetArticles query err:%v", err)
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Preload("Tag").Where("id=?", id).First(&article).Error
	if err != nil {
		logger.Errorf("GetArticle query err:%v", err)
		return nil, err
	}
	return &article, nil
}

func EditArticle(id int, data interface{}) bool {
	err := db.Model(&Article{}).Where("id=?", id).Updates(data).Error
	if err != nil {
		logger.Errorf("EditArticle query err:%v", err)
		return false
	}
	return true
}

func AddArticle(data map[string]interface{}) (int64, error) {
	article := Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}
	result := db.Create(&article)
	if result.Error != nil {
		logger.Errorf("AddArticle query err:%v", result.Error)
		return 0, result.Error
	}
	return article.ID, nil
}

func DeleteArticle(id int) bool {
	article := Article{}
	err := db.Where("id=?", id).Delete(&article).Error
	if err != nil {
		logger.Errorf("DeleteArticle query err:%v", err)
		return false
	}
	return true
}
