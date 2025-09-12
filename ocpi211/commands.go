package ocpi211

import (
	"time"
)

type CommandType string

const (
	CommandTypeStartSession    CommandType = "START_SESSION"
	CommandTypeStopSession     CommandType = "STOP_SESSION"
	CommandTypeReserveNow      CommandType = "RESERVE_NOW"
	CommandTypeUnlockConnector CommandType = "UNLOCK_CONNECTOR"
)

type CommandResponseType string

const (
	CommandResponseTypeNotSupported   CommandResponseType = "NOT_SUPPORTED"
	CommandResponseTypeRejected       CommandResponseType = "REJECTED"
	CommandResponseTypeAccepted       CommandResponseType = "ACCEPTED"
	CommandResponseTypeTimeout        CommandResponseType = "TIMEOUT"
	CommandResponseTypeUnknownSession CommandResponseType = "UNKNOWN_SESSION"
)

type StartSession struct {
	ResponseUrl string            `json:"response_url"`
	Token       Token             `json:"token"`
	LocationID  string            `json:"location_id"`
	EvseUID     *string           `json:"evse_uid,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type StopSession struct {
	ResponseUrl string `json:"response_url"`
	SessionID   string `json:"session_id"`
}

type ReserveNow struct {
	ResponseUrl   string    `json:"response_url"`
	Token         Token     `json:"token"`
	ExpiryDate    time.Time `json:"expiry_date"`
	ReservationID int       `json:"reservation_id"`
	LocationID    string    `json:"location_id"`
	EvseUID       *string   `json:"evse_uid,omitempty"`
}

type UnlockConnector struct {
	ResponseUrl string `json:"response_url"`
	LocationID  string `json:"location_id"`
	EvseUID     string `json:"evse_uid"`
	ConnectorID string `json:"connector_id"`
}

type CommandResponse struct {
	Result CommandResponseType `json:"result"`
}
