package database

import "github.com/leeliwei930/notion_sdk/enums"

type RollUpProperty struct {
	RelationPropertyName string                   `json:"relation_property_name,omitempty"`
	RelationPropertyID   string                   `json:"relation_property_id,omitempty"`
	RollUpPropertyName   string                   `json:"rollup_property_name,omitempty"`
	RollUpPropertyID     string                   `json:"rollup_property_id,omitempty"`
	Function             enums.RollUpFunctionType `json:"function,omitempty"`
}
