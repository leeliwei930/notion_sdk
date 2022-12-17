package enums

import "encoding/json"

type FormulaPropertyValueType uint8

const (
	FormulaPropertyStringValue FormulaPropertyValueType = iota
	FormulaPropertyNumberValue
	FormulaPropertyBooleanValue
	FormulaPropertyDateValue
)

var formulaPropertyValueTypeMap = map[FormulaPropertyValueType]string{
	FormulaPropertyStringValue:  "string",
	FormulaPropertyNumberValue:  "number",
	FormulaPropertyBooleanValue: "boolean",
	FormulaPropertyDateValue:    "date",
}

var formulaPropertyValueTypeIndexes = map[string]FormulaPropertyValueType{
	"string":  FormulaPropertyStringValue,
	"number":  FormulaPropertyNumberValue,
	"boolean": FormulaPropertyBooleanValue,
	"date":    FormulaPropertyDateValue,
}

func ParsePropertyValueType(s string) FormulaPropertyValueType {
	return formulaPropertyValueTypeIndexes[s]
}
func (f FormulaPropertyValueType) String() string {
	return formulaPropertyValueTypeMap[f]
}

func (f FormulaPropertyValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *FormulaPropertyValueType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*f = ParsePropertyValueType(j)
	return nil
}
