package models

type NotionText struct {
	Content string      `json:"content"`
	Link    *NotionLink `json:"link,omitempty"`
}
