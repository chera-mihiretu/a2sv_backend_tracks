package router

import (
	"fmt"
	"github/chera/task_manager/controller"

	"github.com/gin-gonic/gin"
)

func Routers() {
	router := gin.Default()
	taskController := controller.NewTaskController()
	err := taskController.OpenFile()
	if err != nil {
		fmt.Println(err)
	}
	router.GET("/tasks", taskController.GetTasks)
	router.POST("/tasks", taskController.AddTask)
	router.DELETE("/tasks/:id", taskController.RemoveTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.Run()
}
