package usecases

import (
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
	"github.com/BautistaBianculli/metadata_archivos/src/core/providers"
)

type uploadUseCases struct {
	fileRepository providers.FileRepository
}

func NewUploadCases(fileRepository providers.FileRepository) entities.UploadUseCases {
	return &uploadUseCases{
		fileRepository: fileRepository,
	}
}

func (u *uploadUseCases) Upload() {

}
