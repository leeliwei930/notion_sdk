package models

import "github.com/leeliwei930/notion_sdk/enums"

type RichText struct {
	PlainText   string             `json:"plain_text,omitempty"`
	Href        string             `json:"href,omitempty"`
	Annotations *Annotations       `json:"annotations,omitempty"`
	Mention     *Mention           `json:"mention,omitempty"`
	Text        *Text              `json:"text,omitempty"`
	Type        enums.RichTextType `json:"type,omitempty"`
}
