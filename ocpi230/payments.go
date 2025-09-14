package ocpi230

import "time"

type InvoiceCreator string

const (
	CPO InvoiceCreator = "CPO"
	PTP InvoiceCreator = "PTP"
)

type PaymentCaptureStatus string

const (
	PaymentCaptureStatusSuccess        PaymentCaptureStatus = "SUCCESS"
	PaymentCaptureStatusPartialSuccess PaymentCaptureStatus = "PARTIAL_SUCCESS"
	PaymentCaptureStatusFailed         PaymentCaptureStatus = "FAILED"
)

type Terminal struct {
	TerminalID        string            `json:"terminal_id"`
	CustomerReference *string           `json:"customer_reference,omitempty"`
	PartyID           string            `json:"party_id"`
	CountryCode       string            `json:"country_code"`
	Address           *string           `json:"address,omitempty"`
	City              *string           `json:"city,omitempty"`
	PostalCode        *string           `json:"postal_code,omitempty"`
	State             *string           `json:"state,omitempty"`
	Country           *string           `json:"country,omitempty"`
	Coordinates       *GeoLocation      `json:"coordinates,omitempty"`
	InvoiceBaseUrl    *string           `json:"invoice_base_url,omitempty"`
	InvoiceCreator    *InvoiceCreator   `json:"invoice_creator,omitempty"`
	Reference         *string           `json:"reference,omitempty"`
	LocationIDs       []string          `json:"location_ids,omitempty"`
	EvseUIDs          []string          `json:"evse_uids,omitempty"`
	LastUpdated       time.Time         `json:"last_updated"`
	Metadata          map[string]string `json:"metadata,omitempty"`
}

type FinancialAdviceConfirmation struct {
	ID                     string               `json:"id"`
	AuthorizationReference string               `json:"authorization_reference"`
	TotalCosts             Price                `json:"total_costs"`
	Currency               string               `json:"currency"`
	EFTData                []string             `json:"eft_data"`
	CaptureStatusCode      PaymentCaptureStatus `json:"capture_status_code"`
	CaptureStatusMessage   *string              `json:"capture_status_message,omitempty"`
	LastUpdated            time.Time            `json:"last_updated"`
	Metadata               map[string]string    `json:"metadata,omitempty"`
}
