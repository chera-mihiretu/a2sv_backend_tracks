package data

import (
	"context"
	"fmt"
	"github/chera/task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskManager interface {
	AddTask() error
	RemoveTask() error
	UpdateTask() error
	GetTask() models.Tasks
	GetTasks() []models.Tasks
}

// TaskService struct that open file
type TaskService struct {
	Collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{
		Collection: collection,
	}
}

// TaskService struct that add new task and also write to file
func (t *TaskService) AddTask(task models.Tasks) (models.Tasks, error) {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := t.Collection.InsertOne(context.Background(), task)
	if err != nil {
		return models.Tasks{}, err
	}
	return task, nil
}

// TaskService struct that remove task and also write to file
func (t *TaskService) RemoveTask(id string) error {
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = t.Collection.DeleteOne(context.Background(), bson.M{"id": taskID})
	if err != nil {
		return err
	}
	return nil
}

// TaskService struct that update task, and also write to file
func (t *TaskService) UpdateTask(task models.Tasks) (models.Tasks, error) {

	// check if the data exist
	taskID, err := primitive.ObjectIDFromHex(task.ID.Hex())
	if err != nil {
		return models.Tasks{}, err
	}
	var old_task models.Tasks
	err = t.Collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: taskID}}).Decode(&old_task)
	if err != nil {
		return models.Tasks{}, err
	}
	if task.Title == "" {
		task.Title = old_task.Title
	}
	if task.Description == "" {
		task.Description = old_task.Description
	}
	if task.Status == "" {
		task.Status = old_task.Status
	}
	fmt.Println(time.Now())
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
			"updated_at":  time.Now(),
		},
	}

	_, err = t.Collection.UpdateOne(context.TODO(), bson.D{{Key: "id", Value: taskID}}, update)

	if err != nil {
		return models.Tasks{}, err
	}

	return task, nil
}

// TaskService struct that get task by id and also write to file
func (t *TaskService) GetTask(taskID string) (models.Tasks, error) {
	taskIDHex, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return models.Tasks{}, err
	}
	var task models.Tasks
	err = t.Collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: taskIDHex}}).Decode(&task)
	if err != nil {
		return models.Tasks{}, err
	}
	return task, nil
}

// TaskService struct that get tasks and also write to file
func (t *TaskService) GetTasks() ([]models.Tasks, error) {
	var current []models.Tasks = make([]models.Tasks, 0)
	cursor, err := t.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &current); err != nil {
		return nil, err
	}
	return current, nil
}
