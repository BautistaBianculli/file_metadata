package usecases

import (
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
	"time"
)

func (fileUseCase *fileUseCases) GetOne(file *string) *entities.GetOne {

	urlDownload, err := fileUseCase.fileRepository.GetOneFile(file)

	if err != nil {
		return &entities.GetOne{Err: err}
	}

	return &entities.GetOne{
		Key:      file,
		Url:      urlDownload,
		LifeTime: time.Now().Add(15 * time.Minute),
		Err:      nil,
	}
}
