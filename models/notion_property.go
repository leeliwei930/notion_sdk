package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type SelectProperty struct {
	ID    *uuid.UUID  `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Color enums.Color `json:"color"`
}

type MultiSelectProperty SelectProperty

type DateProperty struct {
	Start    time.Time  `json:"start"`
	End      *time.Time `json:"end,omitempty"`
	TimeZone string     `json:"time_zone,omitempty"`
}

type FormulaProperty struct {
	Type     enums.FormulaPropertyValueType `json:"type"`
	String   string                         `json:"string,omitempty"`
	Number   *float64                       `json:"number,omitempty"`
	Boolean  bool                           `json:"boolean,omitempty"`
	Date     *DateProperty                  `json:"date,omitempty"`
	Relation *RelationProperty              `json:"relation,omitempty"`
}

type RelationProperty struct {
	ID *uuid.UUID `json:"id"`
}

type RollUpProperty struct {
	Number *float64      `json:"number,omitempty"`
	Date   *DateProperty `json:"date,omitempty"`
}

type RollUpArrayProperty struct {
	Type           enums.PropertyType `json:"type"`
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
	Url            *string            `json:"url,omitempty"`
	Email          *string            `json:"email,omitempty"`
	PhoneNumber    *string            `json:"phone_number,omitempty"`
	CreatedTime    *time.Time         `json:"created_time,omitempty"`
	CreatedBy      *User              `json:"created_by,omitempty"`
	LastEditedTime *time.Time         `json:"last_edited_time,omitempty"`
	LastEditedBy   *User              `json:"last_edited_by,omitempty"`
}
