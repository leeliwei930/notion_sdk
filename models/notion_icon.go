package models

type Icon struct {
	Type     string        `json:"type,omitempty"`
	Emoji    string        `json:"emoji,omitempty"`
	External *ExternalFile `json:"external,omitempty"`
	File     *UploadedFile `json:"file,omitempty"`
}
