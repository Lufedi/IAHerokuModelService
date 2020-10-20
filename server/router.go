package server

import (
	"ModelsService/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthController := new(controllers.ModelController)
	router.GET("/health", healthController.Status)

	return router

}
