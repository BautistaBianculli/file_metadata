package entrypoints

import (
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
	"github.com/BautistaBianculli/metadata_archivos/src/estructure/custom_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Upload struct {
	UploadUseCase entities.UploadUseCases
}

func (handler *Upload) Handle(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")

	if err != nil {
		err = custom_errors.NewValidationError(custom_errors.CodeErrorInvalidBody, custom_errors.MessageErrorParseError, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	defer file.Close()

	fileResponse := handler.UploadUseCase.Upload(file, header)

	if fileResponse.Err != nil {
		c.JSON(http.StatusInternalServerError, fileResponse.Err)
		return
	}

	c.JSON(http.StatusOK, fileResponse)
}
