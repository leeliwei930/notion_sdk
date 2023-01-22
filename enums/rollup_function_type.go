package enums

import "encoding/json"

type RollUpFunctionType uint8

const (
	Count RollUpFunctionType = iota
	CountValues
	Empty
	NotEmpty
	Unique
	ShowUnique
	PercentEmpty
	PercentNotEmpty
	Sum
	Average
	Median
	Min
	Max
	Range
	EarliestDate
	LatestDate
	DateRange
	Checked
	Unchecked
	PercentChecked
	PercentUnchecked
	CountPerGroup
	PercentPerGroup
	ShowOriginal
)

var rollUpFunctionEnumMap = map[RollUpFunctionType]string{
	Count:            "count",
	CountValues:      "count_values",
	Empty:            "empty",
	NotEmpty:         "not_empty",
	Unique:           "unique",
	ShowUnique:       "show_unique",
	PercentEmpty:     "percent_empty",
	PercentNotEmpty:  "percent_not_empty",
	Sum:              "sum",
	Average:          "average",
	Median:           "median",
	Min:              "min",
	Max:              "max",
	Range:            "range",
	EarliestDate:     "earliest_date",
	LatestDate:       "latest_date",
	DateRange:        "date_range",
	Checked:          "checked",
	Unchecked:        "unchecked",
	PercentChecked:   "percent_checked",
	PercentUnchecked: "percent_unchecked",
	CountPerGroup:    "count_per_group",
	PercentPerGroup:  "percent_per_group",
	ShowOriginal:     "show_original",
}

var rollUpFunctionIndexes = map[string]RollUpFunctionType{
	"count":             Count,
	"count_values":      CountValues,
	"empty":             Empty,
	"not_empty":         NotEmpty,
	"unique":            Unique,
	"show_unique":       ShowUnique,
	"percent_empty":     PercentEmpty,
	"percent_not_empty": PercentNotEmpty,
	"sum":               Sum,
	"average":           Average,
	"median":            Median,
	"min":               Min,
	"max":               Max,
	"range":             Range,
	"earliest_date":     EarliestDate,
	"latest_date":       LatestDate,
	"date_range":        DateRange,
	"checked":           Checked,
	"unchecked":         Unchecked,
	"percent_checked":   PercentChecked,
	"percent_unchecked": PercentUnchecked,
	"count_per_group":   CountPerGroup,
	"percent_per_group": PercentPerGroup,
	"show_original":     ShowOriginal,
}

func ParseRollUpFunctionType(s string) RollUpFunctionType {
	return rollUpFunctionIndexes[s]
}

func (r RollUpFunctionType) String() string {
	return rollUpFunctionEnumMap[r]
}

func (r RollUpFunctionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *RollUpFunctionType) UnmarshalJSON(d []byte) error {
	var j string
	err := json.Unmarshal(d, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*r = ParseRollUpFunctionType(j)
	return nil
}
