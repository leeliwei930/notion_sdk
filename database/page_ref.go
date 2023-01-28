package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/models"
)

type PageRef struct {
	Object         string                           `json:"object"`
	ID             *uuid.UUID                       `json:"id,omitempty"`
	CreatedTime    *time.Time                       `json:"created_time"`
	LastEditedTime *time.Time                       `json:"last_edited_time"`
	Archived       bool                             `json:"archived"`
	Icon           *models.Icon                     `json:"icon"`
	Cover          *models.Icon                     `json:"cover"`
	Properties     map[string]*models.PropertyValue `json:"properties,omitempty"`
	Parent         *models.PageParent               `json:"parent"`
	Url            string                           `json:"url"`
}
