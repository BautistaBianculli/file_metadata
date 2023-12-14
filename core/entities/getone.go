package entities

import "time"

type GetOne struct {
	Key      *string   `json:"Name"`
	Url      string    `json:"Url"`
	LifeTime time.Time `json:"LifeTime"`
	Err      error     `json:"Error"`
}
