package providers

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
)

type FileRepository interface {
	UploadFile(file multipart.File, header *multipart.FileHeader) (*s3manager.UploadOutput, error)
	GetAllFiles() ([]*s3.Object, error)
	GetOneFile(file *string) (string, error)
}
