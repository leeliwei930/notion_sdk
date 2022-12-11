package models

type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}
