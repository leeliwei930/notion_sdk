package models

import "github.com/google/uuid"

type NotionUser struct {
	Object    string          `json:"object"`
	ID        uuid.UUID       `json:"id"`
	Type      *NotionUserType `json:"type,omitempty"`
	AvatarUrl *string         `json:"avatar_url,omitempty"`
	Person    *NotionPerson   `json:"person,omitempty"`
	Bot       *NotionBot      `json:"bot,omitempty"`
}

type NotionPerson struct {
	Email string `json:"email"`
}

type NotionBot struct {
	Owner *NotionBotOwner `json:"owner"`
}

type NotionBotOwner struct {
	Type      NotionBotOwnerType `json:"type"`
	Workspace bool               `json:"workspace"`
	User      NotionUser         `json:"user"`
}
