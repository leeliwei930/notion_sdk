package filter

type Formula struct {
	String   *Text   `json:"string,omitempty"`
	Checkbox *Text   `json:"checkbox,omitempty"`
	Number   *Number `json:"number,omitempty"`
	Date     *Date   `json:"date,omitempty"`
}
