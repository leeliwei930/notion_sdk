package enums

import "encoding/json"

type RelationDatabasePropertyType uint8

const (
	Single RelationDatabasePropertyType = iota
	Dual
)

var relationDatabasePropertyTypeMap = map[RelationDatabasePropertyType]string{
	Single: "single_property",
	Dual:   "dual_property",
}

var relationDatabasePropertyTypeIndexes = map[string]RelationDatabasePropertyType{
	"single_property": Single,
	"dual_property":   Dual,
}

func ParseRelationDatabasePropertyType(s string) RelationDatabasePropertyType {
	return relationDatabasePropertyTypeIndexes[s]
}

func (rp RelationDatabasePropertyType) String() string {
	return relationDatabasePropertyTypeMap[rp]
}

func (rp RelationDatabasePropertyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(rp.String())
}

func (p *RelationDatabasePropertyType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*p = ParseRelationDatabasePropertyType(j)
	return nil
}
