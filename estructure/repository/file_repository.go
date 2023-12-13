package repository

import (
	"github.com/BautistaBianculli/metadata_archivos/src/application/config"
	"github.com/BautistaBianculli/metadata_archivos/src/core/providers"
	"github.com/aws/aws-sdk-go/aws/session"
)

type fileRepository struct {
	session session.Session
	bucket  string
}

type fileResponse struct {
}

func NewFileRepository(session session.Session, config config.Config) providers.FileRepository {
	return &fileRepository{
		session: session,
		bucket:  config.AWSBucket,
	}
}

func (f *fileRepository) UploadFile() {

}
