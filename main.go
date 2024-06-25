package main

import (
	"net/http"

	"user-app/errors"
	"user-app/model"
	"user-app/repository"
	"user-app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)

	router := gin.Default()

	router.POST("/user", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.CreateUser(user.Name, user.Email, user.Pwd); err != nil {
			if _, ok := err.(*errors.UserExistedError); ok {
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, user)
	})

	router.GET("/user", func(c *gin.Context) {
		type RequestData struct {
			Email string `json:"email"`
			Pwd   string `json:"pwd"`
		}

		var requestData RequestData
		if err := c.ShouldBindBodyWithJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := service.GetUserByEmail(requestData.Email, requestData.Pwd)
		if err != nil {
			if _, ok := err.(*errors.UserNotFoundError); ok {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else if _, ok := err.(*errors.PwdIncorrectError); ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, user)
	})

	router.Run(":8080")
}
