package filter

import "github.com/google/uuid"

type People struct {
	Contains    *uuid.UUID `json:"contains,omitempty"`
	NotContains *uuid.UUID `json:"does_not_contains,omitempty"`
	IsEmpty     *bool      `json:"is_empty,omitempty"`
	IsNotEmpty  *bool      `json:"is_not_empty,omitempty"`
}
