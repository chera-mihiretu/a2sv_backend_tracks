package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Get a list of tasks
	router.GET("/tasks", nil)
	// to get the task detail
	router.GET("/tasks/:id", nil)
	//  to update the task
	router.PUT("/tasks/:id", nil)
	// to delete the task
	router.DELETE("tasks/:id", nil)

	// to add new task
	router.PUT("/tasks", nil)
	router.Run()
}
