package ocpi221

import "time"

type TariffType string

const (
	TariffTypeAdHocPayment TariffType = "AD_HOC_PAYMENT"
	TariffTypeProfileCheap TariffType = "PROFILE_CHEAP"
	TariffTypeProfileFast  TariffType = "PROFILE_FAST"
	TariffTypeProfileGreen TariffType = "PROFILE_GREEN"
	TariffTypeRegular      TariffType = "REGULAR"
)

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

type ReservationRestrictionType string

const (
	ReservationRestrictionTypeReservation        ReservationRestrictionType = "RESERVATION"
	ReservationRestrictionTypeReservationExpires ReservationRestrictionType = "RESERVATION_EXPIRES"
)

type PriceComponent struct {
	Type     TariffDimensionType `json:"type"`
	Price    float64             `json:"price"`
	Vat      *float64            `json:"vat,omitempty"`
	StepSize int                 `json:"step_size"`
}

type TariffRestriction struct {
	StartTime   *string                     `json:"start_time,omitempty"`
	EndTime     *string                     `json:"end_time,omitempty"`
	StartDate   *string                     `json:"start_date,omitempty"`
	EndDate     *string                     `json:"end_date,omitempty"`
	MinKWh      *float64                    `json:"min_kwh,omitempty"`
	MaxKWh      *float64                    `json:"max_kwh,omitempty"`
	MinCurrent  *float64                    `json:"min_current,omitempty"`
	MaxCurrent  *float64                    `json:"max_current,omitempty"`
	MinPower    *float64                    `json:"min_power,omitempty"`
	MaxPower    *float64                    `json:"max_power,omitempty"`
	MinDuration *int                        `json:"min_duration,omitempty"`
	MaxDuration *int                        `json:"max_duration,omitempty"`
	DayOfWeek   []DayOfWeekText             `json:"day_of_week,omitempty"`
	Reservation *ReservationRestrictionType `json:"reservation,omitempty"`
}

type TariffElement struct {
	PriceComponents []PriceComponent   `json:"price_components"`
	Restrictions    *TariffRestriction `json:"restrictions,omitempty"`
}

type Tariff struct {
	CountryCode   string            `json:"country_code"`
	PartyID       string            `json:"party_id"`
	ID            string            `json:"id"`
	Currency      string            `json:"currency"`
	Type          *TariffType       `json:"type,omitempty"`
	TariffAltText []DisplayText     `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string           `json:"tariff_alt_url,omitempty"`
	MinPrice      *Price            `json:"min_price,omitempty"`
	MaxPrice      *Price            `json:"max_price,omitempty"`
	Elements      []TariffElement   `json:"elements"`
	StartDateTime *time.Time        `json:"start_date_time,omitempty"`
	EndDateTime   *time.Time        `json:"end_date_time,omitempty"`
	EnergyMix     *EnergyMix        `json:"energy_mix,omitempty"`
	LastUpdated   time.Time         `json:"last_updated"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}
