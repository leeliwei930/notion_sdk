package enums

type MentionType uint8

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

func (m MentionType) String() string {
	return MentionTypeMap[m]
}
