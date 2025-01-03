package router

import (
	"github/chera/task_manager/controller"
	"github/chera/task_manager/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routers(taskCollection, userCollection *mongo.Collection) {
	router := gin.Default()
	taskController := controller.NewTaskController(taskCollection, userCollection)
	myMiddleWare := middleware.NewMiddleWare()

	router.POST("/register", taskController.RegisterUser)
	router.POST("/login", taskController.LoginUser)
	router.GET("/tasks", myMiddleWare.CheckValidity("user"), taskController.GetTasks)
	router.GET("/tasks/:id", myMiddleWare.CheckValidity("user"), taskController.GetTasks)

	router.POST("/tasks", myMiddleWare.CheckValidity("admin"), taskController.AddTask)
	router.DELETE("/tasks/:id", myMiddleWare.CheckValidity("admin"), taskController.RemoveTask)
	router.PUT("/tasks", myMiddleWare.CheckValidity("admin"), taskController.UpdateTask)
	router.Run()
}
