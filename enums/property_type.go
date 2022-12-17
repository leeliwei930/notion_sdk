package enums

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

var PropertyTypeMap = map[PropertyType]string{
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

func (p PropertyType) String() string {
	return PropertyTypeMap[p]
}
