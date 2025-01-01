package controller

import (
	"github/chera/task_manager/data"
	"github/chera/task_manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskControllerInterface interface {
	AddTask() error
	RemoveTask() error
	UpdateTask() error
	GetTask() models.Tasks
	GetTasks()
}

// the struct wich is going to implement the interface
type TaskController struct {
	TaskService *data.TaskService
}

// the function that returns the struct
func NewTaskController(collection *mongo.Collection) *TaskController {
	return &TaskController{
		TaskService: data.NewTaskService(collection),
	}
}

// the functions that implement the interface
func (tc *TaskController) AddTask(c *gin.Context) {
	var task models.Tasks

	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	result, err := tc.TaskService.AddTask(task)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// the functions that implement the interface
func (tc *TaskController) RemoveTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.TaskService.RemoveTask(id); err != nil {
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
	result, err := tc.TaskService.UpdateTask(task)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// the functions that implement the interface
func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	result, err := tc.TaskService.GetTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// the functions that implement the interface
func (tc *TaskController) GetTasks(c *gin.Context) {
	result, err := tc.TaskService.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
