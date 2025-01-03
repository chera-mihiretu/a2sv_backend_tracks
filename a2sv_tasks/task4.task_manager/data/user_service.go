package data

import (
	"github/chera/task_manager/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserManager interface {
	RegisterUser(user models.User) error
	LoginUser(user models.User) (string, error)
}

type UserService struct {
	Collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{Collection: collection}
}

func (us *UserService) RegisterUser(user models.User) error {
	return nil
}

func (us *UserService) LoginUser(user models.User) (string, error) {
	return "", nil
}
