package filter

import "github.com/leeliwei930/notion_sdk/database"

type Cursor struct {
	Results    []database.PageRef `json:"results,omitempty"`
	Object     string             `json:"object,omitempty"`
	NextCursor string             `json:"next_cursor,omitempty"`
	Type       string             `json:"type,omitempty"`
	Page       any                `json:"page,omitempty"`
	HasMore    bool               `json:"has_more,omitempty"`
}
