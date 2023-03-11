package models

import (
	"fmt"
)

type NotionError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (notionErr *NotionError) Error() string {
	return fmt.Errorf("failed to make request to notion API with error code: %s", notionErr.Code).Error()
}
