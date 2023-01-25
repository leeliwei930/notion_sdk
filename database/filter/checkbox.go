package filter

type Checkbox struct {
	Equals    *bool `json:"equals,omitempty"`
	NotEquals *bool `json:"does_not_equals,omitempty"`
}
