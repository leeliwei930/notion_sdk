package database

import (
	"time"

	"github.com/leeliwei930/notion_sdk/models"
)

type Database struct {
	Object         string                       `json:"object,omitempty"`
	ID             string                       `json:"id,omitempty"`
	CreatedTime    *time.Time                   `json:"created_time,omitempty"`
	CreatedBy      *models.User                 `json:"created_by,omitempty"`
	LastEditedTime *time.Time                   `json:"last_edited_time,omitempty"`
	LastEditedBy   *models.User                 `json:"last_edited_by,omitempty"`
	Title          []models.RichText            `json:"title,omitempty"`
	Description    []models.RichText            `json:"description,omitempty"`
	Icon           *models.Icon                 `json:"icon,omitempty"`
	Cover          *models.Icon                 `json:"cover,omitempty"`
	Properties     map[string]*DatabaseProperty `json:"properties,omitempty"`
	Parent         *models.PageParent           `json:"parent,omitempty"`
	Url            string                       `json:"url,omitempty"`
	Archived       bool                         `json:"archived,omitempty"`
	IsInline       bool                         `json:"is_inline,omitempty"`
}
