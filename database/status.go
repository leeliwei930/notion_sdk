package database

import (
	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type StatusProperty struct {
	Options *[]StatusOptions `json:"options,omitempty"`
	Group   *[]StatusGroup   `json:"group,omitempty"`
}

type StatusOptions struct {
	Name  string      `json:"name,omitempty"`
	ID    *uuid.UUID  `json:"id,omitempty"`
	Color enums.Color `json:"color,omitempty"`
}

type StatusGroup struct {
	Name      string       `json:"name,omitempty"`
	ID        *uuid.UUID   `json:"id,omitempty"`
	Color     enums.Color  `json:"color,omitempty"`
	OptionIDs *[]uuid.UUID `json:"option_ids,omitempty"`
}
