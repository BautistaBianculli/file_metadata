package entrypoints

import (
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
	"github.com/gin-gonic/gin"
)

type FileHandler interface {
	UploadHandle(c *gin.Context)
	GetAllHandler(c *gin.Context)
	GetOneHandler(c *gin.Context)
	DeleteOneHandler(c *gin.Context)
}

type Files struct {
	FileUseCases entities.FileUseCases
}
