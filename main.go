package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mvrsss/go-todo-api/app"
	"github.com/mvrsss/go-todo-api/app/migrate"
	_ "github.com/mvrsss/go-todo-api/app/models"
	"github.com/mvrsss/go-todo-api/config"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.GetConfig().DB))
	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()
	migrate.DBMigrate(config.DB)
	r := app.SetRouters()
	r.Run()
}
