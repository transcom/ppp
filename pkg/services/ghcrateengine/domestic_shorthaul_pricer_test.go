package ghcrateengine

import (
	"strconv"
	"testing"
	"time"

	"github.com/transcom/mymove/pkg/unit"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

const (
	dshTestServiceArea = "006"
	dshTestWeight      = 3600
	dshTestMileage     = 1200
)

func (suite *GHCRateEngineServiceSuite) TestPriceDomesticShorthaulWithServiceItemParamsBadData() {
	suite.setUpDomesticShorthaulData()
	paymentServiceItem := suite.setupPaymentServiceItemWithParams(
		models.ReServiceCodeDSH,
		[]createParams{
			{
				models.ServiceItemParamNameContractCode,
				models.ServiceItemParamTypeString,
				testdatagen.DefaultContractCode,
			},
			{
				models.ServiceItemParamNameRequestedPickupDate,
				models.ServiceItemParamTypeDate,
				time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC).Format(DateParamFormat),
			},
			{
				models.ServiceItemParamNameDistanceZip5,
				models.ServiceItemParamTypeInteger,
				strconv.Itoa(dshTestMileage),
			},
			{
				models.ServiceItemParamNameWeightBilledActual,
				models.ServiceItemParamTypeInteger,
				"0",
			},
			{
				models.ServiceItemParamNameServiceAreaOrigin,
				models.ServiceItemParamTypeString,
				dshTestServiceArea,
			},
		},
	)

	pricer := NewDomesticShorthaulPricer(suite.DB())

	suite.T().Run("failure during pricing bubbles up", func(t *testing.T) {
		_, err := pricer.PriceUsingParams(paymentServiceItem.PaymentServiceItemParams)
		suite.Error(err)
		suite.Equal("Weight must be greater than 0", err.Error())
	})
}

func (suite *GHCRateEngineServiceSuite) TestPriceDomesticShorthaulWithServiceItemParams() {
	suite.setUpDomesticShorthaulData()
	paymentServiceItem := suite.setupDomesticShorthaulServiceItems()

	pricer := NewDomesticShorthaulPricer(suite.DB())

	suite.T().Run("success all params for shorthaul available", func(t *testing.T) {
		cost, err := pricer.PriceUsingParams(paymentServiceItem.PaymentServiceItemParams)
		expectedCost := unit.Cents(6563903)
		suite.NoError(err)
		suite.Equal(expectedCost, cost)
	})

	suite.T().Run("validation errors", func(t *testing.T) {
		// No contract code
		_, err := pricer.PriceUsingParams(models.PaymentServiceItemParams{})
		suite.Error(err)
		suite.Equal("could not find param with key ContractCode", err.Error())

		// No requested pickup date
		missingRequestedPickupDate := suite.removeOnePaymentServiceItem(paymentServiceItem.PaymentServiceItemParams, models.ServiceItemParamNameRequestedPickupDate)
		_, err = pricer.PriceUsingParams(missingRequestedPickupDate)
		suite.Error(err)
		suite.Equal("could not find param with key RequestedPickupDate", err.Error())

		// No distance
		missingDistanceZip5 := suite.removeOnePaymentServiceItem(paymentServiceItem.PaymentServiceItemParams, models.ServiceItemParamNameDistanceZip5)
		_, err = pricer.PriceUsingParams(missingDistanceZip5)
		suite.Error(err)
		suite.Equal("could not find param with key DistanceZip5", err.Error())

		// No weight
		missingBilledActualWeight := suite.removeOnePaymentServiceItem(paymentServiceItem.PaymentServiceItemParams, models.ServiceItemParamNameWeightBilledActual)
		_, err = pricer.PriceUsingParams(missingBilledActualWeight)
		suite.Error(err)
		suite.Equal("could not find param with key WeightBilledActual", err.Error())

		// No service area
		missingServiceAreaOrigin := suite.removeOnePaymentServiceItem(paymentServiceItem.PaymentServiceItemParams, models.ServiceItemParamNameServiceAreaOrigin)
		_, err = pricer.PriceUsingParams(missingServiceAreaOrigin)
		suite.Error(err)
		suite.Equal("could not find param with key ServiceAreaOrigin", err.Error())
	})
}

func (suite *GHCRateEngineServiceSuite) TestPriceDomesticShorthaul() {
	suite.setUpDomesticShorthaulData()

	pricer := NewDomesticShorthaulPricer(suite.DB())

	suite.T().Run("success shorthaul cost within peak period", func(t *testing.T) {
		cost, err := pricer.Price(
			testdatagen.DefaultContractCode,
			time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
			dshTestMileage,
			dshTestWeight,
			dshTestServiceArea,
		)
		expectedCost := unit.Cents(6563903)
		suite.NoError(err)
		suite.Equal(expectedCost, cost)
	})

	suite.T().Run("success shorthaul cost within non-peak period", func(t *testing.T) {
		nonPeakDate := peakStart.addDate(0, -1)
		cost, err := pricer.Price(
			testdatagen.DefaultContractCode,
			time.Date(testdatagen.TestYear, nonPeakDate.month, nonPeakDate.day, 0, 0, 0, 0, time.UTC),
			dshTestMileage,
			dshTestWeight,
			dshTestServiceArea,
		)
		expectedCost := unit.Cents(5709696)
		suite.NoError(err)
		suite.Equal(expectedCost, cost)
	})

	suite.T().Run("failure if contract code bogus", func(t *testing.T) {
		_, err := pricer.Price(
			"bogus_code",
			time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
			dshTestMileage,
			dshTestWeight,
			dshTestServiceArea,
		)

		suite.Error(err)
		suite.Equal("Could not lookup Domestic Service Area Price: sql: no rows in result set", err.Error())
	})

	suite.T().Run("failure if move date is outside of contract year", func(t *testing.T) {
		_, err := pricer.Price(
			testdatagen.DefaultContractCode,
			time.Date(testdatagen.TestYear+1, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
			dshTestMileage,
			dshTestWeight,
			dshTestServiceArea,
		)

		suite.Error(err)
		suite.Equal("Could not lookup contract year: sql: no rows in result set", err.Error())
	})

	suite.T().Run("weight below minimum", func(t *testing.T) {
		cost, err := pricer.Price(
			testdatagen.DefaultContractCode,
			time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
			dshTestMileage,
			unit.Pound(499),
			dshTestServiceArea,
		)
		expectedCost := unit.Cents(911653)
		suite.NoError(err)
		suite.Equal(expectedCost, cost)
	})

	suite.T().Run("validation errors", func(t *testing.T) {
		requestedPickupDate := time.Date(testdatagen.TestYear, time.July, 4, 0, 0, 0, 0, time.UTC)

		// No contract code
		_, err := pricer.Price("", requestedPickupDate, dshTestMileage, dshTestWeight, dshTestServiceArea)
		suite.Error(err)
		suite.Equal("ContractCode is required", err.Error())

		// No requested pickup date
		_, err = pricer.Price(testdatagen.DefaultContractCode, time.Time{}, dshTestMileage, dshTestWeight, dshTestServiceArea)
		suite.Error(err)
		suite.Equal("RequestedPickupDate is required", err.Error())

		// No distance
		_, err = pricer.Price(testdatagen.DefaultContractCode, requestedPickupDate, 0, dshTestWeight, dshTestServiceArea)
		suite.Error(err)
		suite.Equal("Distance must be greater than 0", err.Error())

		// No weight
		_, err = pricer.Price(testdatagen.DefaultContractCode, requestedPickupDate, dshTestMileage, 0, dshTestServiceArea)
		suite.Error(err)
		suite.Equal("Weight must be greater than 0", err.Error())

		// No service area
		_, err = pricer.Price(testdatagen.DefaultContractCode, requestedPickupDate, dshTestMileage, dshTestWeight, "")
		suite.Error(err)
		suite.Equal("ServiceArea is required", err.Error())
	})
}

func (suite *GHCRateEngineServiceSuite) setupDomesticShorthaulServiceItems() models.PaymentServiceItem {
	return suite.setupPaymentServiceItemWithParams(
		models.ReServiceCodeDSH,
		[]createParams{
			{
				models.ServiceItemParamNameContractCode,
				models.ServiceItemParamTypeString,
				testdatagen.DefaultContractCode,
			},
			{
				models.ServiceItemParamNameRequestedPickupDate,
				models.ServiceItemParamTypeDate,
				time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC).Format(DateParamFormat),
			},
			{
				models.ServiceItemParamNameDistanceZip5,
				models.ServiceItemParamTypeInteger,
				strconv.Itoa(dshTestMileage),
			},
			{
				models.ServiceItemParamNameWeightBilledActual,
				models.ServiceItemParamTypeInteger,
				strconv.Itoa(dshTestWeight),
			},
			{
				models.ServiceItemParamNameServiceAreaOrigin,
				models.ServiceItemParamTypeString,
				dshTestServiceArea,
			},
		},
	)
}

func (suite *GHCRateEngineServiceSuite) removeOnePaymentServiceItem(paymentServiceItemParams models.PaymentServiceItemParams, nameToRemove models.ServiceItemParamName) models.PaymentServiceItemParams {
	var params models.PaymentServiceItemParams
	for _, param := range paymentServiceItemParams {
		if param.ServiceItemParamKey.Key != nameToRemove {
			params = append(params, param)
		}
	}
	return params
}

func (suite *GHCRateEngineServiceSuite) setUpDomesticShorthaulData() {
	contractYear := testdatagen.MakeReContractYear(suite.DB(),
		testdatagen.Assertions{
			ReContractYear: models.ReContractYear{
				Escalation:           1.0197,
				EscalationCompounded: 1.0407,
			},
		})

	serviceArea := testdatagen.MakeReDomesticServiceArea(suite.DB(),
		testdatagen.Assertions{
			ReDomesticServiceArea: models.ReDomesticServiceArea{
				Contract:    contractYear.Contract,
				ServiceArea: dshTestServiceArea,
			},
		})

	domesticShorthaulService := testdatagen.MakeReService(suite.DB(),
		testdatagen.Assertions{
			ReService: models.ReService{
				Code: "DSH",
				Name: "Dom. Shorthaul",
			},
		})

	domesticShorthaulPrice := models.ReDomesticServiceAreaPrice{
		ContractID:            contractYear.Contract.ID,
		DomesticServiceAreaID: serviceArea.ID,
		IsPeakPeriod:          true,
		ServiceID:             domesticShorthaulService.ID,
	}

	domesticShorthaulPeakPrice := domesticShorthaulPrice
	domesticShorthaulPeakPrice.PriceCents = 146
	suite.MustSave(&domesticShorthaulPeakPrice)

	domesticShorthaulNonpeakPrice := domesticShorthaulPrice
	domesticShorthaulNonpeakPrice.IsPeakPeriod = false
	domesticShorthaulNonpeakPrice.PriceCents = 127
	suite.MustSave(&domesticShorthaulNonpeakPrice)
}
