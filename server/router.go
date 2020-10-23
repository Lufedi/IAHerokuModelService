package server

import (
	"ModelsService/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	modelController := new(controllers.ModelController)
	router.GET("/health", modelController.Status)
	router.POST("/upload", modelController.UploadModel)
	router.POST("/model", modelController.SaveModel)
	return router
}
