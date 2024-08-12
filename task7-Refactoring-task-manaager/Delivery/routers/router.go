package router

import (
	controller "task7-Refactoring-task-manaager/Delivery/controllers"
	domain "task7-Refactoring-task-manaager/Domain"
	infrastructure "task7-Refactoring-task-manaager/Infrastructure"
	repositories "task7-Refactoring-task-manaager/Repositories"
	usecases "task7-Refactoring-task-manaager/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(timeout time.Duration, db mongo.Database) {
	router := gin.Default()
	tr := repositories.NewTaskRepository(db, domain.CollectionTask)
	ur := repositories.NewUserRepository(db, domain.CollectionUser)
	md := infrastructure.MiddleWare{}

	controller := &controller.Controller{
		TaskUsecase: usecases.NewTaskUsecase(tr, timeout),
		UserUsecase: usecases.NewUserUsecase(ur, timeout),
	}
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.GET("/tasks", md.AuthMiddleWare(), controller.GetTasks)
	router.GET("/tasks/:_id", md.AuthMiddleWare(), controller.GetTask)
	router.POST("users/:_id", md.AuthMiddleWare(), md.AdminMiddleWare(), controller.Promote)
	router.POST("/tasks", md.AuthMiddleWare(), md.AdminMiddleWare(), controller.AddTask)
	router.PUT("/tasks/:_id", md.AuthMiddleWare(), md.AdminMiddleWare(), controller.UpdateTask)
	router.DELETE("/tasks/:_id", md.AuthMiddleWare(), md.AdminMiddleWare(), controller.RemoveTask)
	router.Run(":8080")

}
