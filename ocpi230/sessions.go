package ocpi230

// import "time"

// type SessionStatus string

// const (
// 	SessionStatusActive      SessionStatus = "ACTIVE"
// 	SessionStatusCompleted   SessionStatus = "COMPLETED"
// 	SessionStatusInvalid     SessionStatus = "INVALID"
// 	SessionStatusPending     SessionStatus = "PENDING"
// 	SessionStatusReservation SessionStatus = "RESERVATION"
// )

// type Session struct {
// 	CountryCode            string            `json:"country_code"`
// 	PartyID                string            `json:"party_id"`
// 	ID                     string            `json:"id"`
// 	StartDateTime          time.Time         `json:"start_date_time"`
// 	EndDateTime            *time.Time        `json:"end_date_time,omitempty"`
// 	KWh                    float64           `json:"kwh"`
// 	CDRToken               CDRToken          `json:"cdr_token"`
// 	AuthMethod             AuthMethod        `json:"auth_method"`
// 	AuthorizationReference *string           `json:"authorization_reference,omitempty"`
// 	LocationID             string            `json:"location_id"`
// 	EvseUID                string            `json:"evse_uid"`
// 	ConnectorID            string            `json:"connector_id"`
// 	MeterID                *string           `json:"meter_id,omitempty"`
// 	Currency               string            `json:"currency"`
// 	ChargingPeriods        []ChargingPeriod  `json:"charging_periods,omitempty"`
// 	TotalCost              *Price            `json:"total_cost,omitempty"`
// 	Status                 SessionStatus     `json:"status"`
// 	LastUpdated            time.Time         `json:"last_updated"`
// 	Metadata               map[string]string `json:"metadata,omitempty"`
// }
