package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users map[string]User = make(map[string]User)
var jwtKey = []byte("chera_is_cool")

func main() {
	router := gin.Default()

	router.GET("/", home)

	router.POST("/register", register)
	router.POST("/login", login)
	router.POST("/secured", AuthMiddleWare(), secured)
	router.Run()
}

func secured(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "This is a secured route",
	})
}
func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func register(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
		})
	}

	user.Password = string(hashedPassword)
	users[user.Email] = user
	//TODO: Save user to database

	c.IndentedJSON(http.StatusOK, gin.H{
		"email":    user.Email,
		"password": user.Password,
	})
}

func login(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	storedUser, ok := users[user.Email]

	if !ok || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"email":   storedUser.Email,
	})

	jwtTokent, err := token.SignedString(jwtKey)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   jwtTokent,
	})
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.GetHeader("Authorization")

		if authHandler == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "No token provided",
			})
			c.Abort()
			return
		}

		authParts := strings.Split(authHandler, " ")

		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
