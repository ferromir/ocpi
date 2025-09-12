package ocpi221

import "time"

type Capability string

const (
	CapabilityChargingProfileCapable        Capability = "CHARGING_PROFILE_CAPABLE"
	CapabilityChargingPreferencesCapable    Capability = "CHARGING_PREFERENCES_CAPABLE"
	CapabilityChipCardSupport               Capability = "CHIP_CARD_SUPPORT"
	CapabilityContactlessCardSupport        Capability = "CONTACTLESS_CARD_SUPPORT"
	CapabilityCreditCardPayable             Capability = "CREDIT_CARD_PAYABLE"
	CapabilityDebitCardPayable              Capability = "DEBIT_CARD_PAYABLE"
	CapabilityPedTerminal                   Capability = "PED_TERMINAL"
	CapabilityRemoteStartStopCapable        Capability = "REMOTE_START_STOP_CAPABLE"
	CapabilityReservable                    Capability = "RESERVABLE"
	CapabilityRfidReader                    Capability = "RFID_READER"
	CapabilityStartSessionConnectorRequired Capability = "START_SESSION_CONNECTOR_REQUIRED"
	CapabilityTokenGroupCapable             Capability = "TOKEN_GROUP_CAPABLE"
	CapabilityUnlockCapable                 Capability = "UNLOCK_CAPABLE"
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

type ParkingType string

const (
	ParkingTypeAlongMotorway     ParkingType = "ALONG_MOTORWAY"
	ParkingTypeParkingGarage     ParkingType = "PARKING_GARAGE"
	ParkingTypeParkingLot        ParkingType = "PARKING_LOT"
	ParkingTypeOnDriveway        ParkingType = "ON_DRIVEWAY"
	ParkingTypeOnStreet          ParkingType = "ON_STREET"
	ParkingTypeUndergroundGarage ParkingType = "UNDERGROUND_GARAGE"
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
	FacilityBikeSharing    Facility = "BIKE_SHARING"
	FacilityBusStop        Facility = "BUS_STOP"
	FacilityTaxiStand      Facility = "TAXI_STAND"
	FacilityTramStop       Facility = "TRAM_STOP"
	FacilityMetroStation   Facility = "METRO_STATION"
	FacilityTrainStation   Facility = "TRAIN_STATION"
	FacilityAirport        Facility = "AIRPORT"
	FacilityParkingLot     Facility = "PARKING_LOT"
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
	Weekday     DayOfWeekNumber `json:"weekday"`
	PeriodBegin string          `json:"period_begin"`
	PeriodEnd   string          `json:"period_end"`
}

type ExceptionalPeriod struct {
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
}

type Hours struct {
	TwentyFourSeven     bool                `json:"twentyfourseven"`
	RegularHours        []RegularHours      `json:"regular_hours,omitempty"`
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
	MaxVoltage         int             `json:"max_voltage"`
	MaxAmperage        int             `json:"max_amperage"`
	MaxElectricPower   *int            `json:"max_electric_power,omitempty"`
	TariffIDs          *string         `json:"tariff_ids,omitempty"`
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

type PublishTokenType struct {
	UID          *string    `json:"uid,omitempty"`
	TokenType    *TokenType `json:"token_type,omitempty"`
	VisualNumber *string    `json:"visual_number,omitempty"`
	Issuer       *string    `json:"issuer,omitempty"`
	GroupID      *string    `json:"groupId,omitempty"`
}

type Location struct {
	CountryCode        string                  `json:"country_code"`
	PartyID            string                  `json:"party_id"`
	ID                 string                  `json:"id"`
	Publish            bool                    `json:"publish"`
	PublishAllowedTo   *PublishTokenType       `json:"publish_allowed_to,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	Address            string                  `json:"address"`
	City               string                  `json:"city"`
	PostalCode         *string                 `json:"postal_code,omitempty"`
	State              *string                 `json:"state,omitempty"`
	Country            string                  `json:"country"`
	Coordinates        GeoLocation             `json:"coordinates"`
	RelatedLocations   []AdditionalGeoLocation `json:"related_locations,omitempty"`
	ParkingType        *ParkingType            `json:"parking_type,omitempty"`
	Evses              []EVSE                  `json:"evses,omitempty"`
	Directions         []DisplayText           `json:"directions,omitempty"`
	Operator           *BusinessDetails        `json:"operator,omitempty"`
	Suboperator        *BusinessDetails        `json:"suboperator,omitempty"`
	Owner              *BusinessDetails        `json:"owner,omitempty"`
	Facilities         []Facility              `json:"facilities,omitempty"`
	TimeZone           string                  `json:"time_zone"`
	OpeningTimes       *Hours                  `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                   `json:"charging_when_closed,omitempty"`
	Images             []Image                 `json:"images,omitempty"`
	EnergyMix          *EnergyMix              `json:"energy_mix,omitempty"`
	LastUpdated        time.Time               `json:"last_updated"`
	Metadata           map[string]string       `json:"metadata,omitempty"`
}
