package database

import (
	"github.com/leeliwei930/notion_sdk/enums"
	"github.com/leeliwei930/notion_sdk/models"
)

type DatabaseProperty struct {
	ID          string                     `json:"id,omitempty"`
	Type        enums.DatabasePropertyType `json:"type,omitempty"`
	Name        string                     `json:"name,omitempty"`
	RichText    *[]models.RichText         `json:"rich_text,omitempty"`
	Number      *NumberProperty            `json:"number,omitempty"`
	Select      *SelectProperty            `json:"select,omitempty"`
	Status      *StatusProperty            `json:"status,omitempty"`
	MultiSelect *MultiSelectProperty       `json:"multi_select,omitempty"`
	Formula     *FormulaProperty           `json:"formula,omitempty"`
	Relation    *RelationProperty          `json:"relation,omitempty"`
	RollUp      *RollUpProperty            `json:"rollup,omitempty"`
}
