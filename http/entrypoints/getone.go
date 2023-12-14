package entrypoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Files) GetOneHandler(c *gin.Context) {
	fileName := c.Param("filename")

	file := handler.FileUseCases.GetOne(&fileName)

	if file.Err != nil {
		c.JSON(http.StatusInternalServerError, file.Err)
		return
	}

	c.JSON(http.StatusOK, file)
}
