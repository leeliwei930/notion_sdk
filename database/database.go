package database

import (
	"time"

	"github.com/leeliwei930/notion_sdk/models"
)

type Database struct {
	Title          []models.RichText            `json:"title,omitempty"`
	Description    []models.RichText            `json:"description,omitempty"`
	ID             string                       `json:"id,omitempty"`
	Url            string                       `json:"url,omitempty"`
	Object         string                       `json:"object,omitempty"`
	CreatedBy      *models.User                 `json:"created_by,omitempty"`
	LastEditedBy   *models.User                 `json:"last_edited_by,omitempty"`
	LastEditedTime *time.Time                   `json:"last_edited_time,omitempty"`
	Icon           *models.Icon                 `json:"icon,omitempty"`
	Cover          *models.Icon                 `json:"cover,omitempty"`
	Properties     map[string]*DatabaseProperty `json:"properties,omitempty"`
	Parent         *models.PageParent           `json:"parent,omitempty"`
	CreatedTime    *time.Time                   `json:"created_time,omitempty"`
	Archived       bool                         `json:"archived,omitempty"`
	IsInline       bool                         `json:"is_inline,omitempty"`
}
