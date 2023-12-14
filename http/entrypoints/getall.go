package entrypoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Files) GetAllHandler(c *gin.Context) {
	files, err := handler.FileUseCases.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": files,
	})
}
