package enums

type PropertyType string

const (
	RichTextPropertyType       PropertyType = "rich_text"
	NumberPropertyType         PropertyType = "number"
	SelectPropertyType         PropertyType = "select"
	MultiSelectPropertyType    PropertyType = "multi_select"
	DatePropertyType           PropertyType = "date"
	FormulaPropertyType        PropertyType = "formula"
	RelationPropertyType       PropertyType = "relation"
	RollUpPropertyType         PropertyType = "rollup"
	TitlePropertyType          PropertyType = "title"
	PeoplePropertyType         PropertyType = "people"
	FilesPropertyType          PropertyType = "files"
	CheckboxPropertyType       PropertyType = "checkbox"
	UrlPropertyType            PropertyType = "url"
	EmailPropertyType          PropertyType = "email"
	PhoneNumberPropertyType    PropertyType = "phone_number"
	CreatedTimePropertyType    PropertyType = "created_time"
	CreatedByPropertyType      PropertyType = "created_by"
	LastEditedTimePropertyType PropertyType = "last_edited_time"
	LastEditedByPropertyType   PropertyType = "last_edited_by"
)
