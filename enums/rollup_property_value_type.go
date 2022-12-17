package enums

type RollUpPropertyValueType uint8

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

func (r RollUpPropertyValueType) String() string {
	return RollUpPropertyValueTypeMap[r]
}
