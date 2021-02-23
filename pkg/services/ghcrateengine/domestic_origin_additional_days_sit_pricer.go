package ghcrateengine

import (
	"time"

	"github.com/gobuffalo/pop/v5"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/unit"
)

type domesticOriginAdditionalDaysSITPricer struct {
	db *pop.Connection
}

// NewDomesticOriginAdditionalDaysSITPricer creates a new pricer for domestic origin additional days SIT
func NewDomesticOriginAdditionalDaysSITPricer(db *pop.Connection) services.DomesticOriginAdditionalDaysSITPricer {
	return &domesticOriginAdditionalDaysSITPricer{
		db: db,
	}
}

// Price determines the price for domestic origin additional days SIT
func (p domesticOriginAdditionalDaysSITPricer) Price(contractCode string, requestedPickupDate time.Time, isPeakPeriod bool, weight unit.Pound, serviceArea string, numberOfDaysInSIT int) (unit.Cents, []services.PricingParam, error) {
	return priceDomesticAdditionalDaysSIT(p.db, models.ReServiceCodeDOASIT, contractCode, requestedPickupDate, isPeakPeriod, weight, serviceArea, numberOfDaysInSIT)
}

// PriceUsingParams determines the price for domestic origin first day SIT given PaymentServiceItemParams
func (p domesticOriginAdditionalDaysSITPricer) PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, []services.PricingParam, error) {
	contractCode, err := getParamString(params, models.ServiceItemParamNameContractCode)
	if err != nil {
		return unit.Cents(0), nil, err
	}

	requestedPickupDate, err := getParamTime(params, models.ServiceItemParamNameRequestedPickupDate)
	if err != nil {
		return unit.Cents(0), nil, err
	}

	weightBilledActual, err := getParamInt(params, models.ServiceItemParamNameWeightBilledActual)
	if err != nil {
		return unit.Cents(0), nil, err
	}

	serviceAreaOrigin, err := getParamString(params, models.ServiceItemParamNameServiceAreaOrigin)
	if err != nil {
		return unit.Cents(0), nil, err
	}

	numberOfDaysInSIT, err := getParamInt(params, models.ServiceItemParamNameNumberDaysSIT)
	if err != nil {
		return unit.Cents(0), nil, err
	}

	isPeakPeriod := IsPeakPeriod(requestedPickupDate)

	return p.Price(contractCode, requestedPickupDate, isPeakPeriod, unit.Pound(weightBilledActual), serviceAreaOrigin, numberOfDaysInSIT)
}
