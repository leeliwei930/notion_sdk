package filter

// as status type filter props is identical with select type
type Status = Select

type Select struct {
	Equals     *string `json:"equals,omitempty"`
	NotEqual   *string `json:"does_not_equal,omitempty"`
	IsEmpty    *bool   `json:"is_empty,omitempty"`
	IsNotEmpty *bool   `json:"is_not_empty,omitempty"`
}
