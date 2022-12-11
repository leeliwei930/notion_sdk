package models

import "time"

type Database struct {
	Object         string                       `json:"object,omitempty"`
	ID             string                       `json:"id,omitempty"`
	CreatedTime    *time.Time                   `json:"created_time,omitempty"`
	CreatedBy      *User                        `json:"created_by,omitempty"`
	LastEditedTime *time.Time                   `json:"last_edited_time,omitempty"`
	LastEditedBy   *User                        `json:"last_edited_by,omitempty"`
	Title          []RichText                   `json:"title,omitempty"`
	Description    []RichText                   `json:"description,omitempty"`
	Icon           *Icon                        `json:"icon,omitempty"`
	Cover          *Icon                        `json:"cover,omitempty"`
	Properties     map[string]*DatabaseProperty `json:"properties,omitempty"`
	Parent         *PageParent                  `json:"parent,omitempty"`
	Url            string                       `json:"url,omitempty"`
	Archived       bool                         `json:"archived,omitempty"`
	IsInline       bool                         `json:"is_inline,omitempty"`
}
