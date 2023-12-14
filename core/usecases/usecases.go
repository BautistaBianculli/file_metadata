package usecases

import (
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
	"github.com/BautistaBianculli/metadata_archivos/src/core/providers"
)

type fileUseCases struct {
	fileRepository providers.FileRepository
}

func NewFileCases(fileRepository providers.FileRepository) entities.FileUseCases {
	return &fileUseCases{
		fileRepository: fileRepository,
	}
}
