package router

import (
	controller "task5-Task_Mangement_API_using_MongoDB_and_Go_driver/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Controller *controller.Controller
}

func (r *Router) Route() {
	router := gin.Default()
	router.GET("/tasks", r.Controller.GetTasks)
	router.GET("/tasks/:id", r.Controller.GetTask)
	router.POST("/tasks", r.Controller.AddTask)
	router.PUT("/tasks/:id", r.Controller.UpdateTask)
	router.DELETE("/tasks/:id", r.Controller.RemoveTask)
	router.Run(":8080")
}
