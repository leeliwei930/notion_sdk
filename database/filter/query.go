package filter

type QueryProps struct {
	Or          []*QueryProps `json:"or,omitempty"`
	And         []*QueryProps `json:"and,omitempty"`
	Property    string        `json:"property,omitempty"`
	Relation    *Relation     `json:"relation,omitempty"`
	Timestamp   *Timestamp    `json:"timestamp,omitempty"`
	Select      *Select       `json:"select,omitempty"`
	Number      *Number       `json:"number,omitempty"`
	MultiSelect *MultiSelect  `json:"multi_select,omitempty"`
	Status      *Status       `json:"status,omitempty"`
	People      *People       `json:"people,omitempty"`
	Files       *Files        `json:"files,omitempty"`
	RollUp      *RollUp       `json:"rollup,omitempty"`
	Formula     *Formula      `json:"formula,omitempty"`
	Checkbox    *Checkbox     `json:"checkbox,omitempty"`
	RichText    *Text         `json:"rich_text,omitempty"`
	Title       *Text         `json:"title,omitempty"`
}
