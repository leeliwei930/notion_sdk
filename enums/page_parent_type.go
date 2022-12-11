package enums

type PageParentType string

const (
	PageParentDatabase  PageParentType = "database_id"
	PageParentPage      PageParentType = "page_id"
	PageParentWorkspace PageParentType = "workspace"
)
