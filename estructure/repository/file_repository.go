package repository

import (
	"github.com/BautistaBianculli/metadata_archivos/src/application/config"
	"github.com/BautistaBianculli/metadata_archivos/src/core/providers"
	"github.com/BautistaBianculli/metadata_archivos/src/estructure/custom_errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
)

type fileRepository struct {
	session *session.Session
	bucket  string
}

func NewFileRepository(session *session.Session, config config.Config) providers.FileRepository {
	return &fileRepository{
		session: session,
		bucket:  config.AWSBucket,
	}
}

func (f *fileRepository) UploadFile(file multipart.File, header *multipart.FileHeader) (*s3manager.UploadOutput, error) {

	uploader := s3manager.NewUploader(f.session)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(f.bucket),
		Key:    aws.String(header.Filename),
		Body:   file,
	})

	if err != nil {
		return nil, custom_errors.NewInternalServerError(custom_errors.CodeErrorInternalServer, custom_errors.MessageErrorInternalServer, err.Error())
	}

	return upload, nil
}
