package data

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github/chera/task_manager/models"
	"os"
	"strconv"
	"time"
)

type TaskManager interface {
	OpenFile() error
	AddTask() error
	RemoveTask() error
	UpdateTask() error
	GetTask() models.Tasks
	GetTasks() []models.Tasks
}

// TaskService struct that open file
type TaskService struct {
	ID       int
	AllTasks map[int]models.Tasks
}

func writeFile(t *TaskService) error {
	file, err := os.Create("task.csv")

	if err != nil {
		return errors.New("cant write the fild")
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	for _, element := range t.AllTasks {
		row := []string{
			fmt.Sprintf("%v", element.ID),
			fmt.Sprintf("%v", element.Title),
			fmt.Sprintf("%v", element.Description),
			fmt.Sprintf("%v", element.Status),
			fmt.Sprintf("%v", element.CreatedAt),
			fmt.Sprintf("%v", element.UpdatedAt),
		}
		if err := writer.Write(row); err != nil {
			return errors.New("couldn't write all the data")
		}
	}

	writer.Flush()
	return nil
}

// file existance check
func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
func NewTaskService() *TaskService {
	return &TaskService{
		ID:       0,
		AllTasks: make(map[int]models.Tasks),
	}
}
func (t *TaskService) OpenFile() error {
	exist := fileExist("task.csv")
	var file *os.File
	var err error
	if !exist {
		file, err = os.Create("task.csv")
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		file, err = os.OpenFile("task.csv", os.O_RDWR, 0644)
		if err != nil {
			return err
		}

		defer file.Close()
	}

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()

	if err != nil {
		return err
	}
	layout := "2006-01-02 15:04:05 -0700 MST"
	for _, row := range rows {

		id, _ := strconv.Atoi(row[0])

		createdAt, err_time := time.Parse(layout, row[4])
		updatedAt, err_time_1 := time.Parse(layout, row[5])
		if err_time != nil || err_time_1 != nil {
			return err_time
		}
		t.AllTasks[id] = models.Tasks{
			ID:          id,
			Title:       row[1],
			Description: row[2],
			Status:      row[3],
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
	}

	t.ID += len(rows)

	return nil
}

// TaskService struct that add new task and also write to file
func (t *TaskService) AddTask(task models.Tasks) error {
	task.ID = t.ID
	t.AllTasks[t.ID] = task
	t.ID++

	return writeFile(t)
}

// TaskService struct that remove task and also write to file
func (t *TaskService) RemoveTask(id int) error {
	if _, exist := t.AllTasks[id]; !exist {
		return errors.New("no task with this id")
	}
	t.ID--
	task := t.AllTasks[t.ID]
	task.ID = id
	t.AllTasks[id] = task

	delete(t.AllTasks, t.ID)

	return writeFile(t)
}

// TaskService struct that update task, and also write to file
func (t *TaskService) UpdateTask(task models.Tasks) error {

	if _, exist := t.AllTasks[task.ID]; !exist {
		return errors.New("no task with this id")
	}

	new_task := t.AllTasks[task.ID]
	new_task.UpdatedAt = time.Now()
	if task.Title != "" {

		new_task.Title = task.Title
	}
	if task.Status != "" {

		new_task.Status = task.Status
	}
	if task.Description != "" {

		new_task.Description = task.Description
	}

	t.AllTasks[new_task.ID] = new_task

	return writeFile(t)
}

// TaskService struct that get task by id and also write to file
func (t *TaskService) GetTask() models.Tasks {
	return models.Tasks{}
}

// TaskService struct that get tasks and also write to file
func (t *TaskService) GetTasks() []models.Tasks {
	var current []models.Tasks = make([]models.Tasks, 0)
	for _, value := range t.AllTasks {
		current = append(current, value)
	}
	return current
}
