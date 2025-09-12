package ocpi211

import "time"

type TokenType string

const (
	TokenTypeOther TokenType = "OTHER"
	TokenTypeRFID  TokenType = "RFID"
)

type WhitelistType string

const (
	WhitelistTypeAlways         WhitelistType = "ALWAYS"
	WhitelistTypeAllowed        WhitelistType = "ALLOWED"
	WhitelistTypeAllowedOffline WhitelistType = "ALLOWED_OFFLINE"
	WhitelistTypeNever          WhitelistType = "NEVER"
)

type TokenAuthIDType string

const (
	TokenAuthIDTypeUID    TokenAuthIDType = "UID"
	TokenAuthIDTypeAuthID TokenAuthIDType = "AUTH_ID"
)

type Token struct {
	UID          string            `json:"uid"`
	Type         TokenType         `json:"type"`
	AuthID       string            `json:"auth_id"`
	VisualNumber *string           `json:"visual_number,omitempty"`
	Issuer       string            `json:"issuer"`
	Valid        bool              `json:"valid"`
	Whitelist    WhitelistType     `json:"whitelist"`
	LanguageCode *string           `json:"language_code,omitempty"`
	LastUpdated  time.Time         `json:"last_updated"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}
