package web

import (
	"crud/web/controller"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUser)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users", controller.UpdateUser)
	router.DELETE("/users", controller.DeleteUser)

	router.Run(":8080")
}
