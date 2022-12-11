package models

type NumberFormatType string

const (
	NumberFormat NumberFormatType = "number"
	WithCommas   NumberFormatType = "number_with_commas"
	Percent      NumberFormatType = "percent"
)
