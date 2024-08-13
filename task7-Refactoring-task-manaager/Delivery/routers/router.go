package router

import (
	controller "task7-Refactoring-task-manaager/Delivery/controllers"
	domain "task7-Refactoring-task-manaager/Domain"
	inf "task7-Refactoring-task-manaager/Infrastructure"
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

	controller := &controller.Controller{
		TaskUsecase: usecases.NewTaskUsecase(tr, timeout),
		UserUsecase: usecases.NewUserUsecase(ur, timeout),
	}
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.GET("/tasks", inf.AuthMiddleWare(), controller.GetTasks)
	router.GET("/tasks/:_id", inf.AuthMiddleWare(), controller.GetTask)
	router.POST("users/:_id", inf.AuthMiddleWare(), inf.AdminMiddleWare(), controller.Promote)
	router.POST("/tasks", inf.AuthMiddleWare(), inf.AdminMiddleWare(), controller.AddTask)
	router.PUT("/tasks/:_id", inf.AuthMiddleWare(), inf.AdminMiddleWare(), controller.UpdateTask)
	router.DELETE("/tasks/:_id", inf.AuthMiddleWare(), inf.AdminMiddleWare(), controller.RemoveTask)
	router.Run(":8080")

}
