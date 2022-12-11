package models

import "time"

type File struct {
	External *ExternalFile `json:"external,omitempty"`
	File     *UploadedFile `json:"file,omitempty"`
}

type UploadedFile struct {
	Url        string     `json:"url"`
	ExpiryTime *time.Time `json:"expiry_time"`
}

type ExternalFile struct {
	Url string `json:"url"`
}
