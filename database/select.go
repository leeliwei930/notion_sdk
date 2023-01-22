package database

import (
	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type SelectProperty struct {
	SelectOptions []*SelectOptions `json:"options,omitempty"`
}

type SelectOptions struct {
	Name  string      `json:"name,omitempty"`
	ID    *uuid.UUID  `json:"id,omitempty"`
	Color enums.Color `json:"color,omitempty"`
}
