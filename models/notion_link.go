package models

import "github.com/leeliwei930/notion_sdk/enums"

type Link struct {
	Url string `json:"url"`
}

type Mention struct {
	Type enums.MentionType `json:"type"`
	User *User             `json:"user,omitempty"`
}
