package models

type NotionError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
