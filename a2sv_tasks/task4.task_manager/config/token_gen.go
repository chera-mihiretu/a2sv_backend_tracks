package config

import (
	"github/chera/task_manager/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user models.User) (string, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"useremail": user.Email,
		"userid":    user.ID,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil

}
