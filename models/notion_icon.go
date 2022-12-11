package models

type NotionIcon struct {
	Type     string              `json:"type,omitempty"`
	External *NotionExternalFile `json:"external,omitempty"`
	File     *NotionUploadedFile `json:"file,omitempty"`
	Emoji    string              `json:"emoji,omitempty"`
}
