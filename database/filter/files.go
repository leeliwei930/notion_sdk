package filter

type Files struct {
	IsEmpty    *bool `json:"is_empty,omitempty"`
	IsNotEmpty *bool `json:"is_not_empty,omitempty"`
}
