package models

type NotionLink struct {
	Url string `json:"url"`
}

type NotionMention struct {
	Type NotionMentionType `json:"type"`
	User *NotionUser       `json:"user,omitempty"`
}
