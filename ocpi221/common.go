package ocpi221

import "time"

type TokenType string

const (
	TokenTypeAdhoc   TokenType = "AD_HOC_USER"
	TokenTypeAppUser TokenType = "APP_USER"
	TokenTypeOther   TokenType = "OTHER"
	TokenTypeRFID    TokenType = "RFID"
)

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
	AuthMethodCommand     AuthMethod = "COMMAND"
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

type GeoLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type ConnectorType string

const (
	ConnectorTypeCHAdeMO          ConnectorType = "CHADEMO"
	ConnectorTypeIEC62196Type1    ConnectorType = "IEC_62196_T1"
	ConnectorTypeIEC62196Type1CCS ConnectorType = "IEC_62196_T1_COMBO"
	ConnectorTypeIEC62196Type2    ConnectorType = "IEC_62196_T2"
	ConnectorTypeIEC62196Type2CCS ConnectorType = "IEC_62196_T2_COMBO"
	ConnectorTypeIEC62196Type3    ConnectorType = "IEC_62196_T3A"
	ConnectorTypeDomesticA        ConnectorType = "DOMESTIC_A"
	ConnectorTypeDomesticB        ConnectorType = "DOMESTIC_B"
	ConnectorTypeDomesticC        ConnectorType = "DOMESTIC_C"
	ConnectorTypeDomesticD        ConnectorType = "DOMESTIC_D"
	ConnectorTypeDomesticE        ConnectorType = "DOMESTIC_E"
	ConnectorTypeDomesticF        ConnectorType = "DOMESTIC_F"
	ConnectorTypeDomesticG        ConnectorType = "DOMESTIC_G"
	ConnectorTypeDomesticH        ConnectorType = "DOMESTIC_H"
	ConnectorTypeDomesticI        ConnectorType = "DOMESTIC_I"
	ConnectorTypeDomesticJ        ConnectorType = "DOMESTIC_J"
	ConnectorTypeDomesticK        ConnectorType = "DOMESTIC_K"
	ConnectorTypeDomesticL        ConnectorType = "DOMESTIC_L"
	ConnectorTypeTeslaR           ConnectorType = "TESLA_R"
	ConnectorTypeTeslaS           ConnectorType = "TESLA_S"
)

type ConnectorFormat string

const (
	ConnectorFormatSocket ConnectorFormat = "SOCKET"
	ConnectorFormatCable  ConnectorFormat = "CABLE"
)

type PowerType string

const (
	PowerTypeAC1Phase PowerType = "AC_1_PHASE"
	PowerTypeAC3Phase PowerType = "AC_3_PHASE"
	PowerTypeDC       PowerType = "DC"
)

type Price struct {
	ExclVat float64 `json:"excl_vat"`
	InclVat float64 `json:"incl_vat"`
}
