package enums

type FormulaPropertyValueType string

const (
	FormulaPropertyStringValue  FormulaPropertyValueType = "string"
	FormulaPropertyNumberValue  FormulaPropertyValueType = "number"
	FormulaPropertyBooleanValue FormulaPropertyValueType = "boolean"
	FormulaPropertyDateValue    FormulaPropertyValueType = "date"
)
