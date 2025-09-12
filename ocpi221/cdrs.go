package ocpi221

import "time"

type CDRLocation struct {
	ID                 string          `json:"id"`
	Name               *string         `json:"name,omitempty"`
	Address            string          `json:"address"`
	City               string          `json:"city"`
	PostalCode         *string         `json:"postal_code,omitempty"`
	State              *string         `json:"state,omitempty"`
	Country            string          `json:"country"`
	Coordinates        GeoLocation     `json:"coordinates"`
	EvseUID            string          `json:"evse_uid"`
	EvseID             string          `json:"evse_id"`
	ConnectorID        string          `json:"connector_id"`
	ConnectorStandard  ConnectorType   `json:"connector_standard"`
	ConnectorFormat    ConnectorFormat `json:"connector_format"`
	ConnectorPowerType PowerType       `json:"connector_power_type"`
}

type SignedValue struct {
	Nature     string `json:"nature"`
	PlainData  string `json:"plain_data"`
	SignedData string `json:"signed_data"`
}

type SignedData struct {
	EncodingMethod        string        `json:"encoding_method"`
	EncodingMethodVersion *int          `json:"encoding_method_version,omitempty"`
	PublicKey             *string       `json:"public_key,omitempty"`
	SignedValues          []SignedValue `json:"signed_values"`
	Url                   *string       `json:"url,omitempty"`
}

type CDR struct {
	CountryCode            string      `json:"country_code"`
	PartyID                string      `json:"party_id"`
	ID                     string      `json:"id"`
	StartDateTime          time.Time   `json:"start_date_time"`
	StopDateTime           time.Time   `json:"stop_date_time"`
	SessionID              *string     `json:"session_id,omitempty"`
	CDRToken               CDRToken    `json:"cdr_token"`
	AuthMethod             AuthMethod  `json:"auth_method"`
	AuthorizationReference *string     `json:"authorization_reference,omitempty"`
	CDRLocation            CDRLocation `json:"cdr_location"`
	MeterID                *string     `json:"meter_id,omitempty"`
	Currency               string      `json:"currency"`
	// Tariffs                  []Tariff          `json:"tariffs,omitempty"`
	ChargingPeriods          []ChargingPeriod  `json:"charging_periods"`
	SignedData               *SignedData       `json:"signed_data,omitempty"`
	TotalCost                Price             `json:"total_cost"`
	TotalFixedCost           *Price            `json:"total_fixed_cost,omitempty"`
	TotalEnergy              float64           `json:"total_energy"`
	TotalEnergyCost          *Price            `json:"total_energy_cost,omitempty"`
	TotalTime                float64           `json:"total_time"`
	TotalTimeCost            *Price            `json:"total_time_cost,omitempty"`
	TotalParkingTime         float64           `json:"total_parking_time"`
	TotalParkingCost         *Price            `json:"total_parking_cost,omitempty"`
	TotalReservationCost     *Price            `json:"total_reservation_cost,omitempty"`
	Remark                   *string           `json:"remark,omitempty"`
	InvoiceReferenceId       *string           `json:"invoice_reference_id,omitempty"`
	Credit                   *bool             `json:"credit,omitempty"`
	CreditReferenceId        *string           `json:"credit_reference_id,omitempty"`
	HomeChargingCompensation *bool             `json:"home_charging_compensation,omitempty"`
	LastUpdated              time.Time         `json:"last_updated"`
	Metadata                 map[string]string `json:"metadata,omitempty"`
}
