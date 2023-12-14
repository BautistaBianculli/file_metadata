package usecases

import (
	"github.com/BautistaBianculli/metadata_archivos/src/core/entities"
	"github.com/BautistaBianculli/metadata_archivos/src/core/providers"
	"github.com/aws/aws-sdk-go/aws"
	"mime/multipart"
)

type uploadUseCases struct {
	fileRepository providers.FileRepository
}

func NewUploadCases(fileRepository providers.FileRepository) entities.UploadUseCases {
	return &uploadUseCases{
		fileRepository: fileRepository,
	}
}

func (u *uploadUseCases) Upload(file multipart.File, header *multipart.FileHeader) *entities.FileResponse {

	uploadResponse, err := u.fileRepository.UploadFile(file, header)
	if err != nil {
		return &entities.FileResponse{Err: err}
	}
	return &entities.FileResponse{
		Location:  uploadResponse.Location,
		VersionID: aws.StringValue(uploadResponse.VersionID),
		UploadID:  uploadResponse.UploadID,
		ETag:      aws.StringValue(uploadResponse.ETag),
		Err:       nil,
	}
}
