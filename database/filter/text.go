package filter

type Text struct {
	Equals      string `json:"equals,omitempty"`
	NotEqual    string `json:"does_not_equal,omitempty"`
	StartsWith  string `json:"starts_with,omitempty"`
	EndsWith    string `json:"ends_with,omitempty"`
	Contains    string `json:"contains,omitempty"`
	NotContains string `json:"does_not_contains,omitempty"`
	IsEmpty     bool   `json:"is_empty,omitempty"`
	IsNotEmpty  bool   `json:"is_not_empty,omitempty"`
}
