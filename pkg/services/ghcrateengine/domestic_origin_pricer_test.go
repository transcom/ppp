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
	dopTestServiceArea = "006"
	dopTestWeight      = 3600
)

func (suite *GHCRateEngineServiceSuite) TestPriceDomesticOriginWithServiceItemParamsBadData() {
	suite.setUpDomesticOriginData()
	paymentServiceItem := suite.setupPaymentServiceItemWithParams(
		models.ReServiceCodeDOP,
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
				models.ServiceItemParamNameWeightBilledActual,
				models.ServiceItemParamTypeInteger,
				"0",
			},
			{
				models.ServiceItemParamNameServiceAreaOrigin,
				models.ServiceItemParamTypeString,
				dopTestServiceArea,
			},
		},
	)

	pricer := NewDomesticOriginPricer(suite.DB())

	suite.T().Run("failure during pricing bubbles up", func(t *testing.T) {
		_, err := pricer.PriceUsingParams(paymentServiceItem.PaymentServiceItemParams)
		suite.Error(err)
		suite.Equal("Weight must be a minimum of 500", err.Error())
	})
}

func (suite *GHCRateEngineServiceSuite) TestPriceDomesticOriginWithServiceItemParams() {
	suite.setUpDomesticOriginData()
	paymentServiceItem := suite.setupDomesticOriginServiceItems()

	pricer := NewDomesticOriginPricer(suite.DB())

	suite.T().Run("success all params for domestic origin available", func(t *testing.T) {
		cost, err := pricer.PriceUsingParams(paymentServiceItem.PaymentServiceItemParams)
		expectedCost := unit.Cents(5470)
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

func (suite *GHCRateEngineServiceSuite) TestPriceDomesticOrigin() {
	suite.setUpDomesticOriginData()

	pricer := NewDomesticOriginPricer(suite.DB())

	suite.T().Run("success domestic origin cost within peak period", func(t *testing.T) {
		cost, err := pricer.Price(
			testdatagen.DefaultContractCode,
			time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
			dopTestWeight,
			dopTestServiceArea,
		)
		expectedCost := unit.Cents(5470)
		suite.NoError(err)
		suite.Equal(expectedCost, cost)
	})

	suite.T().Run("success domestic origin cost within non-peak period", func(t *testing.T) {
		nonPeakDate := peakStart.addDate(0, -1)
		cost, err := pricer.Price(
			testdatagen.DefaultContractCode,
			time.Date(testdatagen.TestYear, nonPeakDate.month, nonPeakDate.day, 0, 0, 0, 0, time.UTC),
			dopTestWeight,
			dopTestServiceArea,
		)
		expectedCost := unit.Cents(4758)
		suite.NoError(err)
		suite.Equal(expectedCost, cost)
	})

	suite.T().Run("failure if contract code bogus", func(t *testing.T) {
		_, err := pricer.Price(
			"bogus_code",
			time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
			dopTestWeight,
			dopTestServiceArea,
		)

		suite.Error(err)
		suite.Equal("Could not lookup Domestic Service Area Price: sql: no rows in result set", err.Error())
	})

	suite.T().Run("failure if move date is outside of contract year", func(t *testing.T) {
		_, err := pricer.Price(
			testdatagen.DefaultContractCode,
			time.Date(testdatagen.TestYear+1, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
			dopTestWeight,
			dopTestServiceArea,
		)

		suite.Error(err)
		suite.Equal("Could not lookup contract year: sql: no rows in result set", err.Error())
	})

	suite.T().Run("weight below minimum", func(t *testing.T) {
		cost, err := pricer.Price(
			testdatagen.DefaultContractCode,
			time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
			unit.Pound(499),
			dopTestServiceArea,
		)
		suite.Equal(unit.Cents(0), cost)
		suite.Error(err)
		suite.Equal("Weight must be a minimum of 500", err.Error())
	})

	suite.T().Run("validation errors", func(t *testing.T) {
		requestedPickupDate := time.Date(testdatagen.TestYear, time.July, 4, 0, 0, 0, 0, time.UTC)

		// No contract code
		_, err := pricer.Price("", requestedPickupDate, dshTestWeight, dopTestServiceArea)
		suite.Error(err)
		suite.Equal("ContractCode is required", err.Error())

		// No requested pickup date
		_, err = pricer.Price(testdatagen.DefaultContractCode, time.Time{}, dshTestWeight, dopTestServiceArea)
		suite.Error(err)
		suite.Equal("RequestedPickupDate is required", err.Error())

		// No weight
		_, err = pricer.Price(testdatagen.DefaultContractCode, requestedPickupDate, 0, dopTestServiceArea)
		suite.Error(err)
		suite.Equal("Weight must be a minimum of 500", err.Error())

		// No service area
		_, err = pricer.Price(testdatagen.DefaultContractCode, requestedPickupDate, dshTestWeight, "")
		suite.Error(err)
		suite.Equal("ServiceArea is required", err.Error())
	})
}

func (suite *GHCRateEngineServiceSuite) setupDomesticOriginServiceItems() models.PaymentServiceItem {
	return suite.setupPaymentServiceItemWithParams(
		models.ReServiceCodeDOP,
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
				models.ServiceItemParamNameWeightBilledActual,
				models.ServiceItemParamTypeInteger,
				strconv.Itoa(dshTestWeight),
			},
			{
				models.ServiceItemParamNameServiceAreaOrigin,
				models.ServiceItemParamTypeString,
				dopTestServiceArea,
			},
		},
	)
}

func (suite *GHCRateEngineServiceSuite) setUpDomesticOriginData() {
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
				ServiceArea: dopTestServiceArea,
			},
		})

	domesticOriginService := testdatagen.MakeReService(suite.DB(),
		testdatagen.Assertions{
			ReService: models.ReService{
				Code: "DOP",
				Name: "Dom. Origin Price",
			},
		})

	domesticOriginPrice := models.ReDomesticServiceAreaPrice{
		ContractID:            contractYear.Contract.ID,
		DomesticServiceAreaID: serviceArea.ID,
		IsPeakPeriod:          true,
		ServiceID:             domesticOriginService.ID,
	}

	domesticOriginPeakPrice := domesticOriginPrice
	domesticOriginPeakPrice.PriceCents = 146
	suite.MustSave(&domesticOriginPeakPrice)

	domesticOriginNonpeakPrice := domesticOriginPrice
	domesticOriginNonpeakPrice.IsPeakPeriod = false
	domesticOriginNonpeakPrice.PriceCents = 127
	suite.MustSave(&domesticOriginNonpeakPrice)
}
