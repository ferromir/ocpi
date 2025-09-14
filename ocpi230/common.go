package ocpi230

import "time"

type Role string

const (
	RoleCPO   Role = "CPO"
	RoleEMSP  Role = "EMSP"
	RoleNAP   Role = "NAP"
	RoleNSP   Role = "NSP"
	RoleOTHER Role = "OTHER"
	RoleSCSP  Role = "SCSP"
)

type TokenType string

const (
	TokenTypeAdhoc   TokenType = "AD_HOC_USER"
	TokenTypeAppUser TokenType = "APP_USER"
	TokenTypeEMAID   TokenType = "EMAID"
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
	DimensionTypeCurrent         DimensionType = "CURRENT"
	DimensionTypeEnergy          DimensionType = "ENERGY"
	DimensionTypeEnergyExport    DimensionType = "ENERGY_EXPORT"
	DimensionTypeEnergyImport    DimensionType = "ENERGY_IMPORT"
	DimensionTypeMaxCurrent      DimensionType = "MAX_CURRENT"
	DimensionTypeMinCurrent      DimensionType = "MIN_CURRENT"
	DimensionTypeMaxPower        DimensionType = "MAX_POWER"
	DimensionTypeMinPower        DimensionType = "MIN_POWER"
	DimensionTypeParkingTime     DimensionType = "PARKING_TIME"
	DimensionTypePower           DimensionType = "POWER"
	DimensionTypeReservationTime DimensionType = "RESERVATION_TIME"
	DimensionTypeStateOfCharge   DimensionType = "STATE_OF_CHARGE"
	DimensionTypeTime            DimensionType = "TIME"
)

type AuthMethod string

const (
	AuthMethodAuthRequest AuthMethod = "AUTH_REQUEST"
	AuthMethodCommand     AuthMethod = "COMMAND"
	AuthMethodWhitelist   AuthMethod = "WHITELIST"
)

type ConnectorType string

const (
	ConnectorTypeCHADEMO               ConnectorType = "CHADEMO"
	ConnectorTypeCHAOJI                ConnectorType = "CHAOJI"
	ConnectorTypeDOMESTIC_A            ConnectorType = "DOMESTIC_A"
	ConnectorTypeDOMESTIC_B            ConnectorType = "DOMESTIC_B"
	ConnectorTypeDOMESTIC_C            ConnectorType = "DOMESTIC_C"
	ConnectorTypeDOMESTIC_D            ConnectorType = "DOMESTIC_D"
	ConnectorTypeDOMESTIC_E            ConnectorType = "DOMESTIC_E"
	ConnectorTypeDOMESTIC_F            ConnectorType = "DOMESTIC_F"
	ConnectorTypeDOMESTIC_G            ConnectorType = "DOMESTIC_G"
	ConnectorTypeDOMESTIC_H            ConnectorType = "DOMESTIC_H"
	ConnectorTypeDOMESTIC_I            ConnectorType = "DOMESTIC_I"
	ConnectorTypeDOMESTIC_J            ConnectorType = "DOMESTIC_J"
	ConnectorTypeDOMESTIC_K            ConnectorType = "DOMESTIC_K"
	ConnectorTypeDOMESTIC_L            ConnectorType = "DOMESTIC_L"
	ConnectorTypeDOMESTIC_M            ConnectorType = "DOMESTIC_M"
	ConnectorTypeDOMESTIC_N            ConnectorType = "DOMESTIC_N"
	ConnectorTypeDOMESTIC_O            ConnectorType = "DOMESTIC_O"
	ConnectorTypeGBT_AC                ConnectorType = "GBT_AC"
	ConnectorTypeGBT_DC                ConnectorType = "GBT_DC"
	ConnectorTypeIEC_60309_2_single_16 ConnectorType = "IEC_60309_2_single_16"
	ConnectorTypeIEC_60309_2_three_16  ConnectorType = "IEC_60309_2_three_16"
	ConnectorTypeIEC_60309_2_three_32  ConnectorType = "IEC_60309_2_three_32"
	ConnectorTypeIEC_60309_2_three_64  ConnectorType = "IEC_60309_2_three_64"
	ConnectorTypeIEC_62196_T1          ConnectorType = "IEC_62196_T1"
	ConnectorTypeIEC_62196_T1_COMBO    ConnectorType = "IEC_62196_T1_COMBO"
	ConnectorTypeIEC_62196_T2          ConnectorType = "IEC_62196_T2"
	ConnectorTypeIEC_62196_T2_COMBO    ConnectorType = "IEC_62196_T2_COMBO"
	ConnectorTypeIEC_62196_T3A         ConnectorType = "IEC_62196_T3A"
	ConnectorTypeIEC_62196_T3C         ConnectorType = "IEC_62196_T3C"
	ConnectorTypeIEC_MCS               ConnectorType = "MCS"
	ConnectorTypeNEMA_5_20             ConnectorType = "NEMA_5_20"
	ConnectorTypeNEMA_6_30             ConnectorType = "NEMA_6_30"
	ConnectorTypeNEMA_6_50             ConnectorType = "NEMA_6_50"
	ConnectorTypeNEMA_10_30            ConnectorType = "NEMA_10_30"
	ConnectorTypeNEMA_10_50            ConnectorType = "NEMA_10_50"
	ConnectorTypeNEMA_14_30            ConnectorType = "NEMA_14_30"
	ConnectorTypeNEMA_14_50            ConnectorType = "NEMA_14_50"
	ConnectorTypePANTOGRAPH_BOTTOM_UP  ConnectorType = "PANTOGRAPH_BOTTOM_UP"
	ConnectorTypePANTOGRAPH_TOP_DOWN   ConnectorType = "PANTOGRAPH_TOP_DOWN"
	ConnectorTypeSAE_J3400             ConnectorType = "SAE_J3400"
	ConnectorTypeTESLA_R               ConnectorType = "TESLA_R"
	ConnectorTypeTESLA_S               ConnectorType = "TESLA_S"
)

type ConnectorFormat string

const (
	ConnectorFormatSocket ConnectorFormat = "SOCKET"
	ConnectorFormatCable  ConnectorFormat = "CABLE"
)

type PowerType string

const (
	PowerTypeAC1Phase      PowerType = "AC_1_PHASE"
	PowerTypeAC2Phase      PowerType = "AC_2_PHASE"
	PowerTypeAC2PhaseSplit PowerType = "AC_2_PHASE_SPLIT"
	PowerTypeAC3Phase      PowerType = "AC_3_PHASE"
	PowerTypeDC            PowerType = "DC"
)

type TaxAmount struct {
	Name        string   `json:"name"`
	AccountName *string  `json:"account_name,omitempty"`
	Percentage  *float32 `json:"percentage,omitempty"`
	Amount      float32  `json:"amount"`
}

type Price struct {
	BeforeTaxes float64     `json:"before_taxes"`
	Taxes       []TaxAmount `json:"taxes,omitempty"`
}

type CDRToken struct {
	CountryCode string    `json:"country_code"`
	PartyID     string    `json:"party_id"`
	UID         string    `json:"uid"`
	Type        TokenType `json:"type"`
	ContractID  string    `json:"contract_id"`
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

type Dimension struct {
	Type   DimensionType `json:"type"`
	Volume float64       `json:"volume"`
}

type ChargingPeriod struct {
	StartDateTime time.Time   `json:"start_date_time"`
	Dimensions    []Dimension `json:"dimensions"`
	TariffId      *string     `json:"tariff_id"`
}

type EnvironmentalImpact struct {
	Category EnvironmentalImpactCategory `json:"category"`
	Amount   float64                     `json:"amount"`
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
