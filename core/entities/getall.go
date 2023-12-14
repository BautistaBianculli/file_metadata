package entities

import "time"

type GetAllResponse struct {
	Key        *string    `json:"Name"`
	Size       *int64     `json:"Size"`
	LastModify *time.Time `json:"LastModify"`
}
