package router

import (
	"github/chera/task_manager/controller"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routers(collection *mongo.Collection) {
	router := gin.Default()
	taskController := controller.NewTaskController(collection)

	router.GET("/tasks", taskController.GetTasks)
	router.GET("/tasks/:id", taskController.GetTasks)

	router.POST("/tasks", taskController.AddTask)
	router.DELETE("/tasks/:id", taskController.RemoveTask)
	router.PUT("/tasks", taskController.UpdateTask)
	router.Run()
}
