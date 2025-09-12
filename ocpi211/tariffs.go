package ocpi211

import "time"

type TariffDimensionType string

const (
	TariffDimensionTypeEnergy      TariffDimensionType = "ENERGY"
	TariffDimensionTypeFlat        TariffDimensionType = "FLAT"
	TariffDimensionTypeParkingTime TariffDimensionType = "PARKING_TIME"
	TariffDimensionTypeTime        TariffDimensionType = "TIME"
)

type DayOfWeekText string

const (
	DayOfWeekTextMonday    DayOfWeekText = "MONDAY"
	DayOfWeekTextTuesday   DayOfWeekText = "TUESDAY"
	DayOfWeekTextWednesday DayOfWeekText = "WEDNESDAY"
	DayOfWeekTextThursday  DayOfWeekText = "THURSDAY"
	DayOfWeekTextFriday    DayOfWeekText = "FRIDAY"
	DayOfWeekTextSaturday  DayOfWeekText = "SATURDAY"
	DayOfWeekTextSunday    DayOfWeekText = "SUNDAY"
)

type PriceComponent struct {
	Type     TariffDimensionType `json:"type"`
	Price    float64             `json:"price"`
	StepSize int                 `json:"step_size"`
}

type TariffRestriction struct {
	StartTime   *string         `json:"start_time,omitempty"` // HH:MM
	EndTime     *string         `json:"end_time,omitempty"`   // HH:MM
	StartDate   *string         `json:"start_date,omitempty"` // YYYY-MM-DD
	EndDate     *string         `json:"end_date,omitempty"`
	MinKWh      *float64        `json:"min_kwh,omitempty"`
	MaxKWh      *float64        `json:"max_kwh,omitempty"`
	MinPower    *float64        `json:"min_power,omitempty"`
	MaxPower    *float64        `json:"max_power,omitempty"`
	MinDuration *int            `json:"min_duration,omitempty"` // in seconds
	MaxDuration *int            `json:"max_duration,omitempty"`
	DayOfWeek   []DayOfWeekText `json:"day_of_week,omitempty"`
}

type TariffElement struct {
	PriceComponents []PriceComponent   `json:"price_components"`
	Restrictions    *TariffRestriction `json:"restrictions,omitempty"`
}

type Tariff struct {
	ID            string          `json:"id"`
	Currency      string          `json:"currency"` // ISO 4217
	TariffAltText []DisplayText   `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string         `json:"tariff_alt_url,omitempty"`
	Elements      []TariffElement `json:"elements"`
	EnergyMix     *EnergyMix      `json:"energy_mix,omitempty"`
	LastUpdated   time.Time       `json:"last_updated"`
}
