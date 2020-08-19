package ghcrateengine

import (
	"time"

	"github.com/gobuffalo/pop"

	"github.com/transcom/mymove/pkg/models"
)

func fetchTaskOrderFee(db *pop.Connection, contractCode string, serviceCode models.ReServiceCode, mtoAvailableToPrimeAt time.Time) (models.ReTaskOrderFee, error) {
	var taskOrderFee models.ReTaskOrderFee
	err := db.Q().
		Join("re_contract_years cy", "re_task_order_fees.contract_year_id = cy.id").
		Join("re_contracts c", "cy.contract_id = c.id").
		Join("re_services s", "re_task_order_fees.service_id = s.id").
		Where("c.code = $1", contractCode).
		Where("s.code = $2", serviceCode).
		Where("$3 between cy.start_date and cy.end_date", mtoAvailableToPrimeAt).
		First(&taskOrderFee)

	if err != nil {
		return models.ReTaskOrderFee{}, err
	}

	return taskOrderFee, nil
}

func fetchDomOtherPrice(db *pop.Connection, contractCode string, serviceCode models.ReServiceCode, servicesSchedule int, isPeakPeriod bool) (models.ReDomesticOtherPrice, error) {
	var domOtherPrice models.ReDomesticOtherPrice
	err := db.Q().
		Join("re_services", "service_id = re_services.id").
		Join("re_contracts", "re_contracts.id = re_domestic_other_prices.contract_id").
		Where("re_contracts.code = $1", contractCode).
		Where("re_services.code = $2", serviceCode).
		Where("schedule = $3", servicesSchedule).
		Where("is_peak_period = $4", isPeakPeriod).
		First(&domOtherPrice)

	if err != nil {
		return models.ReDomesticOtherPrice{}, err
	}

	return domOtherPrice, nil
}
