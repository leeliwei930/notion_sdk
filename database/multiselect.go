package database

import "github.com/leeliwei930/notion_sdk/enums"

type MultiSelectProperty struct {
	Options []MultiSelectOptionsProperty `json:"options,omitempty"`
}

type MultiSelectOptionsProperty struct {
	Name  string      `json:"name,omitempty"`
	ID    string      `json:"id,omitempty"`
	Color enums.Color `json:"color,omitempty"`
}
