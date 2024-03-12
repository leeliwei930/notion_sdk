package models

import (
	"github.com/google/uuid"
)

type PageProperties map[string]*PropertyValue

func (p PageProperties) Title(propertyName string) string {
	if len(p[propertyName].Title) != 0 {
		return p[propertyName].Title[0].PlainText
	}
	return ""
}

func (p PageProperties) RichTextAt(propertyName string, index int) string {
	if len(p[propertyName].RichText) != 0 {
		return p[propertyName].RichText[index].PlainText
	}
	return ""
}

func (p PageProperties) SelectName(propertyName string) string {
	return p[propertyName].Select.Name
}

func (p PageProperties) Url(propertyName string) string {
	return p[propertyName].Url
}

func (p PageProperties) RelationFirstItem(propertyName string) *RelationPropertyValue {
	if len(p[propertyName].Relation) == 0 {
		return nil
	}

	return &p[propertyName].Relation[0]
}

func (p PageProperties) RelationFirstItemID(propertyName string) (uuid.UUID, error) {
	if p.RelationFirstItem(propertyName) == nil {
		return uuid.Nil, nil
	}

	id, err := uuid.Parse(p.RelationFirstItem(propertyName).ID)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
