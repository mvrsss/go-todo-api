package migrate

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mvrsss/go-todo-api/app/models"
)

// type Todo struct {
// 	gorm.Model
// 	Key      int    `gorm:"column:key" json:"key"`
// 	Semester string `gorm:"column:semester" json:"semester"`
// 	Grade    string `gorm:"column:grade" json:"grade"`
// }

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Todo{})
	isPresent := db.HasTable(&models.Todo{})
	fmt.Println("Table course is present", isPresent)
	if !isPresent {
		db.CreateTable(models.Todo{})
	}

	return db
}
