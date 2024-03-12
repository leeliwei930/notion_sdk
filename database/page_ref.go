package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/models"
)

type PageRef struct {
	Object         string                `json:"object"`
	Url            string                `json:"url"`
	ID             *uuid.UUID            `json:"id,omitempty"`
	CreatedTime    *time.Time            `json:"created_time"`
	LastEditedTime *time.Time            `json:"last_edited_time"`
	Icon           *models.Icon          `json:"icon"`
	Cover          *models.Icon          `json:"cover"`
	Properties     models.PageProperties `json:"properties,omitempty"`
	Parent         *models.PageParent    `json:"parent"`
	Archived       bool                  `json:"archived"`
}
