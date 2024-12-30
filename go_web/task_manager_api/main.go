package main

import (
	"github/cheramihiretu/task_manager_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	router := gin.Default()

	// Get a list of tasks
	router.GET("/tasks", getAllTask)
	// to get the task detail
	router.GET("/tasks/:id", getSpecificTask)
	//  to update the task accepts JSON
	router.PUT("/tasks/:id", updateTask)
	// to delete the task
	router.DELETE("tasks/:id", deleteTask)

	// to add new task, accepts JSON
	router.PUT("/tasks", nil)

	router.Run()
}

// returns all the tasks
func getAllTask(c *gin.Context) {
	c.IndentedJSON(200, tasks)
}

// returns a specific task based on the id
func getSpecificTask(c *gin.Context) {
	id := c.Param("id")

	for _, val := range tasks {
		if val.ID == id {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "task not found")
}

// Updates the task given with id
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var index_to_update int = -1

	var updated_task models.Task

	if err := c.BindJSON(&updated_task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// update the task
	for i, task := range tasks {
		if task.ID == id {
			index_to_update = i
			break
		}
	}
	if index_to_update == -1 {
		c.IndentedJSON(http.StatusNotFound, "task not found")
		return
	}

	// update the task
	if updated_task.Title != "" {
		tasks[index_to_update].Title = updated_task.Title
	}

	if updated_task.Description != "" {
		tasks[index_to_update].Description = updated_task.Description
	}

	if updated_task.Status != "" {
		tasks[index_to_update].Status = updated_task.Status
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
	})

}

// delete task with given id
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Task deleted successfully",
			})
			return
		}
	}
}
