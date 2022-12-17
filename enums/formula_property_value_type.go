package enums

type FormulaPropertyValueType int

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
