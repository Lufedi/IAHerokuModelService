package controllers

import (
	"ModelsService/client"
	model2 "ModelsService/model"
	"ModelsService/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


var s3Client = new(client.S3Client)
var modelRepository = new(repositories.ModelRepository)

type ModelController struct{

}

func (h ModelController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

/*
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/appleboy/test.zip" \
  -H "Content-Type: multipart/form-data"
 */
func(h ModelController) UploadModel(c *gin.Context)  {
	fileHeader, error := c.FormFile("file")
	if error != nil {
		fmt.Println("Error receiving the file", error)
	}
	file,_ := fileHeader.Open()
	s3Client.UploadFile(file, fileHeader.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fileHeader.Filename))
}

func (h ModelController) SaveModel(c *gin.Context)  {
	var model model2.Model
	err := c.BindJSON(&model)
	if err != nil {
		fmt.Errorf("Error binding model", err)
		c.String(http.StatusInternalServerError, "Error creating model")
	}

	err = modelRepository.Insert(&model)

	if err != nil {
		fmt.Println("Error creating the object")
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error creating the object")
	}

	c.String(http.StatusCreated, "")
}