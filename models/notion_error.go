package models

type NotionError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
