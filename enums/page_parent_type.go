package enums

type PageParentType int

const (
	PageParentDatabase PageParentType = iota
	PageParentPage
	PageParentWorkspace
)

var PageParentTypeMap = map[PageParentType]string{
	PageParentDatabase:  "database_id",
	PageParentPage:      "page_id",
	PageParentWorkspace: "workspace",
}
