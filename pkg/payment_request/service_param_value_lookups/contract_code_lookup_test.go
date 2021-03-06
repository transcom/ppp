package serviceparamvaluelookups

import (
	"testing"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/testdatagen"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services/ghcrateengine"
)

func (suite *ServiceParamValueLookupsSuite) TestContractCodeLookup() {
	key := models.ServiceItemParamNameContractCode

	suite.T().Run("golden path", func(t *testing.T) {
		mtoServiceItem := testdatagen.MakeDefaultMTOServiceItem(suite.DB())

		paramLookup, err := ServiceParamLookupInitialize(suite.DB(), suite.planner, mtoServiceItem.ID, uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), nil)
		suite.FatalNoError(err)

		valueStr, err := paramLookup.ServiceParamValue(key)
		suite.FatalNoError(err)
		suite.Equal(ghcrateengine.DefaultContractCode, valueStr)
	})

	suite.T().Run("golden path with param cache", func(t *testing.T) {

		// MS
		reService1 := testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
			ReService: models.ReService{
				Code: "MS",
			},
		})

		// MS
		mtoServiceItem1 := testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
			ReService: reService1,
		})

		// ContractCode
		serviceItemParamKey1 := testdatagen.MakeServiceItemParamKey(suite.DB(), testdatagen.Assertions{
			ServiceItemParamKey: models.ServiceItemParamKey{
				Key:         models.ServiceItemParamNameContractCode,
				Description: "contract code",
				Type:        models.ServiceItemParamTypeString,
				Origin:      models.ServiceItemParamOriginSystem,
			},
		})

		_ = testdatagen.MakeServiceParam(suite.DB(), testdatagen.Assertions{
			ServiceParam: models.ServiceParam{
				ServiceID:             mtoServiceItem1.ReServiceID,
				ServiceItemParamKeyID: serviceItemParamKey1.ID,
				ServiceItemParamKey:   serviceItemParamKey1,
			},
		})

		paramCache := ServiceParamsCache{}
		paramCache.Initialize(suite.DB())
		paramLookup, err := ServiceParamLookupInitialize(suite.DB(), suite.planner, mtoServiceItem1.ID, uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), &paramCache)
		suite.FatalNoError(err)

		valueStr, err := paramLookup.ServiceParamValue(serviceItemParamKey1.Key)
		suite.FatalNoError(err)
		suite.Equal(ghcrateengine.DefaultContractCode, valueStr)

		// Verify value from paramCache
		paramCacheValue := paramCache.ParamValue(*mtoServiceItem1.MTOShipmentID, key)
		suite.Equal(ghcrateengine.DefaultContractCode, *paramCacheValue)
	})
}
