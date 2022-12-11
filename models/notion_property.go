package models

import (
	"time"

	"github.com/google/uuid"
)

type NotionSelectProperty struct {
	ID    *uuid.UUID  `json:"id,omitempty"`
	Name  *string     `json:"name,omitempty"`
	Color NotionColor `json:"color"`
}

type NotionMultiSelectProperty NotionSelectProperty

type NotionDateProperty struct {
	Start time.Time  `json:"start"`
	End   *time.Time `json:"end,omitempty"`
}

type NotionFormulaProperty struct {
	Type     NotionFormulaPropertyValueType `json:"type"`
	String   *string                        `json:"string,omitempty"`
	Number   *float64                       `json:"number,omitempty"`
	Boolean  bool                           `json:"boolean,omitempty"`
	Date     *NotionDateProperty            `json:"date,omitempty"`
	Relation *NotionRelationProperty        `json:"relation,omitempty"`
}

type NotionRelationProperty struct {
	ID *uuid.UUID `json:"id"`
}

type NotionRollUpProperty struct {
	Number *float64            `json:"number,omitempty"`
	Date   *NotionDateProperty `json:"date,omitempty"`
}

type NotionRollUpArrayProperty struct {
	Type           NotionPropertyType      `json:"type"`
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
