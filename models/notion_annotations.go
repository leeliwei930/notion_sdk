package models

type NotionAnnotations struct {
	Bold          bool        `json:"bold"`
	Italic        bool        `json:"italic"`
	Strikethrough bool        `json:"strikethrough"`
	Underline     bool        `json:"underline"`
	Code          bool        `json:"code"`
	Color         NotionColor `json:"color"`
}
