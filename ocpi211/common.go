package ocpi211

import "time"

type ImageCategory string

const (
	ImageCategoryCharger  ImageCategory = "CHARGER"
	ImageCategoryEntrance ImageCategory = "ENTRANCE"
	ImageCategoryLocation ImageCategory = "LOCATION"
	ImageCategoryNetwork  ImageCategory = "NETWORK"
	ImageCategoryOperator ImageCategory = "OPERATOR"
	ImageCategoryOther    ImageCategory = "OTHER"
	ImageCategoryOwner    ImageCategory = "OWNER"
)

type EnergySourceCategory string

const (
	EnergySourceCategoryNuclear       EnergySourceCategory = "NUCLEAR"
	EnergySourceCategoryGeneralFossil EnergySourceCategory = "GENERAL_FOSSIL"
	EnergySourceCategoryCoal          EnergySourceCategory = "COAL"
	EnergySourceCategoryGas           EnergySourceCategory = "GAS"
	EnergySourceCategoryGeneralGreen  EnergySourceCategory = "GENERAL_GREEN"
	EnergySourceCategorySolar         EnergySourceCategory = "SOLAR"
	EnergySourceCategoryWind          EnergySourceCategory = "WIND"
	EnergySourceCategoryWater         EnergySourceCategory = "WATER"
)

type EnvironmentalImpactCategory string

const (
	EnvironmentalImpactCategoryNuclearWaste  EnvironmentalImpactCategory = "NUCLEAR_WASTE"
	EnvironmentalImpactCategoryCarbonDioxide EnvironmentalImpactCategory = "CARBON_DIOXIDE"
)

type DimensionType string

const (
	DimensionTypeEnergy      DimensionType = "ENERGY"
	DimensionTypeFlat        DimensionType = "FLAT"
	DimensionTypeMaxCurrent  DimensionType = "MAX_CURRENT"
	DimensionTypeMinCurrent  DimensionType = "MIN_CURRENT"
	DimensionTypeParkingTime DimensionType = "PARKING_TIME"
	DimensionTypeTime        DimensionType = "TIME"
)

type AuthMethod string

const (
	AuthMethodAuthRequest AuthMethod = "AUTH_REQUEST"
	AuthMethodWhitelist   AuthMethod = "WHITELIST"
)

type Dimension struct {
	Type   DimensionType `json:"type"`
	Volume float64       `json:"volume"`
}

type ChargingPeriod struct {
	StartDateTime time.Time   `json:"start_date_time"`
	Dimensions    []Dimension `json:"dimensions"`
}

type EnvironmentalImpact struct {
	Source     EnvironmentalImpactCategory `json:"source"`
	Percentage float64                     `json:"percentage"`
}

type EnergySource struct {
	Source     EnergySourceCategory `json:"source"`
	Percentage float64              `json:"percentage"`
}

type EnergyMix struct {
	IsGreenEnergy       bool                  `json:"is_green_energy"`
	EnergySources       []EnergySource        `json:"energy_sources,omitempty"`
	EnvironmentalImpact []EnvironmentalImpact `json:"environ_impact,omitempty"`
	SupplierName        *string               `json:"supplier_name,omitempty"`
	EnergyProductName   *string               `json:"energy_product_name,omitempty"`
}

type Image struct {
	Url       string        `json:"url"`
	Thumbnail *string       `json:"thumbnail,omitempty"`
	Category  ImageCategory `json:"category"`
	Type      string        `json:"type,omitempty"`
	Width     *int          `json:"width,omitempty"`
	Height    *int          `json:"height,omitempty"`
}

type BusinessDetails struct {
	Name    string  `json:"name"`
	Website *string `json:"website,omitempty"`
	Logo    *Image  `json:"logo,omitempty"`
}

type DisplayText struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}
