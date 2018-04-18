package models_test

import (
	"time"

	. "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/unit"
)

func (suite *ModelSuite) Test_EffectiveDateValidation() {
	now := time.Now()

	validPackRate := Tariff400ngFullPackRate{
		WeightLbsLower:     100,
		WeightLbsUpper:     200,
		EffectiveDateLower: now,
		EffectiveDateUpper: now.AddDate(1, 0, 0),
	}

	expErrors := map[string][]string{}
	suite.verifyValidationErrors(&validPackRate, expErrors)

	invalidPackRate := Tariff400ngFullPackRate{
		WeightLbsLower:     100,
		WeightLbsUpper:     200,
		EffectiveDateLower: now,
		EffectiveDateUpper: now.AddDate(-1, 0, 0),
	}

	expErrors = map[string][]string{
		"effective_date_upper": []string{"EffectiveDateUpper must be after EffectiveDateLower."},
	}
	suite.verifyValidationErrors(&invalidPackRate, expErrors)
}

func (suite *ModelSuite) Test_WeightValidation() {
	validPackRate := Tariff400ngFullPackRate{
		WeightLbsLower: 100,
		WeightLbsUpper: 200,
	}

	expErrors := map[string][]string{}
	suite.verifyValidationErrors(&validPackRate, expErrors)

	invalidPackRate := Tariff400ngFullPackRate{
		WeightLbsLower: 200,
		WeightLbsUpper: 100,
	}

	expErrors = map[string][]string{
		"weight_lbs_lower": []string{"200 is not less than 100."},
	}
	suite.verifyValidationErrors(&invalidPackRate, expErrors)
}

func (suite *ModelSuite) Test_RateValidation() {
	validPackRate := Tariff400ngFullPackRate{
		RateCents:      100,
		WeightLbsLower: 100,
		WeightLbsUpper: 200,
	}

	expErrors := map[string][]string{}
	suite.verifyValidationErrors(&validPackRate, expErrors)

	invalidPackRate := Tariff400ngFullPackRate{
		RateCents:      -1,
		WeightLbsLower: 100,
		WeightLbsUpper: 200,
	}

	expErrors = map[string][]string{
		"rate_cents": []string{"-1 is not greater than -1."},
	}
	suite.verifyValidationErrors(&invalidPackRate, expErrors)
}

func (suite *ModelSuite) Test_FetchFullPackRateCents() {
	t := suite.T()

	rate1 := unit.Cents(100)
	cwt := 15
	schedule := 1

	fpr1 := Tariff400ngFullPackRate{
		Schedule:           schedule,
		WeightLbsLower:     1000,
		WeightLbsUpper:     2000,
		RateCents:          rate1,
		EffectiveDateLower: testdatagen.PeakRateCycleStart,
		EffectiveDateUpper: testdatagen.PeakRateCycleEnd,
	}
	suite.mustSave(&fpr1)

	rate, err := FetchTariff400ngFullPackRateCents(suite.db, cwt, schedule)
	if err != nil {
		t.Fatalf("Unable to query full pack rate: %v", err)
	}
	if rate != rate1 {
		t.Errorf("Incorrect full pack rate recieved. Got: %d. Expected: %d.", rate, rate1)
	}

	// Test upper bound exclusivity of EffectiveDateUpper
	rate, err = FetchTariff400ngFullPackRateCents(suite.db, cwt, schedule)
	if err == nil && rate == rate1 {
		t.Errorf("EffectiveDateUpper is inclusive of upper bound.")
	}
}
