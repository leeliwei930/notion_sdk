package models

import "time"

type NotionDatabase struct {
	Object         string                             `json:"object,omitempty"`
	ID             string                             `json:"id,omitempty"`
	CreatedTime    *time.Time                         `json:"created_time,omitempty"`
	CreatedBy      *NotionUser                        `json:"created_by,omitempty"`
	LastEditedTime *time.Time                         `json:"last_edited_time,omitempty`
	LastEditedBy   *NotionUser                        `json:"last_edited_by,omitempty"`
	Title          []NotionRichText                   `json:"title,omitempty"`
	Description    []NotionRichText                   `json:"description,omitempty"`
	Icon           *NotionIcon                        `json:"icon,omitempty"`
	Cover          *NotionIcon                        `json:"cover,omitempty"`
	Properties     map[string]*NotionDatabaseProperty `json:"properties,omitempty"`
	Parent         *NotionPageParent                  `json:"parent,omitempty"`
	Url            string                             `json:"url,omitempty"`
	Archived       bool                               `json:"archived,omitempty"`
	IsInline       bool                               `json:"is_inline,omitempty"`
}
