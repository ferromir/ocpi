package ocpi211

import "time"

type SessionStatus string

const (
	SessionStatusActive    SessionStatus = "ACTIVE"
	SessionStatusCompleted SessionStatus = "COMPLETED"
	SessionStatusInvalid   SessionStatus = "INVALID"
	SessionStatusPending   SessionStatus = "PENDING"
)

type Session struct {
	ID              string            `json:"id"`
	StartDateTime   time.Time         `json:"start_date_time"`
	EndDateTime     *time.Time        `json:"end_date_time,omitempty"`
	KWh             float32           `json:"kwh"`
	AuthID          string            `json:"auth_id"`
	AuthMethod      AuthMethod        `json:"auth_method"`
	Location        Location          `json:"location"`
	MeterID         *string           `json:"meter_id,omitempty"`
	Currency        string            `json:"currency"`
	ChargingPeriods []ChargingPeriod  `json:"charging_periods,omitempty"`
	TotalCost       *float32          `json:"total_cost,omitempty"`
	Status          SessionStatus     `json:"status"`
	LastUpdated     time.Time         `json:"last_updated"`
	Metadata        map[string]string `json:"metadata,omitempty"`
}
