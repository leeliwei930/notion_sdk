package filter

import "github.com/leeliwei930/notion_sdk/models"

type Cursor struct {
	Object     string        `json:"object,omitempty"`
	Results    []models.Page `json:"results,omitempty"`
	NextCursor string        `json:"next_cursor,omitempty"`
	HasMore    bool          `json:"has_more,omitempty"`
	Type       string        `json:"type,omitempty"`
	Page       any           `json:"page,omitempty"`
}
