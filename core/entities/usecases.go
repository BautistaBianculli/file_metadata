package entities

import "mime/multipart"

type FileUseCases interface {
	Upload(file multipart.File, header *multipart.FileHeader) *UploadFileResponse
	GetAll() ([]*GetAllResponse, error)
	GetOne(file *string) *GetOne
}
