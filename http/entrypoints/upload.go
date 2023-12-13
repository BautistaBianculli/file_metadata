package entrypoints

import (
	"fmt"
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
	"github.com/gin-gonic/gin"
)

type Upload struct {
	UploadUseCase entities.UploadUseCases
}

func (handler *Upload) Handle(c *gin.Context) {

	fmt.Println("Entre")
}
