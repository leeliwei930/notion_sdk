package enums

import "encoding/json"

type MentionType uint8

const (
	User MentionType = iota
	Page
	Database
	Data
	LinkPreview
)

var mentionTypeMap = map[MentionType]string{
	User:        "user",
	Page:        "page",
	Database:    "database",
	Data:        "data",
	LinkPreview: "link_preview",
}

var mentionTypeIndexes = map[string]MentionType{
	"user":         User,
	"page":         Page,
	"database":     Database,
	"data":         Data,
	"link_preview": LinkPreview,
}

func ParseMentionType(s string) MentionType {
	return mentionTypeIndexes[s]
}

func (m MentionType) String() string {
	return mentionTypeMap[m]
}

func (m MentionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *MentionType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*m = ParseMentionType(j)
	return nil
}
