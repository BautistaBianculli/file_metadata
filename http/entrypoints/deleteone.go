package entrypoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Files) DeleteOneHandler(c *gin.Context) {

	fileName := c.Param("filename")

	err := handler.FileUseCases.DeleteOne(&fileName)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.Status(http.StatusOK)
}
