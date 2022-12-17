package enums

type DatabasePropertyType int

const (
	Title DatabasePropertyType = iota
	RichText
	Number
	Select
	MultiSelect
	Date
	People
	Files
	Checkbox
	Url
	Email
	PhoneNumber
	Formula
	Relation
	RollUp
	CreatedTime
	CreatedBy
	LastEditedBy
	LastEditedTime
	Status
)

var DatabasePropertyTypeMap = map[DatabasePropertyType]string{
	Title:          "title",
	RichText:       "rich_text",
	Number:         "number",
	Select:         "select",
	MultiSelect:    "multi_select",
	Date:           "date",
	People:         "people",
	Files:          "files",
	Checkbox:       "checkbox",
	Url:            "url",
	Email:          "email",
	PhoneNumber:    "phone_number",
	Formula:        "formula",
	Relation:       "relation",
	RollUp:         "rollup",
	CreatedTime:    "created_time",
	CreatedBy:      "created_by",
	LastEditedBy:   "last_edited_by",
	LastEditedTime: "last_edited_time",
	Status:         "status",
}
