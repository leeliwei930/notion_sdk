package enums

type FormulaPropertyValueType uint8

const (
	FormulaPropertyStringValue FormulaPropertyValueType = iota
	FormulaPropertyNumberValue
	FormulaPropertyBooleanValue
	FormulaPropertyDateValue
)

var FormulaPropertyValueTypeMap = map[FormulaPropertyValueType]string{
	FormulaPropertyStringValue:  "string",
	FormulaPropertyNumberValue:  "number",
	FormulaPropertyBooleanValue: "boolean",
	FormulaPropertyDateValue:    "date",
}

func (f FormulaPropertyValueType) String() string {
	return FormulaPropertyValueTypeMap[f]
}
