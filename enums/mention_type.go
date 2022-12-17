package enums

type MentionType int

const (
	User MentionType = iota
	Page
	Database
	Data
	LinkPreview
)

var MentionTypeMap = map[MentionType]string{
	User:        "user",
	Page:        "page",
	Database:    "database",
	Data:        "data",
	LinkPreview: "link_preview",
}
