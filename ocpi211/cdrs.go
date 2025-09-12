package ocpi211

import (
	"time"
)

type CDR struct {
	ID               string            `json:"id"`
	StartDateTime    time.Time         `json:"start_date_time"`
	StopDateTime     time.Time         `json:"stop_date_time"`
	AuthID           string            `json:"auth_id"`
	AuthMethod       AuthMethod        `json:"auth_method"`
	Location         Location          `json:"location"`
	MeterID          *string           `json:"meter_id,omitempty"`
	Currency         string            `json:"currency"`
	Tariffs          []Tariff          `json:"tariffs,omitempty"`
	ChargingPeriods  []ChargingPeriod  `json:"charging_periods"`
	TotalCost        float64           `json:"total_cost"`
	TotalEnergy      float64           `json:"total_energy"`
	TotalTime        float64           `json:"total_time"`
	TotalParkingTime *float64          `json:"total_parking_time,omitempty"`
	Remark           *string           `json:"remark,omitempty"`
	LastUpdated      time.Time         `json:"last_updated"`
	Metadata         map[string]string `json:"metadata,omitempty"`
}
