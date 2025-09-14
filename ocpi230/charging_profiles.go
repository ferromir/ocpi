package ocpi230

import "time"

type ChargingProfileResultType string

const (
	ChargingProfileResultTypeAccepted ChargingProfileResultType = "ACCEPTED"
	ChargingProfileResultTypeRejected ChargingProfileResultType = "REJECTED"
	ChargingProfileResultTypeUnknown  ChargingProfileResultType = "UNKNOWN"
)

type ChargingRateUnit string

const (
	ChargingRateUnitW ChargingRateUnit = "W"
	ChargingRateUnitA ChargingRateUnit = "A"
)

type ChargingProfileResponseType string

const (
	ChargingProfileResponseTypeAccepted       ChargingProfileResponseType = "ACCEPTED"
	ChargingProfileResponseTypeNotSupported   ChargingProfileResponseType = "NOT_SUPPORTED"
	ChargingProfileResponseTypeRejected       ChargingProfileResponseType = "REJECTED"
	ChargingProfileResponseTypeTooOften       ChargingProfileResponseType = "TOO_OFTEN"
	ChargingProfileResponseTypeUnknownSession ChargingProfileResponseType = "UNKNOWN_SESSION"
)

type ChargingProfilePeriod struct {
	StartPeriod int     `json:"start_period"`
	Limit       float32 `json:"limit"`
}

type ChargingProfile struct {
	StartDateTime         *time.Time             `json:"start_date_time,omitempty"`
	Duration              *int                   `json:"duration,omitempty"`
	ChargingRateUnit      ChargingRateUnit       `json:"charging_rate_unit"`
	MinChargeRate         *float32               `json:"min_charge_rate,omitempty"`
	ChargingProfilePeriod *ChargingProfilePeriod `json:"charging_profile_period,omitempty"`
}

type ActiveChargingProfile struct {
	StartDateTime   time.Time       `json:"start_date_time"`
	ChargingProfile ChargingProfile `json:"charging_profile"`
}

type SetChargingProfile struct {
	ChargingProfile ChargingProfile `json:"charging_profile"`
	ResponseUrl     string          `json:"response_url"`
}

type ChargingProfileResult struct {
	Result ChargingProfileResultType `json:"result"`
}

type ActiveChargingProfileResult struct {
	Result  ChargingProfileResultType `json:"result"`
	Profile *ActiveChargingProfile    `json:"profile,omitempty"`
}

type ChargingProfileResponse struct {
	Result  ChargingProfileResponseType `json:"result"`
	Timeout int                         `json:"timeout"`
}
