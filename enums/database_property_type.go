package enums

type DatabasePropertyType string

const (
	Title          DatabasePropertyType = "title"
	RichText       DatabasePropertyType = "rich_text"
	Number         DatabasePropertyType = "number"
	Select         DatabasePropertyType = "select"
	MultiSelect    DatabasePropertyType = "multi_select"
	Date           DatabasePropertyType = "date"
	People         DatabasePropertyType = "people"
	Files          DatabasePropertyType = "files"
	Checkbox       DatabasePropertyType = "checkbox"
	Url            DatabasePropertyType = "url"
	Email          DatabasePropertyType = "email"
	PhoneNumber    DatabasePropertyType = "phone_number"
	Formula        DatabasePropertyType = "formula"
	Relation       DatabasePropertyType = "relation"
	RollUp         DatabasePropertyType = "rollup"
	CreatedTime    DatabasePropertyType = "created_time"
	CreatedBy      DatabasePropertyType = "created_by"
	LastEditedBy   DatabasePropertyType = "last_edited_by"
	LastEditedTime DatabasePropertyType = "last_edited_time"
	Status         DatabasePropertyType = "status"
)
