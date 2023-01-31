package models

import (
	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type User struct {
	Object    string          `json:"object"`
	AvatarUrl string          `json:"avatar_url,omitempty"`
	Type      *enums.UserType `json:"type,omitempty"`
	Person    *Person         `json:"person,omitempty"`
	Bot       *Bot            `json:"bot,omitempty"`
	ID        uuid.UUID       `json:"id"`
}

type Person struct {
	Email string `json:"email"`
}

type Bot struct {
	Owner *BotOwner `json:"owner"`
}

type BotOwner struct {
	User      User               `json:"user"`
	Type      enums.BotOwnerType `json:"type"`
	Workspace bool               `json:"workspace"`
}
