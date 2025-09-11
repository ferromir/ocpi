package v2_3_0

import "time"

// Response envelope for v2.3
type Response[T any] struct {
	Data          T         `json:"data"`
	StatusCode    int       `json:"status_code" validate:"required"`
	StatusMessage string    `json:"status_message,omitempty"`
	Timestamp     time.Time `json:"timestamp" validate:"required"`
}

// Sample CDR (Charge Detail Record) (very simplified)
type CDR struct {
	ID          string    `json:"id" validate:"required"`
	StartTime   time.Time `json:"start_time" validate:"required"`
	EndTime     time.Time `json:"end_time,omitempty"`
	TotalCost   float64   `json:"total_cost,omitempty"`
	Currency    string    `json:"currency,omitempty" validate:"omitempty,len=3"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
