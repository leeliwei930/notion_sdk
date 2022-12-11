package models

type NotionDatabaseProperty struct {
	ID       string                     `json:"id,omitempty"`
	Type     NotionDatabasePropertyType `json:"type,omitempty"`
	Name     string                     `json:"name,omitempty"`
	RichText []NotionRichText           `json:"rich_text,omitempty"`
}

type NotionDatabasePropertyType string

const (
	Title          NotionDatabasePropertyType = "title"
	RichText       NotionDatabasePropertyType = "rich_text"
	Number         NotionDatabasePropertyType = "number"
	Select         NotionDatabasePropertyType = "select"
	MultiSelect    NotionDatabasePropertyType = "multi_select"
	Date           NotionDatabasePropertyType = "date"
	People         NotionDatabasePropertyType = "people"
	Files          NotionDatabasePropertyType = "files"
	Checkbox       NotionDatabasePropertyType = "checkbox"
	Url            NotionDatabasePropertyType = "url"
	Email          NotionDatabasePropertyType = "email"
	PhoneNumber    NotionDatabasePropertyType = "phone_number"
	Formula        NotionDatabasePropertyType = "formula"
	Relation       NotionDatabasePropertyType = "relation"
	RollUp         NotionDatabasePropertyType = "rollup"
	CreatedTime    NotionDatabasePropertyType = "created_time"
	CreatedBy      NotionDatabasePropertyType = "created_by"
	LastEditedBy   NotionDatabasePropertyType = "last_edited_by"
	LastEditedTime NotionDatabasePropertyType = "last_edited_time"
	Status         NotionDatabasePropertyType = "status"
)
