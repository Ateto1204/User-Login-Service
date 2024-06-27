package api

import (
	"net/http"

	"github.com/Ateto/User-Login-Service/errors"
	"github.com/Ateto/User-Login-Service/model"
	"github.com/Ateto/User-Login-Service/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	type RequestData struct {
		Email string `json:"email"`
		Pwd   string `json:"pwd"`
	}

	var requestData RequestData
	if err := c.ShouldBindBodyWithJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.service.GetUserByEmail(requestData.Email, requestData.Pwd)
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
}

func (ctrl *UserController) SaveUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.CreateUser(user.Name, user.Email, user.Pwd); err != nil {
		if _, ok := err.(*errors.UserExistedError); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}
