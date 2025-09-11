package v2_2_1

import "time"

// Response envelope for v2.2
type Response[T any] struct {
	Data          T         `json:"data"`
	StatusCode    int       `json:"status_code" validate:"required"`
	StatusMessage string    `json:"status_message,omitempty"`
	Timestamp     time.Time `json:"timestamp" validate:"required"`
}

// Sample Tariff type (simplified)
type Tariff struct {
	ID          string    `json:"id" validate:"required"`
	Currency    string    `json:"currency" validate:"required,len=3"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
