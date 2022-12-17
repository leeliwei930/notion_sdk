package enums

type RollUpPropertyValueType int

const (
	RollUpNumberPropertyValue RollUpPropertyValueType = iota
	RollUpDatePropertyValue
	RollUpArrayPropertyValue
)

var RollUpPropertyValueTypeMap = map[RollUpPropertyValueType]string{
	RollUpNumberPropertyValue: "number",
	RollUpDatePropertyValue:   "date",
	RollUpArrayPropertyValue:  "array",
}
