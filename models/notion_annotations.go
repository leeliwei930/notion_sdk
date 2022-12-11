package models

import "github.com/leeliwei930/notion_sdk/enums"

type Annotations struct {
	Bold          bool        `json:"bold"`
	Italic        bool        `json:"italic"`
	Strikethrough bool        `json:"strikethrough"`
	Underline     bool        `json:"underline"`
	Code          bool        `json:"code"`
	Color         enums.Color `json:"color"`
}
