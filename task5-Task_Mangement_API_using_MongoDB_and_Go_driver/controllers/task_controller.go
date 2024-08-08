package controllers

import (
	"net/http"
	service "task5-Task_Mangement_API_using_MongoDB_and_Go_driver/data"
	"task5-Task_Mangement_API_using_MongoDB_and_Go_driver/models"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TaskService service.TaskService
}

func (controller *Controller) GetTasks(context *gin.Context) {
	tasks := controller.TaskService.GetTasks()
	context.IndentedJSON(http.StatusOK, tasks)
}
func (controller *Controller) GetTask(context *gin.Context) {
	id := context.Param("id")
	task, err := controller.TaskService.GetTask(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	}

	context.JSON(http.StatusOK, task)

}
func (controller *Controller) AddTask(context *gin.Context) {
	var newTask models.Task

	if err := context.BindJSON(&newTask); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.TaskService.AddTask(newTask)
	context.JSON(http.StatusCreated, gin.H{"message": "Task Created"})
}
func (controller *Controller) RemoveTask(context *gin.Context) {
	id := context.Param("id")
	err := controller.TaskService.RemoveTask(id)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"message": "Task removed"})
		return
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
}
func (controller *Controller) UpdateTask(context *gin.Context) {
	var updatedTask models.Task
	id := context.Param("id")
	if err := context.BindJSON(&updatedTask); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orginal_task, err := controller.TaskService.GetTask(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	if updatedTask.Title == "" {
		updatedTask.Title = orginal_task.Title
	}
	if updatedTask.Description == "" {
		updatedTask.Description = orginal_task.Description
	}
	if updatedTask.Status == "" {
		updatedTask.Status = orginal_task.Status
	}
	if updatedTask.DueDate.IsZero() {
		updatedTask.DueDate = orginal_task.DueDate
	}
	errr := controller.TaskService.UpdateTask(id, updatedTask)
	if errr == nil {
		context.JSON(http.StatusOK, gin.H{"message": "Task Updated"})
		return
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
