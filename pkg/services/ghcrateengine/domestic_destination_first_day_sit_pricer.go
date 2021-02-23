package ghcrateengine

import (
	"time"

	"github.com/gobuffalo/pop/v5"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/unit"
)

type domesticDestinationFirstDaySITPricer struct {
	db *pop.Connection
}

// NewDomesticDestinationFirstDaySITPricer creates a new pricer for domestic destination first day SIT
func NewDomesticDestinationFirstDaySITPricer(db *pop.Connection) services.DomesticDestinationFirstDaySITPricer {
	return &domesticDestinationFirstDaySITPricer{
		db: db,
	}
}

// Price determines the price for domestic destination first day SIT
func (p domesticDestinationFirstDaySITPricer) Price(contractCode string, requestedPickupDate time.Time, weight unit.Pound, serviceArea string) (unit.Cents, []services.PricingParam, error) {
	return priceDomesticFirstDaySIT(p.db, models.ReServiceCodeDDFSIT, contractCode, requestedPickupDate, weight, serviceArea)
}

// PriceUsingParams determines the price for domestic destination first day SIT given PaymentServiceItemParams
func (p domesticDestinationFirstDaySITPricer) PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, []services.PricingParam, error) {
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

	serviceAreaDest, err := getParamString(params, models.ServiceItemParamNameServiceAreaDest)
	if err != nil {
		return unit.Cents(0), nil, err
	}

	return p.Price(contractCode, requestedPickupDate, unit.Pound(weightBilledActual), serviceAreaDest)
}
