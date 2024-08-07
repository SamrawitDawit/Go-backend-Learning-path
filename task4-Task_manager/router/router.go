package router

import (
	"task4-Task_manager/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	controller controllers.Controller
}

func (r *Router) Route() {
	router := gin.Default()
	router.GET("/tasks", r.controller.GetTasks)
	router.GET("/tasks/:id", r.controller.GetTask)
	router.POST("/tasks", r.controller.AddTask)
	router.PUT("/tasks/:id", r.controller.UpdateTask)
	router.DELETE("/tasks/:id", r.controller.RemoveTask)
	router.Run(":8080")
}
