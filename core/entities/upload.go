package entities

type UploadFileResponse struct {
	Location  string `json:"Location"`
	VersionID string `json:"VersionID"`
	UploadID  string `json:"UploadID"`
	ETag      string `json:"ETag"`
	Err       error  `json:"Error"`
}
