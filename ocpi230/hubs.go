package ocpi230

import "time"

type ConnectionStatus string

const (
	ConnectionStatusConnected ConnectionStatus = "CONNECTED"
	ConnectionStatusOffline   ConnectionStatus = "OFFLINE"
	ConnectionStatusPlanned   ConnectionStatus = "PLANNED"
	ConnectionStatusSuspended ConnectionStatus = "SUSPENDED"
)

type ClientInfo struct {
	CountryCode string            `json:"country_code"`
	PartyID     string            `json:"party_id"`
	Role        Role              `json:"role"`
	Status      ConnectionStatus  `json:"status"`
	LastUpdated time.Time         `json:"last_updated"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}
