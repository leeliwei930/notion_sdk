package filter

import (
	"time"
)

type Date struct {
	Equals     *time.Time     `json:"equals,omitempty"`
	Before     *time.Time     `json:"before,omitempty"`
	After      *time.Time     `json:"after,omitempty"`
	OnOrBefore *time.Time     `json:"on_or_before,omitempty"`
	OnOrAfter  *time.Time     `json:"on_or_after,omitempty"`
	PastWeek   *EmptyProperty `json:"past_week,omitempty"`
	ThisWeek   *EmptyProperty `json:"this_week,omitempty"`
	NextWeeks  *EmptyProperty `json:"next_week,omitempty"`
	PastMonth  *EmptyProperty `json:"past_month,omitempty"`
	PastYear   *EmptyProperty `json:"past_year,omitempty"`
	NextMonth  *EmptyProperty `json:"next_month,omitempty"`
	NextYear   *EmptyProperty `json:"next_year,omitempty"`
}
