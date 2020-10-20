package controllers
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ModelController struct{}

func (h ModelController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
