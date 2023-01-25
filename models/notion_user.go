package models

import (
	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type User struct {
	Object    string          `json:"object"`
	ID        uuid.UUID       `json:"id"`
	Type      *enums.UserType `json:"type,omitempty"`
	AvatarUrl string          `json:"avatar_url,omitempty"`
	Person    *Person         `json:"person,omitempty"`
	Bot       *Bot            `json:"bot,omitempty"`
}

type Person struct {
	Email string `json:"email"`
}

type Bot struct {
	Owner *BotOwner `json:"owner"`
}

type BotOwner struct {
	Type      enums.BotOwnerType `json:"type"`
	Workspace bool               `json:"workspace"`
	User      User               `json:"user"`
}
