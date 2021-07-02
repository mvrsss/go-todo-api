package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	gorm.Model
	Key         int    `gorm:"column:key"`
	ID          int    `gorm:"column:id"json:"id"`
	Title       string `gorm:"column:title"json:"title"`
	Description string `gorm:"column:description"json:"description"`
}

func (t *Todo) TableName() string {
	return "todo"
}

func AddTask(db *gorm.DB, t *Todo) (err error) {
	if err = db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTasks(db *gorm.DB, t *[]Todo) (err error) {
	if err = db.Order("id desc").Find(t).Error; err != nil {
		return err
	}
	return nil
}

func GetATask(t *Todo, id int, db *gorm.DB) (err error) {
	if err = db.Where("key = ?", id).First(t).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTasks(t *Todo, id int, db *gorm.DB) (err error) {
	if err = db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTasks(t *Todo, id int, db *gorm.DB) (err error) {
	if err = db.Where("key = ?", id).Delete(t).Error; err != nil {
		return err
	}
	return nil
}
