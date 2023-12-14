package entities

import (
	"mime/multipart"
)

type UploadUseCases interface {
	Upload(file multipart.File, header *multipart.FileHeader) *FileResponse
}

type FileResponse struct {
	Location  string `json:"Location"`
	VersionID string `json:"VersionID"`
	UploadID  string `json:"UploadID"`
	ETag      string `json:"ETag"`
	Err       error  `json:"Error"`
}
