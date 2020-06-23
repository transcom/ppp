package ghcdieselfuelprice

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/transcom/mymove/pkg/testingsuite"
	"go.uber.org/zap"
)

type GHCDieselFuelPriceServiceSuite struct {
	testingsuite.PopTestSuite
	logger *zap.Logger
}

func (suite *GHCDieselFuelPriceServiceSuite) SetupTest() {
	suite.DB().TruncateAll()
}

func TestGHCDieselFuelPriceServiceSuite(t *testing.T) {
	// Use a no-op logger during testing
	logger := zap.NewNop()

	ts := &GHCDieselFuelPriceServiceSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage()),
		logger:       logger,
	}
	suite.Run(t, ts)
	ts.PopTestSuite.TearDown()
}

func (suite *GHCDieselFuelPriceServiceSuite) helperStubEIAData(url string) (eiaData, error) {
	var eiaData eiaData
	re := suite.helperRemoveURLQuerystring(url)

	if re.MatchString("EIA Open Data API error - invalid or missing api_key") {
		eiaData.responseStatusCode = 200
		eiaData.RequestData = requestData {
			Command:  "series",
			SeriesID: "pet.emd_epd2d_pte_nus_dpg.ws",
		}
		eiaData.ErrorData = errorData {
			Error: "invalid or missing api_key. For key registration, documentation, and examples see https://www.eia.gov/developer/",
		}
		eiaData.SeriesData = []seriesData{}

		return eiaData, nil
	}

	if re.MatchString("EIA Open Data API error - invalid series_id") {
		eiaData.responseStatusCode = 200
		eiaData.RequestData = requestData {
			Command:  "series",
			SeriesID: "pet.emd_epd2d_pte_nus_dpg.ws",
		}
		eiaData.ErrorData = errorData {
			Error: "invalid series_id. For key registration, documentation, and examples see https://www.eia.gov/developer/",
		}
		eiaData.SeriesData = []seriesData{}

		return eiaData, nil
	}

	if re.MatchString("nil series data") {
		eiaData.responseStatusCode = 200
		eiaData.RequestData = requestData {
			Command:  "series",
			SeriesID: "pet.emd_epd2d_pte_nus_dpg.ws",
		}
		eiaData.ErrorData = errorData{}
		eiaData.SeriesData = []seriesData{}

		return eiaData, nil
	}

	if re.MatchString("extract diesel fuel price data") {
		eiaData.responseStatusCode = 200
		eiaData.RequestData = requestData {
			Command:  "series",
			SeriesID: "pet.emd_epd2d_pte_nus_dpg.ws",
		}
		eiaData.ErrorData = errorData{}
		eiaData.SeriesData = []seriesData {
			0: {
				Updated: "2020-06-08T19:30:09-0400",
				Data: [][]interface{} {
					0: {0: "20200608", 1: 2.396},
					1: {0: "20200601", 1: 2.386},
					2: {0: "20200525", 1: 2.39},
					3: {0: "20200518", 1: 2.386},
					4: {0: "20200511", 1: 2.394},
				},
			},
		}

		return eiaData, nil
	}

	return eiaData, nil
}

func (suite *GHCDieselFuelPriceServiceSuite) helperRemoveURLQuerystring(url string) *regexp.Regexp {
	re := regexp.MustCompile(`%20`)
	url = re.ReplaceAllLiteralString(url, ` `)
	url = strings.Split(url, "?")[0]
	re = regexp.MustCompile(`^` + url + `.*`)

	return re
}