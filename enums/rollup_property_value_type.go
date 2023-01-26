package enums

import "encoding/json"

type RollUpPropertyValueType uint8

const (
	RollUpNumberPropertyValue RollUpPropertyValueType = iota
	RollUpDatePropertyValue
	RollUpArrayPropertyValue
)

var rollUpPropertyValueTypeMap = map[RollUpPropertyValueType]string{
	RollUpNumberPropertyValue: "number",
	RollUpDatePropertyValue:   "date",
	RollUpArrayPropertyValue:  "array",
}
var rollUpPropertyValueTypeIndexes = map[string]RollUpPropertyValueType{
	"number": RollUpNumberPropertyValue,
	"date":   RollUpDatePropertyValue,
	"array":  RollUpArrayPropertyValue,
}

func ParseRollUpPropertyValueType(s string) RollUpPropertyValueType {
	return rollUpPropertyValueTypeIndexes[s]
}

func (r RollUpPropertyValueType) String() string {
	return rollUpPropertyValueTypeMap[r]
}

func (r RollUpPropertyValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *RollUpPropertyValueType) UnmarshalJSON(d []byte) error {
	var j string
	err := json.Unmarshal(d, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*r = ParseRollUpPropertyValueType(j)
	return nil
}
