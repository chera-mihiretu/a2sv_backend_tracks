package services

import "github/chera/task_manager/models"

type TaskServices interface {
	AddTask() error
	RemoveTask() error
	UpdateTask() error
	GetTask() models.Taks
	GetTasks() []models.Taks
}
