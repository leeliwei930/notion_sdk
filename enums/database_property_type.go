package enums

import "encoding/json"

type DatabasePropertyType uint8

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

var databasePropertyTypeMap = map[DatabasePropertyType]string{
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

var databasePropertyTypeIndexes = map[string]DatabasePropertyType{
	"title":            Title,
	"rich_text":        RichText,
	"number":           Number,
	"select":           Select,
	"multi_select":     MultiSelect,
	"date":             Date,
	"people":           People,
	"files":            Files,
	"checkbox":         Checkbox,
	"url":              Url,
	"email":            Email,
	"phone_number":     PhoneNumber,
	"formula":          Formula,
	"relation":         Relation,
	"rollup":           RollUp,
	"created_time":     CreatedTime,
	"created_by":       CreatedBy,
	"last_edited_by":   LastEditedBy,
	"last_edited_time": LastEditedTime,
	"status":           Status,
}

func ParseDatabasePropertyType(s string) DatabasePropertyType {
	return databasePropertyTypeIndexes[s]
}
func (d DatabasePropertyType) String() string {
	return databasePropertyTypeMap[d]
}

func (d DatabasePropertyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *DatabasePropertyType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*d = ParseDatabasePropertyType(j)
	return nil
}
