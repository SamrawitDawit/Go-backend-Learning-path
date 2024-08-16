package router

import (
	controller "task8-Testing/Delivery/controllers"
	domain "task8-Testing/Domain"
	inf "task8-Testing/Infrastructure"
	repositories "task8-Testing/Repositories"
	usecases "task8-Testing/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(timeout time.Duration, db mongo.Database, jwtService inf.Jwt_interface) {
	router := gin.Default()
	tr := repositories.NewTaskRepository(db, domain.CollectionTask)
	ur := repositories.NewUserRepository(db, domain.CollectionUser)

	controller := &controller.Controller{
		TaskUsecase: usecases.NewTaskUsecase(tr, timeout),
		UserUsecase: usecases.NewUserUsecase(ur, timeout),
	}
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.GET("/tasks", inf.AuthMiddleWare(jwtService), controller.GetTasks)
	router.GET("/tasks/:_id", inf.AuthMiddleWare(jwtService), controller.GetTask)
	router.POST("/users/:_id", inf.AuthMiddleWare(jwtService), inf.AdminMiddleWare(), controller.Promote)
	router.POST("/tasks", inf.AuthMiddleWare(jwtService), inf.AdminMiddleWare(), controller.AddTask)
	router.PUT("/tasks/:_id", inf.AuthMiddleWare(jwtService), inf.AdminMiddleWare(), controller.UpdateTask)
	router.DELETE("/tasks/:_id", inf.AuthMiddleWare(jwtService), inf.AdminMiddleWare(), controller.RemoveTask)
	router.Run(":8080")

}
