package filter

type Number struct {
	Equals     float64 `json:"equals,omitempty"`
	NotEqual   float64 `json:"does_not_equal,omitempty"`
	Gt         float64 `json:"greater_than,omitempty"`
	Lt         float64 `json:"less_than,omitempty"`
	Gte        float64 `json:"greater_than_or_equal_to,omitempty"`
	Lte        float64 `json:"less_than_or_equal_to,omitempty"`
	IsEmpty    bool    `json:"is_empty,omitempty"`
	IsNotEmpty bool    `json:"is_not_empty,omitempty"`
}
