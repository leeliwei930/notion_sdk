package filter

type QueryProps struct {
	Property    string        `json:"property,omitempty"`
	RichText    *Text         `json:"rich_text,omitempty"`
	Checkbox    *Checkbox     `json:"checkbox,omitempty"`
	Number      *Number       `json:"number,omitempty"`
	Timestamp   *Timestamp    `json:"timestamp,omitempty"`
	Select      *Select       `json:"select,omitempty"`
	MultiSelect *MultiSelect  `json:"multi_select,omitempty"`
	Status      *Status       `json:"status,omitempty"`
	And         []*QueryProps `json:"and,omitempty"`
	Or          []*QueryProps `json:"or,omitempty"`
}
