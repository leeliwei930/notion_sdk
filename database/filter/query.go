package filter

type QueryProps struct {
	Property string        `json:"property,omitempty"`
	RichText *Text         `json:"rich_text,omitempty"`
	Checkbox *Checkbox     `json:"checkbox,omitempty"`
	Number   *Number       `json:"number,omitempty"`
	And      []*QueryProps `json:"and,omitempty"`
	Or       []*QueryProps `json:"or,omitempty"`
}
