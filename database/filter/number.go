package filter

type Number struct {
	Equals     any   `json:"equals,omitempty"`
	NotEqual   any   `json:"does_not_equal,omitempty"`
	Gt         any   `json:"greater_than,omitempty"`
	Lt         any   `json:"less_than,omitempty"`
	Gte        any   `json:"greater_than_or_equal_to,omitempty"`
	Lte        any   `json:"less_than_or_equal_to,omitempty"`
	IsEmpty    *bool `json:"is_empty,omitempty"`
	IsNotEmpty *bool `json:"is_not_empty,omitempty"`
}
