package ocpi221

import "time"

type LocationType string

const (
	LocationTypeOnStreet      LocationType = "ON_STREET"
	LocationTypeParkingGarage LocationType = "PARKING_GARAGE"
	LocationTypeUnderground   LocationType = "UNDERGROUND_GARAGE"
	LocationTypeParkingLot    LocationType = "PARKING_LOT"
	LocationTypeOther         LocationType = "OTHER"
	LocationTypeUnknown       LocationType = "UNKNOWN"
)

type Capability string

const (
	CapabilityChargingProfile        Capability = "CHARGING_PROFILE_CAPABLE"
	CapabilityCreditCardPayable      Capability = "CREDIT_CARD_PAYABLE"
	CapabilityRemoteStartStopCapable Capability = "REMOTE_START_STOP_CAPABLE"
	CapabilityReservable             Capability = "RESERVABLE"
	CapabilityRFIDReader             Capability = "RFID_READER"
	CapabilityUnlockCapable          Capability = "UNLOCK_CAPABLE"
)

type EvseStatus string

const (
	EvseStatusAvailable   EvseStatus = "AVAILABLE"
	EvseStatusBlocked     EvseStatus = "BLOCKED"
	EvseStatusCharging    EvseStatus = "CHARGING"
	EvseStatusInoperative EvseStatus = "INOPERATIVE"
	EvseStatusOutOfOrder  EvseStatus = "OUTOFORDER"
	EvseStatusPlanned     EvseStatus = "PLANNED"
	EvseStatusRemoved     EvseStatus = "REMOVED"
	EvseStatusReserved    EvseStatus = "RESERVED"
	EvseStatusUnknown     EvseStatus = "UNKNOWN"
)

type ParkingRestriction string

const (
	ParkingRestrictionEvOnly      ParkingRestriction = "EV_ONLY"
	ParkingRestrictionPlugged     ParkingRestriction = "PLUGGED"
	ParkingRestrictionDisabled    ParkingRestriction = "DISABLED"
	ParkingRestrictionCustomers   ParkingRestriction = "CUSTOMERS"
	ParkingRestrictionMotorcycles ParkingRestriction = "MOTORCYCLES"
)

type Facility string

const (
	FacilityHotel          Facility = "HOTEL"
	FacilityRestaurant     Facility = "RESTAURANT"
	FacilityCafe           Facility = "CAFE"
	FacilityMall           Facility = "MALL"
	FacilitySupermarket    Facility = "SUPERMARKET"
	FacilitySport          Facility = "SPORT"
	FacilityRecreationArea Facility = "RECREATION_AREA"
	FacilityNature         Facility = "NATURE"
	FacilityMuseum         Facility = "MUSEUM"
	FacilityBusStop        Facility = "BUS_STOP"
	FacilityTaxiStand      Facility = "TAXI_STAND"
	FacilityTrainStation   Facility = "TRAIN_STATION"
	FacilityAirport        Facility = "AIRPORT"
	FacilityCarpoolParking Facility = "CARPOOL_PARKING"
	FacilityFuelStation    Facility = "FUEL_STATION"
	FacilityWifi           Facility = "WIFI"
)

type DayOfWeekNumber int

const (
	DayOfWeekNumberMonday    DayOfWeekNumber = 1
	DayOfWeekNumberTuesday   DayOfWeekNumber = 2
	DayOfWeekNumberWednesday DayOfWeekNumber = 3
	DayOfWeekNumberThursday  DayOfWeekNumber = 4
	DayOfWeekNumberFriday    DayOfWeekNumber = 5
	DayOfWeekNumberSaturday  DayOfWeekNumber = 6
	DayOfWeekNumberSunday    DayOfWeekNumber = 7
)

type RegularHours struct {
	PeriodBegin string          `json:"period_begin"`
	PeriodEnd   string          `json:"period_end"`
	Weekday     DayOfWeekNumber `json:"weekday"`
}

type ExceptionalPeriod struct {
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
}

type Hours struct {
	RegularHours        []RegularHours      `json:"regular_hours,omitempty"`
	TwentyFourSeven     bool                `json:"twentyfourseven"`
	ExceptionalOpenings []ExceptionalPeriod `json:"exceptional_openings,omitempty"`
	ExceptionalClosings []ExceptionalPeriod `json:"exceptional_closings,omitempty"`
}

type StatusSchedule struct {
	PeriodBegin time.Time  `json:"period_begin"`
	PeriodEnd   *time.Time `json:"period_end,omitempty"`
	Status      EvseStatus `json:"status"`
}

type AdditionalGeoLocation struct {
	Latitude  string       `json:"latitude"`
	Longitude string       `json:"longitude"`
	Name      *DisplayText `json:"name,omitempty"`
}

type Connector struct {
	ID                 string          `json:"id"`
	Standard           ConnectorType   `json:"standard"`
	Format             ConnectorFormat `json:"format"`
	PowerType          PowerType       `json:"power_type"`
	Voltage            int             `json:"voltage"`
	Amperage           int             `json:"amperage"`
	TariffID           *string         `json:"tariff_id,omitempty"`
	TermsAndConditions *string         `json:"terms_and_conditions,omitempty"`
	LastUpdated        time.Time       `json:"last_updated"`
}

type EVSE struct {
	UID                 string               `json:"uid"`
	EvseID              *string              `json:"evse_id,omitempty"`
	Status              EvseStatus           `json:"status"`
	StatusSchedule      []StatusSchedule     `json:"status_schedule,omitempty"`
	Capabilities        []Capability         `json:"capabilities,omitempty"`
	Connectors          []Connector          `json:"connectors"`
	FloorLevel          *string              `json:"floor_level,omitempty"`
	Coordinates         *GeoLocation         `json:"coordinates"`
	PhysicalReference   *string              `json:"physical_reference,omitempty"`
	Directions          []DisplayText        `json:"directions,omitempty"`
	ParkingRestrictions []ParkingRestriction `json:"parking_restrictions,omitempty"`
	Images              []Image              `json:"images,omitempty"`
	LastUpdated         time.Time            `json:"last_updated"`
}

type Location struct {
	ID                 string                  `json:"id"`
	Type               LocationType            `json:"type"`
	Name               *string                 `json:"name,omitempty"`
	Address            string                  `json:"address"`
	City               string                  `json:"city"`
	PostalCode         string                  `json:"postal_code"`
	Country            string                  `json:"country"`
	Coordinates        GeoLocation             `json:"coordinates"`
	RelatedLocations   []AdditionalGeoLocation `json:"related_locations,omitempty"`
	Evses              []EVSE                  `json:"evses,omitempty"`
	Directions         []DisplayText           `json:"directions,omitempty"`
	Operator           *BusinessDetails        `json:"operator,omitempty"`
	Suboperator        *BusinessDetails        `json:"suboperator,omitempty"`
	Owner              *BusinessDetails        `json:"owner,omitempty"`
	Facilities         []Facility              `json:"facilities,omitempty"`
	TimeZone           *string                 `json:"time_zone,omitempty"` // IANA
	OpeningTimes       *Hours                  `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                   `json:"charging_when_closed,omitempty"`
	Images             []Image                 `json:"images,omitempty"`
	EnergyMix          *EnergyMix              `json:"energy_mix,omitempty"`
	LastUpdated        time.Time               `json:"last_updated"`
	Metadata           map[string]string       `json:"metadata,omitempty"`
}
