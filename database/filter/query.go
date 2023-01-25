package filter

type QueryProps struct {
	Property    string        `json:"property,omitempty"`
	RichText    *Text         `json:"rich_text,omitempty"`
	Checkbox    *Checkbox     `json:"checkbox,omitempty"`
	Number      *Number       `json:"number,omitempty"`
	Timestamp   *Timestamp    `json:"timestamp,omitempty"`
	Select      *Select       `json:"select,omitempty"`
	Relation    *Relation     `json:"relation,omitempty"`
	MultiSelect *MultiSelect  `json:"multi_select,omitempty"`
	Status      *Status       `json:"status,omitempty"`
	People      *People       `json:"people,omitempty"`
	Files       *Files        `json:"files,omitempty"`
	RollUp      *RollUp       `json:"rollup,omitempty"`
	Formula     *Formula      `json:"formula,omitempty"`
	And         []*QueryProps `json:"and,omitempty"`
	Or          []*QueryProps `json:"or,omitempty"`
}
