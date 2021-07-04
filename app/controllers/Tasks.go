package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mvrsss/go-todo-api/app/models"
)

func GetATask(a *gin.Context) {
	id := a.Params.ByName("key")
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
	err := models.AddTask(&t)
	if err != nil {
		a.AbortWithStatus(http.StatusNotFound)
	} else {
		a.JSON(http.StatusOK, t)
	}
}

func UpdateTasks(a *gin.Context) {
	var t models.Todo
	id := a.Params.ByName("key")
	err := models.GetATask(&t, id)
	if err != nil {
		a.JSON(http.StatusNotFound, t)
	}
	a.BindJSON(&t)
	err = models.UpdateTasks(&t, id)
	if err != nil {
		a.AbortWithStatus(http.StatusNotFound)
	} else {
		a.JSON(http.StatusOK, t)
	}
}

func DeleteTasks(a *gin.Context) {
	var t models.Todo
	id := a.Params.ByName("key")
	err := models.DeleteTasks(&t, id)
	if err != nil {
		a.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
