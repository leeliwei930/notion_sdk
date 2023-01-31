package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type SelectProperty struct {
	Name  string      `json:"name,omitempty"`
	ID    *uuid.UUID  `json:"id,omitempty"`
	Color enums.Color `json:"color"`
}

type MultiSelectProperty SelectProperty

type DateProperty struct {
	TimeZone string    `json:"time_zone,omitempty"`
	Start    *DateTime `json:"start"`
	End      *DateTime `json:"end,omitempty"`
}

type FormulaProperty struct {
	String   string                         `json:"string,omitempty"`
	Number   *float64                       `json:"number,omitempty"`
	Date     *DateProperty                  `json:"date,omitempty"`
	Relation *RelationProperty              `json:"relation,omitempty"`
	Type     enums.FormulaPropertyValueType `json:"type"`
	Boolean  bool                           `json:"boolean,omitempty"`
}

type RelationProperty struct {
	ID *uuid.UUID `json:"id,omitempty"`
}

type RollUpProperty struct {
	Number *float64      `json:"number,omitempty"`
	Date   *DateProperty `json:"date,omitempty"`
}

type RollUpArrayProperty struct {
	People         []User             `json:"people,omitempty"`
	MultiSelect    []SelectProperty   `json:"multi_select,omitempty"`
	RichText       []RichText         `json:"rich_text,omitempty"`
	Files          []File             `json:"files,omitempty"`
	Title          []RichText         `json:"title,omitempty"`
	CreatedBy      *User              `json:"created_by,omitempty"`
	PhoneNumber    *string            `json:"phone_number,omitempty"`
	Relation       *RelationProperty  `json:"relation,omitempty"`
	LastEditedBy   *User              `json:"last_edited_by,omitempty"`
	Select         *SelectProperty    `json:"select,omitempty"`
	Formula        *FormulaProperty   `json:"formula,omitempty"`
	Number         *float64           `json:"number,omitempty"`
	LastEditedTime *time.Time         `json:"last_edited_time,omitempty"`
	Url            *string            `json:"url,omitempty"`
	Email          *string            `json:"email,omitempty"`
	Date           *DateProperty      `json:"date,omitempty"`
	CreatedTime    *time.Time         `json:"created_time,omitempty"`
	RollUp         *RollUpProperty    `json:"rollup,omitempty"`
	Checkbox       bool               `json:"checkbox,omitempty"`
	Type           enums.PropertyType `json:"type"`
}
