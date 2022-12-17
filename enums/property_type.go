package enums

import "encoding/json"

type PropertyType uint8

const (
	RichTextPropertyType PropertyType = iota
	NumberPropertyType
	SelectPropertyType
	MultiSelectPropertyType
	DatePropertyType
	FormulaPropertyType
	RelationPropertyType
	RollUpPropertyType
	TitlePropertyType
	PeoplePropertyType
	FilesPropertyType
	CheckboxPropertyType
	UrlPropertyType
	EmailPropertyType
	PhoneNumberPropertyType
	CreatedTimePropertyType
	CreatedByPropertyType
	LastEditedTimePropertyType
	LastEditedByPropertyType
)

var propertyTypeMap = map[PropertyType]string{
	RichTextPropertyType:       "rich_text",
	NumberPropertyType:         "number",
	SelectPropertyType:         "select",
	MultiSelectPropertyType:    "multi_select",
	DatePropertyType:           "date",
	FormulaPropertyType:        "formula",
	RelationPropertyType:       "relation",
	RollUpPropertyType:         "rollup",
	TitlePropertyType:          "title",
	PeoplePropertyType:         "people",
	FilesPropertyType:          "files",
	CheckboxPropertyType:       "checkbox",
	UrlPropertyType:            "url",
	EmailPropertyType:          "email",
	PhoneNumberPropertyType:    "phone_number",
	CreatedTimePropertyType:    "created_time",
	CreatedByPropertyType:      "created_by",
	LastEditedTimePropertyType: "last_edited_time",
	LastEditedByPropertyType:   "last_edited_by",
}

var propertyTypeIndexes = map[string]PropertyType{
	"rich_text":        RichTextPropertyType,
	"number":           NumberPropertyType,
	"select":           SelectPropertyType,
	"multi_select":     MultiSelectPropertyType,
	"date":             DatePropertyType,
	"formula":          FormulaPropertyType,
	"relation":         RelationPropertyType,
	"rollup":           RollUpPropertyType,
	"title":            TitlePropertyType,
	"people":           PeoplePropertyType,
	"files":            FilesPropertyType,
	"checkbox":         CheckboxPropertyType,
	"url":              UrlPropertyType,
	"email":            EmailPropertyType,
	"phone_number":     PhoneNumberPropertyType,
	"created_time":     CreatedTimePropertyType,
	"created_by":       CreatedByPropertyType,
	"last_edited_time": LastEditedTimePropertyType,
	"last_edited_by":   LastEditedByPropertyType,
}

func ParsePropertyType(s string) PropertyType {
	return propertyTypeIndexes[s]
}

func (p PropertyType) String() string {
	return propertyTypeMap[p]
}

func (p PropertyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PropertyType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*p = ParsePropertyType(j)
	return nil
}
