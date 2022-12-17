package enums

import "encoding/json"

type PageParentType uint8

const (
	PageParentDatabase PageParentType = iota
	PageParentPage
	PageParentWorkspace
)

var pageParentTypeMap = map[PageParentType]string{
	PageParentDatabase:  "database_id",
	PageParentPage:      "page_id",
	PageParentWorkspace: "workspace",
}

var pageParentTypeIndexes = map[string]PageParentType{
	"database_id": PageParentDatabase,
	"page_id":     PageParentPage,
	"workspace":   PageParentWorkspace,
}

func ParseParentType(s string) PageParentType {
	return pageParentTypeIndexes[s]
}
func (p PageParentType) String() string {
	return pageParentTypeMap[p]
}

func (p PageParentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PageParentType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*p = ParseParentType(j)
	return nil
}
