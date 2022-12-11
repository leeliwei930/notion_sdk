package models

import "time"

type NotionFile struct {
	External *NotionExternalFile `json:"external,omitempty"`
	File     *NotionUploadedFile `json:"file,omitempty"`
}

type NotionUploadedFile struct {
	Url        string     `json:"url"`
	ExpiryTime *time.Time `json:"expiry_time"`
}

type NotionExternalFile struct {
	Url string `json:"url"`
}
