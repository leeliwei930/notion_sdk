package database

import (
	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type RelationProperty struct {
	DatabaseID   *uuid.UUID                         `json:"database_id,omitempty"`
	Type         enums.RelationDatabasePropertyType `json:"type,omitempty"`
	DualProperty *DualPropertyRelationConfiguration `json:"dual_property,omitempty"`
}

type DualPropertyRelationConfiguration struct {
	SyncedPropertyName string `json:"synced_property_name,omitempty"`
	SyncedPropertyID   string `json:"synced_property_id,omitempty"`
}
