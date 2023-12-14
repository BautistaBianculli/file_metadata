package usecases

import (
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
)

func (fileUseCase *fileUseCases) GetAll() ([]*entities.GetAllResponse, error) {
	content, err := fileUseCase.fileRepository.GetAllFiles()
	if err != nil {
		return nil, err
	}
	var files []*entities.GetAllResponse
	for _, file := range content {

		files = append(files, &entities.GetAllResponse{
			Key:        file.Key,
			Size:       file.Size,
			LastModify: file.LastModified,
		})
	}

	return files, nil
}
