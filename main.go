package main

import (
	"fmt"
	"github.com/mvrsss/go-todo-api/config"
	"github.com/mvrsss/go-todo-api/app"
	"github.com/mvrsss/go-todo-api/app/models"
	"github.com/mvrsss/go-todo-api/app/migrate"
	"github.com/jinzhu/gorm"
)

func main() {
	Config.DB, err := gorm.Open("mysql", Config.DbURL(Config.GetConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()
	migrate.DBMigrate(Config.DB)
	r := app.SetRouters()
	r.Run()
}