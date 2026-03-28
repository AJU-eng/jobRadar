package entities

import (
	"github.com/google/uuid"
)

type JobPost struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Comp_id     uuid.UUID `json:"comp_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Time        string    `json:"time"`
	TimeRange   string    `json:"time_range"`
	Period      string    `json:"period"`
	Location    string    `json:"location"`
}
