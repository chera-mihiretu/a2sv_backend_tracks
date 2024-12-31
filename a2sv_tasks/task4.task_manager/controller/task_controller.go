package controller

import (
	services "github/chera/task_manager/data"
	"github/chera/task_manager/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskControllerInterface interface {
	OpenFile() error
	AddTask() error
	RemoveTask() error
	UpdateTask() error
	GetTask() models.Tasks
	GetTasks()
}

// the struct wich is going to implement the interface
type TaskController struct {
	TaskService *services.TaskService
}

func (tc *TaskController) OpenFile() error {
	return tc.TaskService.OpenFile()

}

// the function that returns the struct
func NewTaskController() *TaskController {
	return &TaskController{
		TaskService: services.NewTaskService(),
	}
}

// the functions that implement the interface
func (tc *TaskController) AddTask(c *gin.Context) {
	var task models.Tasks

	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "bad json body",
		})
		return
	}

	if err := tc.TaskService.AddTask(task); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Written Succesfully",
	})
}

// the functions that implement the interface
func (tc *TaskController) RemoveTask(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
	}

	if err := tc.TaskService.RemoveTask(int_id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}

}

// the functions that implement the interface
func (tc *TaskController) UpdateTask(c *gin.Context) {
	var task models.Tasks
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad json Format",
		})
		return
	}

	if err := tc.TaskService.UpdateTask(task); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "updated succesfully"})
}

// the functions that implement the interface
func (tc *TaskController) GetTask() models.Tasks {
	return models.Tasks{}
}

// the functions that implement the interface
func (tc *TaskController) GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tc.TaskService.GetTasks())
}
