package ocpi221

import "time"

type CommandType string

const (
	CommandTypeStartSession      CommandType = "START_SESSION"
	CommandTypeStopSession       CommandType = "STOP_SESSION"
	CommandTypeReserveNow        CommandType = "RESERVE_NOW"
	CommandTypeCancelReservation CommandType = "CANCEL_RESERVATION"
	CommandTypeUnlockConnector   CommandType = "UNLOCK_CONNECTOR"
)

type CommandResponseType string

const (
	CommandResponseTypeNotSupported   CommandResponseType = "NOT_SUPPORTED"
	CommandResponseTypeRejected       CommandResponseType = "REJECTED"
	CommandResponseTypeAccepted       CommandResponseType = "ACCEPTED"
	CommandResponseTypeUnknownSession CommandResponseType = "UNKNOWN_SESSION"
)

type CommandResultType string

const (
	CommandResultTypeAccepted            CommandResultType = "ACCEPTED"
	CommandResultTypeCanceledReservation CommandResultType = "CANCELED_RESERVATION"
	CommandResultTypeEvseOccupied        CommandResultType = "EVSE_OCCUPIED"
	CommandResultTypeEvseInoperative     CommandResultType = "EVSE_INOPERATIVE"
	CommandResultTypeFailed              CommandResultType = "FAILED"
	CommandResultTypeNotSupported        CommandResultType = "NOT_SUPPORTED"
	CommandResultTypeRejected            CommandResultType = "REJECTED"
	CommandResultTypeTimeout             CommandResultType = "TIMEOUT"
	CommandResultTypeUnknownReservation  CommandResultType = "UNKNOWN_RESERVATION"
)

type StartSession struct {
	ResponseUrl string `json:"response_url"`
	// Token                  Token             `json:"token"`
	LocationID             string            `json:"location_id"`
	EvseUID                *string           `json:"evse_uid,omitempty"`
	ConnectorID            *string           `json:"connector_id,omitempty"`
	AuthorizationReference *string           `json:"authorization_reference,omitempty"`
	Metadata               map[string]string `json:"metadata,omitempty"`
}

type StopSession struct {
	ResponseUrl string `json:"response_url"`
	SessionID   string `json:"session_id"`
}

type CancelReservation struct {
	ResponseUrl   string `json:"response_url"`
	ReservationID int    `json:"reservation_id"`
}

type ReserveNow struct {
	ResponseUrl string `json:"response_url"`
	// Token                  Token     `json:"token"`
	ExpiryDate             time.Time `json:"expiry_date"`
	ReservationID          int       `json:"reservation_id"`
	LocationID             string    `json:"location_id"`
	EvseUID                *string   `json:"evse_uid,omitempty"`
	AuthorizationReference *string   `json:"authorization_reference,omitempty"`
}

type UnlockConnector struct {
	ResponseUrl string `json:"response_url"`
	LocationID  string `json:"location_id"`
	EvseUID     string `json:"evse_uid"`
	ConnectorID string `json:"connector_id"`
}

type CommandResponse struct {
	Result  CommandResponseType `json:"result"`
	Timeout int                 `json:"timeout"`
	Message *DisplayText        `json:"message,omitempty"`
}

type CommandResult struct {
	Result  CommandResultType `json:"result"`
	Message *DisplayText      `json:"message,omitempty"`
}
