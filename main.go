package main

import (
	"log"

	"github.com/Ateto/User-Login-Service/api"
	"github.com/Ateto/User-Login-Service/db"
	"github.com/Ateto/User-Login-Service/repository"
	"github.com/Ateto/User-Login-Service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.NewDB("./config.json", "./db/init.sql")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	controller := api.NewUserController(service)

	router := SetUpRouter(controller)

	router.Run(":8080")
}

func SetUpRouter(ctrl *api.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("/user", ctrl.GetUser)
	router.POST("user", ctrl.SaveUser)

	return router
}
