package enums

type NumberFormatType string

const (
	NumberFormatTypeGeneral    NumberFormatType = "number"
	NumberFormatTypeWithCommas NumberFormatType = "number_with_commas"
	NumberFormatTypePercent    NumberFormatType = "percent"
)
