package models

import "github.com/leeliwei930/notion_sdk/enums"

type DatabaseProperty struct {
	ID       string                     `json:"id,omitempty"`
	Type     enums.DatabasePropertyType `json:"type,omitempty"`
	Name     string                     `json:"name,omitempty"`
	RichText []RichText                 `json:"rich_text,omitempty"`
}
