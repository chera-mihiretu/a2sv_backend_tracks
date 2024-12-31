package services

import "github/chera/task_manager/models"

type TaskManager interface {
	OpennFile() error
	AddTask() error
	RemoveTask() error
	UpdateTask() error
	GetTask() models.Taks
	GetTasks() []models.Taks
}

// TaskService struct that open file
type TaskService struct {
	AllTaks map[int]models.Taks
}

// TaskService struct that add new task
func (t *TaskService) AddTask() error {
	return nil
}

// TaskService struct that remove task
func (t *TaskService) RemoveTask() error {
	return nil
}

// TaskService struct that update task
func (t *TaskService) UpdateTask() error {
	return nil
}

// TaskService struct that get task
func (t *TaskService) GetTask() models.Taks {
	return models.Taks{}
}

// TaskService struct that get tasks
func (t *TaskService) GetTasks() []models.Taks {
	return []models.Taks{}
}
