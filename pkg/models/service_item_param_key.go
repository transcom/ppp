package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"

	"github.com/gofrs/uuid"
)

//ServiceItemParamName is the name of service item parameter
type ServiceItemParamName string

func (s ServiceItemParamName) String() string {
	return string(s)
}

const (
	// ServiceItemParamNameCanStandAlone is the param key name CanStandAlone
	ServiceItemParamNameCanStandAlone ServiceItemParamName = "CanStandAlone"
	// ServiceItemParamNameCubicFeetBilled is the param key name CubicFeetBilled
	ServiceItemParamNameCubicFeetBilled ServiceItemParamName = "CubicFeetBilled"
	// ServiceItemParamNameCubicFeetCrating is the param key name CubicFeetCrating
	ServiceItemParamNameCubicFeetCrating ServiceItemParamName = "CubicFeetCrating"
	// ServiceItemParamNameDistanceZip3 is the param key name DistanceZip3
	ServiceItemParamNameDistanceZip3 ServiceItemParamName = "DistanceZip3"
	// ServiceItemParamNameDistanceZip5 is the param key name DistanceZip5
	ServiceItemParamNameDistanceZip5 ServiceItemParamName = "DistanceZip5"
	// ServiceItemParamNameDistanceZip5SITOrigin is the param key name DistanceZip5SITOrigin
	ServiceItemParamNameDistanceZip5SITOrigin ServiceItemParamName = "DistanceZip5SITOrigin"
	// ServiceItemParamNameDistanceZip5SITDest is the param key name DistanceZip5SITDest
	ServiceItemParamNameDistanceZip5SITDest ServiceItemParamName = "DistanceZip5SITDest"
	// ServiceItemParamNameEIAFuelPrice is the param key name EIAFuelPrice
	ServiceItemParamNameEIAFuelPrice ServiceItemParamName = "EIAFuelPrice"
	// ServiceItemParamNameMarketDest is the param key name MarketDest
	ServiceItemParamNameMarketDest ServiceItemParamName = "MarketDest"
	// ServiceItemParamNameMarketOrigin is the param key name MarketOrigin
	ServiceItemParamNameMarketOrigin ServiceItemParamName = "MarketOrigin"
	// ServiceItemParamNameNumberDaysSIT is the param key name NumberDaysSIT
	ServiceItemParamNameNumberDaysSIT ServiceItemParamName = "NumberDaysSIT"
	// ServiceItemParamNamePriceAreaDest is the param key name PriceAreaDest
	ServiceItemParamNamePriceAreaDest ServiceItemParamName = "PriceAreaDest"
	// ServiceItemParamNamePriceAreaIntlDest is the param key name PriceAreaIntlDest
	ServiceItemParamNamePriceAreaIntlDest ServiceItemParamName = "PriceAreaIntlDest"
	// ServiceItemParamNamePriceAreaIntlOrigin is the param key name PriceAreaIntlOrigin
	ServiceItemParamNamePriceAreaIntlOrigin ServiceItemParamName = "PriceAreaIntlOrigin"
	// ServiceItemParamNamePriceAreaOrigin is the param key name PriceAreaOrigin
	ServiceItemParamNamePriceAreaOrigin ServiceItemParamName = "PriceAreaOrigin"
	// ServiceItemParamNamePSILinehaulDom is the param key name PSI_LinehaulDom
	ServiceItemParamNamePSILinehaulDom ServiceItemParamName = "PSI_LinehaulDom"
	// ServiceItemParamNamePSILinehaulDomPrice is the param key name PSI_LinehaulDomPrice
	ServiceItemParamNamePSILinehaulDomPrice ServiceItemParamName = "PSI_LinehaulDomPrice"
	// ServiceItemParamNamePSILinehaulShort is the param key name PSI_LinehaulShort
	ServiceItemParamNamePSILinehaulShort ServiceItemParamName = "PSI_LinehaulShort"
	// ServiceItemParamNamePSILinehaulShortPrice is the param key name PSI_LinehaulShortPrice
	ServiceItemParamNamePSILinehaulShortPrice ServiceItemParamName = "PSI_LinehaulShortPrice"
	// ServiceItemParamNamePSIPackingDom is the param key name PSI_PackingDom
	ServiceItemParamNamePSIPackingDom ServiceItemParamName = "PSI_PackingDom"
	// ServiceItemParamNamePSIPackingDomPrice is the param key name PSI_PackingDomPrice
	ServiceItemParamNamePSIPackingDomPrice ServiceItemParamName = "PSI_PackingDomPrice"
	// ServiceItemParamNamePSIPackingHHGIntl is the param key name PSI_PackingHHGIntl
	ServiceItemParamNamePSIPackingHHGIntl ServiceItemParamName = "PSI_PackingHHGIntl"
	// ServiceItemParamNamePSIPackingHHGIntlPrice is the param key name PSI_PackingHHGIntlPrice
	ServiceItemParamNamePSIPackingHHGIntlPrice ServiceItemParamName = "PSI_PackingHHGIntlPrice"
	// ServiceItemParamNamePSIPriceDomDest is the param key name PSI_PriceDomDest
	ServiceItemParamNamePSIPriceDomDest ServiceItemParamName = "PSI_PriceDomDest"
	// ServiceItemParamNamePSIPriceDomDestPrice is the param key name PSI_PriceDomDestPrice
	ServiceItemParamNamePSIPriceDomDestPrice ServiceItemParamName = "PSI_PriceDomDestPrice"
	// ServiceItemParamNamePSIPriceDomOrigin is the param key name PSI_PriceDomOrigin
	ServiceItemParamNamePSIPriceDomOrigin ServiceItemParamName = "PSI_PriceDomOrigin"
	// ServiceItemParamNamePSIPriceDomOriginPrice is the param key name PSI_PriceDomOriginPrice
	ServiceItemParamNamePSIPriceDomOriginPrice ServiceItemParamName = "PSI_PriceDomOriginPrice"
	// ServiceItemParamNamePSIShippingLinehaulIntlCO is the param key name PSI_ShippingLinehaulIntlCO
	ServiceItemParamNamePSIShippingLinehaulIntlCO ServiceItemParamName = "PSI_ShippingLinehaulIntlCO"
	// ServiceItemParamNamePSIShippingLinehaulIntlCOPrice is the param key name PSI_ShippingLinehaulIntlCOPrice
	ServiceItemParamNamePSIShippingLinehaulIntlCOPrice ServiceItemParamName = "PSI_ShippingLinehaulIntlCOPrice"
	// ServiceItemParamNamePSIShippingLinehaulIntlOC is the param key name PSI_ShippingLinehaulIntlOC
	ServiceItemParamNamePSIShippingLinehaulIntlOC ServiceItemParamName = "PSI_ShippingLinehaulIntlOC"
	// ServiceItemParamNamePSIShippingLinehaulIntlOCPrice is the param key name PSI_ShippingLinehaulIntlOCPrice
	ServiceItemParamNamePSIShippingLinehaulIntlOCPrice ServiceItemParamName = "PSI_ShippingLinehaulIntlOCPrice"
	// ServiceItemParamNamePSIShippingLinehaulIntlOO is the param key name PSI_ShippingLinehaulIntlOO
	ServiceItemParamNamePSIShippingLinehaulIntlOO ServiceItemParamName = "PSI_ShippingLinehaulIntlOO"
	// ServiceItemParamNamePSIShippingLinehaulIntlOOPrice is the param key name PSI_ShippingLinehaulIntlOOPrice
	ServiceItemParamNamePSIShippingLinehaulIntlOOPrice ServiceItemParamName = "PSI_ShippingLinehaulIntlOOPrice"
	// ServiceItemParamNameRateAreaNonStdDest is the param key name RateAreaNonStdDest
	ServiceItemParamNameRateAreaNonStdDest ServiceItemParamName = "RateAreaNonStdDest"
	// ServiceItemParamNameRateAreaNonStdOrigin is the param key name RateAreaNonStdOrigin
	ServiceItemParamNameRateAreaNonStdOrigin ServiceItemParamName = "RateAreaNonStdOrigin"
	// ServiceItemParamNameRequestedPickupDate is the param key name RequestedPickupDate
	ServiceItemParamNameRequestedPickupDate ServiceItemParamName = "RequestedPickupDate"
	// ServiceItemParamNameServiceAreaDest is the param key name ServiceAreaDest
	ServiceItemParamNameServiceAreaDest ServiceItemParamName = "ServiceAreaDest"
	// ServiceItemParamNameServiceAreaOrigin is the param key name ServiceAreaOrigin
	ServiceItemParamNameServiceAreaOrigin ServiceItemParamName = "ServiceAreaOrigin"
	// ServiceItemParamNameServicesScheduleDest is the param key name ServicesScheduleDest
	ServiceItemParamNameServicesScheduleDest ServiceItemParamName = "ServicesScheduleDest"
	// ServiceItemParamNameServicesScheduleOrigin is the param key name ServicesScheduleOrigin
	ServiceItemParamNameServicesScheduleOrigin ServiceItemParamName = "ServicesScheduleOrigin"
	// ServiceItemParamNameSITScheduleDest is the param key name SITScheduleDest
	ServiceItemParamNameSITScheduleDest ServiceItemParamName = "SITScheduleDest"
	// ServiceItemParamNameSITScheduleOrigin is the param key name SITScheduleOrigin
	ServiceItemParamNameSITScheduleOrigin ServiceItemParamName = "SITScheduleOrigin"
	// ServiceItemParamNameWeightActual is the param key name WeightActual
	ServiceItemParamNameWeightActual ServiceItemParamName = "WeightActual"
	// ServiceItemParamNameWeightBilledActual is the param key name WeightBilledActual
	ServiceItemParamNameWeightBilledActual ServiceItemParamName = "WeightBilledActual"
	// ServiceItemParamNameWeightEstimated is the param key name WeightEstimated
	ServiceItemParamNameWeightEstimated ServiceItemParamName = "WeightEstimated"
	// ServiceItemParamNameZipDestAddress is the param key name ZipDestAddress
	ServiceItemParamNameZipDestAddress ServiceItemParamName = "ZipDestAddress"
	// ServiceItemParamNameZipPickupAddress is the param key name ZipPickupAddress
	ServiceItemParamNameZipPickupAddress ServiceItemParamName = "ZipPickupAddress"
	// ServiceItemParamNameZipSITAddress is the param key name ZipSITAddress
	ServiceItemParamNameZipSITAddress ServiceItemParamName = "ZipSITAddress"
)

// ServiceItemParamType is a type of service item parameter
type ServiceItemParamType string

// String is a string representation of a ServiceItemParamType
func (s ServiceItemParamType) String() string {
	return string(s)
}

const (
	// ServiceItemParamTypeString is a string
	ServiceItemParamTypeString ServiceItemParamType = "STRING"
	// ServiceItemParamTypeDate is a date
	ServiceItemParamTypeDate ServiceItemParamType = "DATE"
	// ServiceItemParamTypeInteger is an integer
	ServiceItemParamTypeInteger ServiceItemParamType = "INTEGER"
	// ServiceItemParamTypeDecimal is a decimal
	ServiceItemParamTypeDecimal ServiceItemParamType = "DECIMAL"
)

// ServiceItemParamOrigin is a type of service item parameter origin
type ServiceItemParamOrigin string

// String is a string representation of a ServiceItemParamOrigin
func (s ServiceItemParamOrigin) String() string {
	return string(s)
}

const (
	// ServiceItemParamOriginPrime is the Prime origin
	ServiceItemParamOriginPrime ServiceItemParamOrigin = "PRIME"
	// ServiceItemParamOriginSystem is the System origin
	ServiceItemParamOriginSystem ServiceItemParamOrigin = "SYSTEM"
)

// ValidServiceItemParamName lists all valid service item param key names
var ValidServiceItemParamName = []string{
	string(ServiceItemParamNameCanStandAlone),
	string(ServiceItemParamNameCubicFeetBilled),
	string(ServiceItemParamNameCubicFeetCrating),
	string(ServiceItemParamNameDistanceZip3),
	string(ServiceItemParamNameDistanceZip5),
	string(ServiceItemParamNameDistanceZip5SITDest),
	string(ServiceItemParamNameDistanceZip5SITOrigin),
	string(ServiceItemParamNameEIAFuelPrice),
	string(ServiceItemParamNameMarketDest),
	string(ServiceItemParamNameMarketOrigin),
	string(ServiceItemParamNameNumberDaysSIT),
	string(ServiceItemParamNamePriceAreaDest),
	string(ServiceItemParamNamePriceAreaIntlDest),
	string(ServiceItemParamNamePriceAreaIntlOrigin),
	string(ServiceItemParamNamePriceAreaOrigin),
	string(ServiceItemParamNamePSILinehaulDom),
	string(ServiceItemParamNamePSILinehaulDomPrice),
	string(ServiceItemParamNamePSILinehaulShort),
	string(ServiceItemParamNamePSILinehaulShortPrice),
	string(ServiceItemParamNamePSIPackingDom),
	string(ServiceItemParamNamePSIPackingDomPrice),
	string(ServiceItemParamNamePSIPackingHHGIntl),
	string(ServiceItemParamNamePSIPackingHHGIntlPrice),
	string(ServiceItemParamNamePSIPriceDomDest),
	string(ServiceItemParamNamePSIPriceDomDestPrice),
	string(ServiceItemParamNamePSIPriceDomOrigin),
	string(ServiceItemParamNamePSIPriceDomOriginPrice),
	string(ServiceItemParamNamePSIShippingLinehaulIntlCO),
	string(ServiceItemParamNamePSIShippingLinehaulIntlCOPrice),
	string(ServiceItemParamNamePSIShippingLinehaulIntlOC),
	string(ServiceItemParamNamePSIShippingLinehaulIntlOCPrice),
	string(ServiceItemParamNamePSIShippingLinehaulIntlOO),
	string(ServiceItemParamNamePSIShippingLinehaulIntlOOPrice),
	string(ServiceItemParamNameRateAreaNonStdDest),
	string(ServiceItemParamNameRateAreaNonStdOrigin),
	string(ServiceItemParamNameRequestedPickupDate),
	string(ServiceItemParamNameServiceAreaDest),
	string(ServiceItemParamNameServiceAreaOrigin),
	string(ServiceItemParamNameServicesScheduleDest),
	string(ServiceItemParamNameServicesScheduleOrigin),
	string(ServiceItemParamNameSITScheduleDest),
	string(ServiceItemParamNameSITScheduleOrigin),
	string(ServiceItemParamNameWeightActual),
	string(ServiceItemParamNameWeightBilledActual),
	string(ServiceItemParamNameWeightEstimated),
	string(ServiceItemParamNameZipDestAddress),
	string(ServiceItemParamNameZipPickupAddress),
	string(ServiceItemParamNameZipSITAddress),
}

var validServiceItemParamType = []string{
	string(ServiceItemParamTypeString),
	string(ServiceItemParamTypeDate),
	string(ServiceItemParamTypeInteger),
	string(ServiceItemParamTypeDecimal),
}

var validServiceItemParamOrigin = []string{
	string(ServiceItemParamOriginPrime),
	string(ServiceItemParamOriginSystem),
}

// ServiceItemParamKey is a key for a Service Item Param
type ServiceItemParamKey struct {
	ID          uuid.UUID              `json:"id" db:"id"`
	Key         ServiceItemParamName   `json:"key" db:"key"`
	Description string                 `json:"description" db:"description"`
	Type        ServiceItemParamType   `json:"type" db:"type"`
	Origin      ServiceItemParamOrigin `json:"origin" db:"origin"`
	CreatedAt   time.Time              `db:"created_at"`
	UpdatedAt   time.Time              `db:"updated_at"`
}

// ServiceItemParamKeys is not required by pop and may be deleted
type ServiceItemParamKeys []ServiceItemParamKey

// Validate validates a ServiceItemParamKey
func (s *ServiceItemParamKey) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Key.String(), Name: "Key"},
		&validators.StringIsPresent{Field: s.Description, Name: "Description"},
		&validators.StringIsPresent{Field: string(s.Type), Name: "Type"},
		&validators.StringIsPresent{Field: string(s.Origin), Name: "Origin"},
		&validators.StringInclusion{Field: s.Type.String(), Name: "Type", List: validServiceItemParamType},
		&validators.StringInclusion{Field: s.Origin.String(), Name: "Origin", List: validServiceItemParamOrigin},
	), nil
}
