package database

import (
	"time"

	"github.com/leeliwei930/notion_sdk/enums"
	"github.com/leeliwei930/notion_sdk/models"
)

type PropertyValue struct {
	ID             string                  `json:"id,omitempty"`
	Type           enums.PropertyType      `json:"type,omitempty"`
	Title          []models.RichText       `json:"title,omitempty"`
	RichText       []models.RichText       `json:"rich_text,omitempty"`
	Number         *float64                `json:"number,omitempty"`
	Select         *SelectPropertyValue    `json:"select,omitempty"`
	MultiSelect    []SelectPropertyValue   `json:"multi_select,omitempty"`
	Date           *models.DateProperty    `json:"date,omitempty"`
	Status         *SelectPropertyValue    `json:"status,omitempty"`
	Formula        *models.FormulaProperty `json:"formula,omitempty"`
	Relation       []RelationPropertyValue `json:"relation,omitempty"`
	RollUp         *RollUpPropertyValue    `json:"rollup,omitempty"`
	People         []models.User           `json:"people,omitempty"`
	Files          []models.File           `json:"files,omitempty"`
	Checkbox       bool                    `json:"checkbox,omitempty"`
	Url            string                  `json:"url,omitempty"`
	Email          string                  `json:"email,omitempty"`
	PhoneNumber    string                  `json:"phone_number,omitempty"`
	CreatedTime    *time.Time              `json:"created_time,omitempty"`
	CreatedBy      *models.User            `json:"created_by,omitempty"`
	LastEditedTime *time.Time              `json:"last_edited_time,omitempty"`
	LastEditedBy   *models.User            `json:"last_edited_by,omitempty"`
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
	Type    enums.RollUpPropertyValueType `json:"type,omitempty"`
	String  string                        `json:"string,omitempty"`
	Number  float64                       `json:"number,omitempty"`
	Date    *models.DateProperty          `json:"date,omitempty"`
	Results []RollUpPropertyValueResults  `json:"results,omitempty"`
}
type RollUpPropertyValueResults struct {
	String string               `json:"string,omitempty"`
	Number float64              `json:"number,omitempty"`
	Date   *models.DateProperty `json:"date,omitempty"`
}
