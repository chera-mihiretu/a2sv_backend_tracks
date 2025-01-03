package data

import (
	"context"
	"errors"
	"fmt"
	"github/chera/task_manager/middleware"
	"github/chera/task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserManager interface {
	RegisterUser(user models.User) error
	LoginUser(user models.User) (string, error)
}

type UserService struct {
	Collection *mongo.Collection
	MiddleWare middleware.MiddleWare
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{Collection: collection, MiddleWare: middleware.NewMiddleWare()}
}

func (us *UserService) RegisterUser(user models.User) (models.User, error) {
	// check if the user exist or not

	var existingUser, newUser models.User

	err := us.Collection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&existingUser)

	if err == nil {
		return models.User{}, errors.New("user with these email exists")
	}

	result, err := us.Collection.Find(context.TODO(), bson.D{}, options.Find().SetLimit(1))

	if err != nil {
		return models.User{}, err
	}
	user.ID = primitive.NewObjectID()
	user.Role = "user"
	hasData := result.Next(context.TODO())
	if !hasData {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	newUser = user
	// hashing the password
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashed_password)

	_, errs := us.Collection.InsertOne(context.TODO(), user)

	if errs != nil {
		return models.User{}, nil
	}

	return newUser, nil

}

func (us *UserService) LoginUser(user models.User) (string, error) {
	// hash the password
	var newUser models.User
	err := us.Collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&newUser)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(user.Password))

	if err != nil {
		return "", errors.New("no such user")
	}

	fmt.Println("-----------------", newUser, "---------------------------")
	token, err := us.MiddleWare.GenerateToken(newUser)

	if err != nil {
		return "", err
	}

	return token, nil
}
