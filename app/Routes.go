package app

import (
	"github.com/mvrsss/go-todo-api/app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/mvrsss/go-todo-api/config"
)

func SetRouters() *ginEngine {
	r := gin.Default()

	v1 := r.Group("/v1") {
		v1.GET("todo", controllers.GetAllTasks),
		v1.POST("todo", controllers.AddTask),
		v1.GET("todo/:id", controllers.GetATask),
		v1.PUT("todo/:id", controllers.UpdateTasks),
		v1.DELETE("todo/:id", controllers.DeleteTasks)
	}

	return r
}