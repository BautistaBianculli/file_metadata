package repository

import (
	"github.com/BautistaBianculli/metadata_archivos/src/application/config"
	"github.com/BautistaBianculli/metadata_archivos/src/core/providers"
	"github.com/BautistaBianculli/metadata_archivos/src/estructure/custom_errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
	"time"
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

func (f *fileRepository) GetAllFiles() ([]*s3.Object, error) {
	svc := s3.New(f.session)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(f.bucket),
	})
	if err != nil {
		return nil, custom_errors.NewInternalServerError(custom_errors.CodeErrorInternalServer, custom_errors.MessageErrorInternalServer, err.Error())
	}
	return resp.Contents, nil
}

func (f *fileRepository) GetOneFile(fileName *string) (string, error) {

	svc := s3.New(f.session)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(f.bucket),
		Key:    fileName,
	})

	urlDownload, err := req.Presign(10 * time.Minute)

	if err != nil {
		return "", custom_errors.NewInternalServerError(custom_errors.CodeErrorInternalServer, custom_errors.MessageErrorInternalServer, err.Error())
	}
	return urlDownload, nil
}
