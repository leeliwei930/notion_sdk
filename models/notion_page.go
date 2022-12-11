package models

import (
	"time"

	"github.com/google/uuid"
)

type NotionPage struct {
	Object         string                      `json:"object"`
	ID             *uuid.UUID                  `json:"id,omitempty"`
	CreatedTime    *time.Time                  `json:"created_time"`
	LastEditedTime *time.Time                  `json:"last_edited_time"`
	Archived       bool                        `json:"archived"`
	Icon           *NotionIcon                 `json:"icon"`
	Cover          *NotionExternalFile         `json:"cover"`
	Properties     *map[string]*NotionProperty `json:"properties,omitempty"`
	Parent         *NotionPageParent           `json:"parent"`
	Url            string                      `json:"url"`
}

type NotionPageParent struct {
	Type       NotionPageParentType `json:"type,omitempty"`
	DatabaseID *uuid.UUID           `json:"database_id,omitempty"`
	PageID     *uuid.UUID           `json:"page_id,omitempty"`
	Workspace  bool                 `json:"workspace,omitempty"`
}

type NotionProperty struct {
	ID             string                  `json:"id,omitempty"`
	Type           NotionPropertyType      `json:"type,omitempty"`
	Title          *[]NotionRichText       `json:"title,omitempty"`
	RichText       *[]NotionRichText       `json:"rich_text,omitempty"`
	Number         *float64                `json:"number,omitempty"`
	Select         *NotionSelectProperty   `json:"select,omitempty"`
	MultiSelect    *[]NotionSelectProperty `json:"multi_select,omitempty"`
	Date           *NotionDateProperty     `json:"date,omitempty"`
	Formula        *NotionFormulaProperty  `json:"formula,omitempty"`
	Relation       *NotionRelationProperty `json:"relation,omitempty"`
	RollUp         *NotionRollUpProperty   `json:"rollup,omitempty"`
	People         *[]NotionUser           `json:"people,omitempty"`
	Files          *[]NotionFile           `json:"files,omitempty"`
	Checkbox       bool                    `json:"checkbox,omitempty"`
	Url            *string                 `json:"url,omitempty"`
	Email          *string                 `json:"email,omitempty"`
	PhoneNumber    *string                 `json:"phone_number,omitempty"`
	CreatedTime    *time.Time              `json:"created_time,omitempty"`
	CreatedBy      *NotionUser             `json:"created_by,omitempty"`
	LastEditedTime *time.Time              `json:"last_edited_time,omitempty"`
	LastEditedBy   *NotionUser             `json:"last_edited_by,omitempty"`
}
