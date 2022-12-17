package enums

type PageParentType uint8

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

func (p PageParentType) String() string {
	return PageParentTypeMap[p]
}
