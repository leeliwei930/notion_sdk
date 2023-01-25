package filter

type RollUp struct {
	Any    QueryProps `json:"any,omitempty"`
	Every  QueryProps `json:"every,omitempty"`
	None   QueryProps `json:"none,omitempty"`
	Number *Number    `json:"number,omitempty"`
	Date   *Date      `json:"date,omitempty"`
}
