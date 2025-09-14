package ocpi230

import "time"

type AllowedType string

const (
	AllowedTypeAllowed    AllowedType = "ALLOWED"
	AllowedTypeBlocked    AllowedType = "BLOCKED"
	AllowedTypeExpired    AllowedType = "EXPIRED"
	AllowedTypeNoCredit   AllowedType = "NO_CREDIT"
	AllowedTypeNotAllowed AllowedType = "NOT_ALLOWED"
)

type WhitelistType string

const (
	WhitelistTypeAlways         WhitelistType = "ALWAYS"
	WhitelistTypeAllowed        WhitelistType = "ALLOWED"
	WhitelistTypeAllowedOffline WhitelistType = "ALLOWED_OFFLINE"
	WhitelistTypeNever          WhitelistType = "NEVER"
)

type ProfileType string

const (
	ProfileTypeCheap   ProfileType = "CHEAP"
	ProfileTypeFast    ProfileType = "FAST"
	ProfileTypeGreen   ProfileType = "GREEN"
	ProfileTypeRegular ProfileType = "REGULAR"
)

type EnergyContract struct {
	SupplierName string  `json:"supplier_name"`
	ContractID   *string `json:"contract_id,omitempty"`
}

type Token struct {
	CountryCode        string            `json:"country_code"`
	PartyID            string            `json:"party_id"`
	UID                string            `json:"uid"`
	Type               TokenType         `json:"type"`
	ContractID         string            `json:"contract_id"`
	VisualNumber       *string           `json:"visual_number,omitempty"`
	Issuer             string            `json:"issuer"`
	GroupID            *string           `json:"group_id,omitempty"`
	Valid              bool              `json:"valid"`
	Whitelist          WhitelistType     `json:"whitelist"`
	Language           *string           `json:"language,omitempty"`
	DefaultProfileType *ProfileType      `json:"default_profile_type,omitempty"`
	EnergyContract     *EnergyContract   `json:"energy_contract,omitempty"`
	LastUpdated        time.Time         `json:"last_updated"`
	Metadata           map[string]string `json:"metadata,omitempty"`
}

type LocationReferences struct {
	LocationID string   `json:"location_id"`
	EvseUIDs   []string `json:"evse_uids,omitempty"`
}

type AuthorizationInfo struct {
	Allowed                AllowedType         `json:"allowed"`
	Token                  Token               `json:"token"`
	Location               *LocationReferences `json:"location"`
	AuthorizationReference *string             `json:"authorization_reference,omitempty"`
	Info                   *DisplayText        `json:"info,omitempty"`
}
