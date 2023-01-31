package filter

import "encoding/json"

type TimestampType uint8

const (
	CreatedTime TimestampType = iota
	LastEditedTime
)

var timestampTypeMap = map[TimestampType]string{
	CreatedTime:    "created_time",
	LastEditedTime: "last_edited_time",
}

var timestampTypeIndexes = map[string]TimestampType{
	"created_time":     CreatedTime,
	"last_edited_time": LastEditedTime,
}

func ParseTimestampType(s string) TimestampType {
	return timestampTypeIndexes[s]
}
func (d TimestampType) String() string {
	return timestampTypeMap[d]
}

func (d TimestampType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *TimestampType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*d = ParseTimestampType(j)
	return nil
}

type Timestamp struct {
	CreatedTime    *Date         `json:"created_time,omitempty"`
	LastEditedTime *Date         `json:"last_edited_time,omitempty"`
	Types          TimestampType `json:"timestamp,omitempty"`
}
