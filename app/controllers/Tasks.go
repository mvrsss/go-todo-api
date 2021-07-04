package controllers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mvrsss/go-todo-api/app/models"
)

func GetATask(a *gin.Context) {
	id, _ := strconv.Atoi(a.Params.ByName("key"))
	var t models.Todo
	var db *gorm.DB
	err := models.GetATask(&t, id, db)
	if err != nil {
		a.AbortWithStatus(http.StatusNotFound)
	} else {
		a.JSON(http.StatusOK, t)
	}
}

func GetAllTasks(a *gin.Context) {
	var t []models.Todo
	var db *gorm.DB
	err := models.GetAllTasks(db, &t)
	if err != nil {
		a.AbortWithStatus(http.StatusNotFound)
	} else {
		a.JSON(http.StatusOK, t)
	}
}

func AddTask(a *gin.Context) {
	var t models.Todo
	a.BindJSON(&t)
	var db *gorm.DB
	err := models.AddTask(db, &t)
	if err != nil {
		a.AbortWithStatus(http.StatusNotFound)
	} else {
		a.JSON(http.StatusOK, t)
	}
}

func UpdateTasks(a *gin.Context) {
	var t models.Todo
	id, _ := strconv.Atoi(a.Params.ByName("key"))
	var db *gorm.DB
	err := models.GetATask(&t, id, db)
	if err != nil {
		a.JSON(http.StatusNotFound, t)
	}
	a.BindJSON(&t)
	err = models.UpdateTasks(&t, id, db)
	if err != nil {
		a.AbortWithStatus(http.StatusNotFound)
	} else {
		a.JSON(http.StatusOK, t)
	}
}

func DeleteTasks(a *gin.Context) {
	var t models.Todo
	id, _ := strconv.Atoi(a.Params.ByName("key"))
	id1 := strconv.Itoa(id)
	var db *gorm.DB
	err := models.DeleteTasks(&t, id, db)
	if err != nil {
		a.AbortWithStatus(http.StatusNotFound)
	} else {
		a.JSON(http.StatusOK, gin.H{"id:" + id1: "deleted"})
	}
}
