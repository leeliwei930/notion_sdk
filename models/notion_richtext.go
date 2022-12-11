package models

type NotionRichText struct {
	PlainText   string             `json:"plain_text,omitempty"`
	Href        *string            `json:"href,omitempty"`
	Annotations *NotionAnnotations `json:"annotations,omitempty"`
	Type        NotionRichTextType `json:"type,omitempty"`
	Mention     *NotionMention     `json:"mention,omitempty"`
	Text        *NotionText        `json:"text,omitempty"`
}
