package filter

type MultiSelect struct {
	Contains    string `json:"contains,omitempty"`
	NotContains string `json:"does_not_contains,omitempty"`
	IsEmpty     bool   `json:"is_empty,omitempty"`
	IsNotEmpty  bool   `json:"is_not_empty,omitempty"`
}
