package models

import (
	"time"

	"github.com/leeliwei930/notion_sdk/enums"
)

type PropertyValue struct {
	MultiSelect    []SelectPropertyValue   `json:"multi_select,omitempty"`
	Files          []File                  `json:"files,omitempty"`
	Title          []RichText              `json:"title,omitempty"`
	RichText       []RichText              `json:"rich_text,omitempty"`
	People         []User                  `json:"people,omitempty"`
	Relation       []RelationPropertyValue `json:"relation,omitempty"`
	Url            string                  `json:"url,omitempty"`
	ID             string                  `json:"id,omitempty"`
	PhoneNumber    string                  `json:"phone_number,omitempty"`
	Email          string                  `json:"email,omitempty"`
	RollUp         *RollUpPropertyValue    `json:"rollup,omitempty"`
	Date           *DateProperty           `json:"date,omitempty"`
	Number         *float64                `json:"number,omitempty"`
	Select         *SelectPropertyValue    `json:"select,omitempty"`
	Formula        *FormulaProperty        `json:"formula,omitempty"`
	Status         *SelectPropertyValue    `json:"status,omitempty"`
	CreatedTime    *time.Time              `json:"created_time,omitempty"`
	CreatedBy      *User                   `json:"created_by,omitempty"`
	LastEditedTime *time.Time              `json:"last_edited_time,omitempty"`
	LastEditedBy   *User                   `json:"last_edited_by,omitempty"`
	Type           enums.PropertyType      `json:"type,omitempty"`
	Checkbox       bool                    `json:"checkbox,omitempty"`
	HasMore        bool                    `json:"has_more,omitempty"`
}

type SelectPropertyValue struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}
type RelationPropertyValue struct {
	ID string `json:"id,omitempty"`
}

type RollUpPropertyValue struct {
	Results []RollUpPropertyValueResults  `json:"results,omitempty"`
	String  string                        `json:"string,omitempty"`
	Number  float64                       `json:"number,omitempty"`
	Date    *DateProperty                 `json:"date,omitempty"`
	Type    enums.RollUpPropertyValueType `json:"type,omitempty"`
}
type RollUpPropertyValueResults struct {
	String string        `json:"string,omitempty"`
	Number float64       `json:"number,omitempty"`
	Date   *DateProperty `json:"date,omitempty"`
}
