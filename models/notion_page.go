package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type Page struct {
	Object         string               `json:"object"`
	ID             *uuid.UUID           `json:"id,omitempty"`
	CreatedTime    *time.Time           `json:"created_time"`
	LastEditedTime *time.Time           `json:"last_edited_time"`
	Archived       bool                 `json:"archived"`
	Icon           *Icon                `json:"icon"`
	Cover          *ExternalFile        `json:"cover"`
	Properties     map[string]*Property `json:"properties,omitempty"`
	Parent         *PageParent          `json:"parent"`
	Url            string               `json:"url"`
}

type PageParent struct {
	Type       enums.PageParentType `json:"type,omitempty"`
	DatabaseID *uuid.UUID           `json:"database_id,omitempty"`
	PageID     *uuid.UUID           `json:"page_id,omitempty"`
	Workspace  bool                 `json:"workspace,omitempty"`
}

type Property struct {
	ID             string             `json:"id,omitempty"`
	Type           enums.PropertyType `json:"type,omitempty"`
	Title          []RichText         `json:"title,omitempty"`
	RichText       []RichText         `json:"rich_text,omitempty"`
	Number         *float64           `json:"number,omitempty"`
	Select         *SelectProperty    `json:"select,omitempty"`
	MultiSelect    []SelectProperty   `json:"multi_select,omitempty"`
	Date           *DateProperty      `json:"date,omitempty"`
	Formula        *FormulaProperty   `json:"formula,omitempty"`
	Relation       *RelationProperty  `json:"relation,omitempty"`
	RollUp         *RollUpProperty    `json:"rollup,omitempty"`
	People         []User             `json:"people,omitempty"`
	Files          []File             `json:"files,omitempty"`
	Checkbox       bool               `json:"checkbox,omitempty"`
	Url            string             `json:"url,omitempty"`
	Email          string             `json:"email,omitempty"`
	PhoneNumber    string             `json:"phone_number,omitempty"`
	CreatedTime    *time.Time         `json:"created_time,omitempty"`
	CreatedBy      *User              `json:"created_by,omitempty"`
	LastEditedTime *time.Time         `json:"last_edited_time,omitempty"`
	LastEditedBy   *User              `json:"last_edited_by,omitempty"`
}
