package ocpi211

import "time"

type StatusCode int

const (
	StatusCodeSuccess              StatusCode = 1000
	StatusCodeClientError          StatusCode = 2000
	StatusCodeInvalidParameters    StatusCode = 2001
	StatusCodeNotEnoughInformation StatusCode = 2002
	StatusCodeUnknownLocation      StatusCode = 2003
	StatusCodeServerError          StatusCode = 3000
	StatusCodeUnableToUseApi       StatusCode = 3001
	StatusCodeUnsupportedVersion   StatusCode = 3002
	StatusCodeNoMatchingEndpoints  StatusCode = 3003
)

type Response[T any] struct {
	Data          T          `json:"data"`
	StatusCode    StatusCode `json:"status_code"`
	StatusMessage *string    `json:"status_message,omitempty"`
	Timestamp     time.Time  `json:"timestamp"`
}
