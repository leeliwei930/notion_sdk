package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type Page struct {
	Object         string                    `json:"object"`
	Url            string                    `json:"url"`
	ID             *uuid.UUID                `json:"id,omitempty"`
	CreatedTime    *time.Time                `json:"created_time"`
	LastEditedTime *time.Time                `json:"last_edited_time"`
	Icon           *Icon                     `json:"icon"`
	Cover          *Icon                     `json:"cover"`
	Properties     map[string]*PropertyValue `json:"properties,omitempty"`
	Parent         *PageParent               `json:"parent"`
	Archived       bool                      `json:"archived"`
}

type PageParent struct {
	DatabaseID *uuid.UUID           `json:"database_id,omitempty"`
	PageID     *uuid.UUID           `json:"page_id,omitempty"`
	Type       enums.PageParentType `json:"type,omitempty"`
	Workspace  bool                 `json:"workspace,omitempty"`
}

type Property struct {
	People         []User             `json:"people,omitempty"`
	Title          []RichText         `json:"title,omitempty"`
	RichText       []RichText         `json:"rich_text,omitempty"`
	MultiSelect    []SelectProperty   `json:"multi_select,omitempty"`
	Files          []File             `json:"files,omitempty"`
	PhoneNumber    string             `json:"phone_number,omitempty"`
	Email          string             `json:"email,omitempty"`
	ID             string             `json:"id,omitempty"`
	Url            string             `json:"url,omitempty"`
	Select         *SelectProperty    `json:"select,omitempty"`
	RollUp         *RollUpProperty    `json:"rollup,omitempty"`
	Relation       *RelationProperty  `json:"relation,omitempty"`
	Formula        *FormulaProperty   `json:"formula,omitempty"`
	Date           *DateProperty      `json:"date,omitempty"`
	Number         *float64           `json:"number,omitempty"`
	CreatedTime    *time.Time         `json:"created_time,omitempty"`
	CreatedBy      *User              `json:"created_by,omitempty"`
	LastEditedTime *time.Time         `json:"last_edited_time,omitempty"`
	LastEditedBy   *User              `json:"last_edited_by,omitempty"`
	Checkbox       bool               `json:"checkbox,omitempty"`
	Type           enums.PropertyType `json:"type,omitempty"`
}
