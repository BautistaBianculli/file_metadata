package usecases

import (
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
	"github.com/aws/aws-sdk-go/aws"
	"mime/multipart"
)

func (fileUseCase *fileUseCases) Upload(file multipart.File, header *multipart.FileHeader) *entities.UploadFileResponse {

	uploadResponse, err := fileUseCase.fileRepository.UploadFile(file, header)

	if err != nil {
		return &entities.UploadFileResponse{Err: err}
	}

	return &entities.UploadFileResponse{
		Location:  uploadResponse.Location,
		VersionID: aws.StringValue(uploadResponse.VersionID),
		UploadID:  uploadResponse.UploadID,
		ETag:      aws.StringValue(uploadResponse.ETag),
		Err:       nil,
	}
}
