package enums

type MentionType string

const (
	User        MentionType = "user"
	Page        MentionType = "page"
	Database    MentionType = "database"
	Data        MentionType = "data"
	LinkPreview MentionType = "link_preview"
)
