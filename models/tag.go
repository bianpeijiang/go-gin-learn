package models

import (
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

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	tx.Statement.SetColumn("CreatedOn", now)
	tx.Statement.SetColumn("UpdatedOn", now)
	return nil
}

func (t *Tag) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.UpdateColumn("UpdatedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, int64, error) {
	tags := make([]Tag, 0)
	result := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return tags, result.RowsAffected, result.Error
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) (int64, error) {
	tags := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	result := db.Create(&tags)

	return tags.ID, result.Error
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data map[string]interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}
