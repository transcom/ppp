package primeapi

import (
	"errors"
	"fmt"
	"net/http/httptest"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/etag"
	"github.com/transcom/mymove/pkg/handlers/primeapi/payloads"
	"github.com/transcom/mymove/pkg/unit"

	movetaskorder "github.com/transcom/mymove/pkg/services/move_task_order"

	"github.com/transcom/mymove/pkg/gen/primemessages"

	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gobuffalo/validate/v3"

	"github.com/stretchr/testify/mock"

	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/services/mocks"

	"github.com/transcom/mymove/pkg/models"

	mtoserviceitemops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/mto_service_item"
	"github.com/transcom/mymove/pkg/handlers"
	mtoserviceitem "github.com/transcom/mymove/pkg/services/mto_service_item"
	"github.com/transcom/mymove/pkg/services/query"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestCreateMTOServiceItemHandler() {
	mto := testdatagen.MakeAvailableMove(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: mto,
	})
	testdatagen.MakeDOFSITReService(suite.DB(), testdatagen.Assertions{
		ReService: models.ReService{
			ID: uuid.FromStringOrNil("9dc919da-9b66-407b-9f17-05c0f03fcb50"),
		},
	})
	builder := query.NewQueryBuilder(suite.DB())
	mtoChecker := movetaskorder.NewMoveTaskOrderChecker(suite.DB())

	req := httptest.NewRequest("POST", "/mto-service-items", nil)
	reason := "lorem ipsum"
	sitEntryDate := time.Now()
	sitPostalCode := "00000"

	// Customer gets new pickup address for SIT Origin Pickup (DOPSIT) which gets added when
	// creating DOFSIT (SIT origin first day).
	actualPickupAddress := testdatagen.MakeAddress2(suite.DB(), testdatagen.Assertions{})

	mtoServiceItem := models.MTOServiceItem{
		MoveTaskOrderID:             mto.ID,
		MTOShipmentID:               &mtoShipment.ID,
		ReService:                   models.ReService{Code: models.ReServiceCodeDOFSIT},
		Reason:                      &reason,
		SITEntryDate:                &sitEntryDate,
		SITPostalCode:               &sitPostalCode,
		SITOriginHHGActualAddress:   &actualPickupAddress,
		SITOriginHHGActualAddressID: &actualPickupAddress.ID,
	}

	params := mtoserviceitemops.CreateMTOServiceItemParams{
		HTTPRequest: req,
		Body:        payloads.MTOServiceItem(&mtoServiceItem),
	}

	suite.T().Run("Successful POST - Integration Test", func(t *testing.T) {
		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemOK{}, response)

		okResponse := response.(*mtoserviceitemops.CreateMTOServiceItemOK)
		suite.NotZero(okResponse.Payload[0].ID())
	})

	suite.T().Run("POST failure - 500", func(t *testing.T) {
		mockCreator := mocks.MTOServiceItemCreator{}
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockCreator,
			mtoChecker,
		}
		err := errors.New("ServerError")

		mockCreator.On("CreateMTOServiceItem",
			mock.Anything,
		).Return(nil, nil, err)

		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemInternalServerError{}, response)

		errResponse := response.(*mtoserviceitemops.CreateMTOServiceItemInternalServerError)
		suite.Equal(handlers.InternalServerErrMessage, *errResponse.Payload.Title, "Payload title is wrong")

	})

	suite.T().Run("POST failure - 422 Unprocessable Entity Error", func(t *testing.T) {
		mockCreator := mocks.MTOServiceItemCreator{}
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockCreator,
			mtoChecker,
		}
		// InvalidInputError should generate an UnprocessableEntity response
		err := services.InvalidInputError{}

		mockCreator.On("CreateMTOServiceItem",
			mock.Anything,
		).Return(nil, nil, err)

		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemUnprocessableEntity{}, response)
	})

	suite.T().Run("POST failure - 409 Conflict Error", func(t *testing.T) {
		mockCreator := mocks.MTOServiceItemCreator{}
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockCreator,
			mtoChecker,
		}
		// ConflictError should generate a Conflict response
		err := services.ConflictError{}

		mockCreator.On("CreateMTOServiceItem",
			mock.Anything,
		).Return(nil, nil, err)

		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemConflict{}, response)
	})

	suite.T().Run("POST failure - 404", func(t *testing.T) {
		mockCreator := mocks.MTOServiceItemCreator{}
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockCreator,
			mtoChecker,
		}
		err := services.NotFoundError{}

		mockCreator.On("CreateMTOServiceItem",
			mock.Anything,
		).Return(nil, nil, err)

		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemNotFound{}, response)
	})

	suite.T().Run("POST failure - 404 - MTO is not available to Prime", func(t *testing.T) {
		mtoNotAvailable := testdatagen.MakeDefaultMove(suite.DB())

		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		body := payloads.MTOServiceItem(&mtoServiceItem)
		body.SetMoveTaskOrderID(handlers.FmtUUID(mtoNotAvailable.ID))

		paramsNotAvailable := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        body,
		}

		response := handler.Handle(paramsNotAvailable)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemNotFound{}, response)

		typedResponse := response.(*mtoserviceitemops.CreateMTOServiceItemNotFound)
		suite.Contains(*typedResponse.Payload.Detail, mtoNotAvailable.ID.String())
	})

	suite.T().Run("POST failure - 404 - Integration - ShipmentID not linked by MoveTaskOrderID", func(t *testing.T) {
		mto2 := testdatagen.MakeAvailableMove(suite.DB())
		mtoShipment2 := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			Move: mto2,
		})
		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		body := payloads.MTOServiceItem(&mtoServiceItem)
		body.SetMoveTaskOrderID(handlers.FmtUUID(mtoShipment.MoveTaskOrderID))
		body.SetMtoShipmentID(strfmt.UUID(mtoShipment2.ID.String()))

		newParams := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        body,
		}

		response := handler.Handle(newParams)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemNotFound{}, response)
	})

	suite.T().Run("POST failure - 422 - Model validation errors", func(t *testing.T) {
		mockCreator := mocks.MTOServiceItemCreator{}
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockCreator,
			mtoChecker,
		}
		verrs := validate.NewErrors()
		verrs.Add("test", "testing")

		mockCreator.On("CreateMTOServiceItem",
			mock.Anything,
		).Return(nil, verrs, nil)

		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemUnprocessableEntity{}, response)
	})

	suite.T().Run("POST failure - 422 - modelType() not supported", func(t *testing.T) {
		mockCreator := mocks.MTOServiceItemCreator{}
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockCreator,
			mtoChecker,
		}
		err := services.NotFoundError{}

		mockCreator.On("CreateMTOServiceItem",
			mock.Anything,
		).Return(nil, nil, err)

		mtoServiceItem := models.MTOServiceItem{
			MoveTaskOrderID: mto.ID,
			MTOShipmentID:   &mtoShipment.ID,
			ReService:       models.ReService{Code: models.ReServiceCodeMS},
			Reason:          nil,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemUnprocessableEntity{}, response)
	})
}

func (suite *HandlerSuite) TestCreateMTOServiceItemDomesticCratingHandler() {
	mto := testdatagen.MakeAvailableMove(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: mto,
	})
	testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
		ReService: models.ReService{
			Code: "DCRT",
		},
	})
	testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
		ReService: models.ReService{
			Code: "DUCRT",
		},
	})
	testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
		ReService: models.ReService{
			Code: "DCRTSA",
		},
	})
	builder := query.NewQueryBuilder(suite.DB())
	mtoChecker := movetaskorder.NewMoveTaskOrderChecker(suite.DB())

	req := httptest.NewRequest("POST", "/mto-service-items", nil)

	mtoServiceItem := models.MTOServiceItem{
		MoveTaskOrderID: mto.ID,
		MTOShipmentID:   &mtoShipment.ID,
		Description:     handlers.FmtString("description"),
		Dimensions: models.MTOServiceItemDimensions{
			models.MTOServiceItemDimension{
				Type:   models.DimensionTypeItem,
				Length: 1000,
				Height: 1000,
				Width:  1000,
			},
			models.MTOServiceItemDimension{
				Type:   models.DimensionTypeCrate,
				Length: 10000,
				Height: 10000,
				Width:  10000,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.T().Run("Successful POST - Integration Test - Domestic Crating", func(t *testing.T) {
		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		mtoServiceItem.ReService.Code = models.ReServiceCodeDCRT
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemOK{}, response)

		okResponse := response.(*mtoserviceitemops.CreateMTOServiceItemOK)
		suite.NotZero(okResponse.Payload[0].ID())
	})

	suite.T().Run("Successful POST - Integration Test - Domestic Uncrating", func(t *testing.T) {
		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		mtoServiceItem.ReService.Code = models.ReServiceCodeDUCRT
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemOK{}, response)

		okResponse := response.(*mtoserviceitemops.CreateMTOServiceItemOK)
		suite.NotZero(okResponse.Payload[0].ID())
	})

	suite.T().Run("Successful POST - Integration Test - Domestic Crating Standalone", func(t *testing.T) {
		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		mtoServiceItem.ReService.Code = models.ReServiceCodeDCRTSA
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemOK{}, response)

		okResponse := response.(*mtoserviceitemops.CreateMTOServiceItemOK)
		suite.NotZero(okResponse.Payload[0].ID())
	})

	suite.T().Run("POST failure - 422", func(t *testing.T) {
		mockCreator := mocks.MTOServiceItemCreator{}
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockCreator,
			mtoChecker,
		}
		err := errors.New("ServerError")

		mockCreator.On("CreateMTOServiceItem",
			mock.Anything,
		).Return(nil, nil, err)

		mtoServiceItem.ReService.Code = models.ReServiceCodeDCRTSA
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		var height int32 = 0
		params.Body.(*primemessages.MTOServiceItemDomesticCrating).Crate.Height = &height
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemUnprocessableEntity{}, response)
	})
}

func (suite *HandlerSuite) TestCreateMTOServiceItemOriginSITHandler() {
	// Under test: createMTOServiceItemHandler function,
	// - no DOPSIT standalone
	// -  DOASIT standalone with DOFSIT

	mto := testdatagen.MakeAvailableMove(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: mto,
	})
	testdatagen.MakeDOFSITReService(suite.DB(), testdatagen.Assertions{})
	builder := query.NewQueryBuilder(suite.DB())
	mtoChecker := movetaskorder.NewMoveTaskOrderChecker(suite.DB())

	reason := "lorem ipsum"
	sitEntryDate := time.Now()
	sitPostalCode := "00000"

	mtoServiceItem := models.MTOServiceItem{
		MoveTaskOrderID: mto.ID,
		MTOShipmentID:   &mtoShipment.ID,
		ReService:       models.ReService{},
		Reason:          &reason,
		SITEntryDate:    &sitEntryDate,
		SITPostalCode:   &sitPostalCode,
	}

	suite.T().Run("POST failure - 422 Cannot create DOPSIT standalone", func(t *testing.T) {
		// Under test: createMTOServiceItemHandler function
		// Set up:     We hit the endpoint with a DOPSIT MTOServiceItem
		// Expected outcome:
		//             Receive a 422 - Unprocessable Entity
		// SETUP
		// Create the payload
		mtoServiceItem.ReService.Code = models.ReServiceCodeDOPSIT

		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		// CALL FUNCTION UNDER TEST
		req := httptest.NewRequest("POST", "/mto-service-items", nil)
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		// CHECK RESULTS
		suite.Error(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemUnprocessableEntity{}, response)

	})

	suite.T().Run("POST Failure - Cannot create DOASIT without DOFSIT", func(t *testing.T) {
		// Under test: createMTOServiceItemHandler function
		// Set up:     We hit the endpoint with a standalone DOASIT MTOServiceItem, no DOFSIT
		// Expected outcome:
		//             Receive a 404 - Not Found
		// SETUP
		// Create the payload
		mtoServiceItem.ReService.Code = models.ReServiceCodeDOASIT

		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		// CALL FUNCTION UNDER TEST
		req := httptest.NewRequest("POST", "/mto-service-items", nil)
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		// CHECK RESULTS
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemNotFound{}, response)

	})

	suite.T().Run("Successful POST - Create DOASIT with DOFSIT", func(t *testing.T) {
		// Under test: createMTOServiceItemHandler function
		// Set up:     We hit the endpoint with a standalone DOASIT MTOServiceItem
		// Expected outcome:
		//             Receive a 404 - Not Found
		// SETUP
		// Create the payload
		testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
			ReService: models.ReService{
				Code: models.ReServiceCodeDOFSIT,
			},
			Move:        mto,
			MTOShipment: mtoShipment,
		})

		mtoServiceItem.ReService.Code = models.ReServiceCodeDOASIT

		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		// CALL FUNCTION UNDER TEST
		req := httptest.NewRequest("POST", "/mto-service-items", nil)
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		// CHECK RESULTS
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemOK{}, response)

	})

}

func (suite *HandlerSuite) TestCreateMTOServiceItemOriginSITHandlerWithDOFSITNoAddress() {
	// Under test: createMTOServiceItemHandler function,
	// - fail to create DOFSIT because of missing sitHHGActualAddress

	mto := testdatagen.MakeAvailableMove(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: mto,
	})
	testdatagen.MakeDOFSITReService(suite.DB(), testdatagen.Assertions{})
	builder := query.NewQueryBuilder(suite.DB())
	mtoChecker := movetaskorder.NewMoveTaskOrderChecker(suite.DB())

	reason := "lorem ipsum"
	sitEntryDate := time.Now()
	sitPostalCode := "00000"

	mtoServiceItem := models.MTOServiceItem{
		MoveTaskOrderID: mto.ID,
		MTOShipmentID:   &mtoShipment.ID,
		ReService:       models.ReService{},
		Reason:          &reason,
		SITEntryDate:    &sitEntryDate,
		SITPostalCode:   &sitPostalCode,
	}

	suite.T().Run("Successful POST - Create DOFSIT", func(t *testing.T) {
		// Under test: createMTOServiceItemHandler function
		// Set up:     We hit the endpoint with a standalone DOFSIT MTOServiceItem
		// Expected outcome:
		//             CreateMTOServiceItemUnprocessableEntity
		// SETUP
		// Create the payload

		mtoServiceItem.ReService.Code = models.ReServiceCodeDOFSIT

		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		// CALL FUNCTION UNDER TEST
		req := httptest.NewRequest("POST", "/mto-service-items", nil)
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		// CHECK RESULTS
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemUnprocessableEntity{}, response)
		unprocessableEntity := response.(*mtoserviceitemops.CreateMTOServiceItemUnprocessableEntity)
		suite.Contains(*unprocessableEntity.Payload.Detail, "must have the sitHHGActualOrigin")
	})

}

func (suite *HandlerSuite) TestCreateMTOServiceItemOriginSITHandlerWithDOFSITWithAddress() {
	// Under test: createMTOServiceItemHandler function,
	// - no DOPSIT standalone
	// -  DOASIT standalone with DOFSIT

	mto := testdatagen.MakeAvailableMove(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: mto,
	})
	testdatagen.MakeDOFSITReService(suite.DB(), testdatagen.Assertions{})
	builder := query.NewQueryBuilder(suite.DB())
	mtoChecker := movetaskorder.NewMoveTaskOrderChecker(suite.DB())

	reason := "lorem ipsum"
	sitEntryDate := time.Now()
	sitPostalCode := "00000"

	// Original customer pickup address
	originalPickupAddress := mtoShipment.PickupAddress
	originalPickupAddressID := mtoShipment.PickupAddressID

	// Customer gets new pickup address
	actualPickupAddress := testdatagen.MakeAddress2(suite.DB(), testdatagen.Assertions{})

	mtoServiceItem := models.MTOServiceItem{
		MoveTaskOrderID:             mto.ID,
		MTOShipmentID:               &mtoShipment.ID,
		ReService:                   models.ReService{},
		Reason:                      &reason,
		SITEntryDate:                &sitEntryDate,
		SITPostalCode:               &sitPostalCode,
		SITOriginHHGActualAddress:   &actualPickupAddress,
		SITOriginHHGActualAddressID: &actualPickupAddress.ID,
	}

	// Verify the addresses for original pickup and new pickup are not the same
	suite.NotEqual(originalPickupAddressID, mtoServiceItem.SITOriginHHGActualAddressID, "address ID is not the same")
	suite.NotEqual(originalPickupAddress.StreetAddress1, mtoServiceItem.SITOriginHHGActualAddress.StreetAddress1, "street address is not the same")
	suite.NotEqual(originalPickupAddress.City, mtoServiceItem.SITOriginHHGActualAddress.City, "city is not the same")
	suite.NotEqual(originalPickupAddress.PostalCode, mtoServiceItem.SITOriginHHGActualAddress.PostalCode, "zip is not the same")

	suite.T().Run("Successful POST - Create DOFSIT", func(t *testing.T) {
		// Under test: createMTOServiceItemHandler function
		// Set up:     We hit the endpoint with a standalone DOFSIT MTOServiceItem
		// Expected outcome:
		//             Successful creation of DOFSIT with DOPSIT added
		// SETUP
		// Create the payload

		mtoServiceItem.ReService.Code = models.ReServiceCodeDOFSIT

		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		// CALL FUNCTION UNDER TEST
		req := httptest.NewRequest("POST", "/mto-service-items", nil)
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		// CHECK RESULTS
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemOK{}, response)

		// Verify address was updated on MTO Shipment
		var updatedMTOShipment models.MTOShipment
		suite.NoError(suite.DB().Eager("PickupAddress").Find(&updatedMTOShipment, mtoShipment.ID))

		// Verify the HHG pickup address is the actual address on the shipment
		// TODO remove if not needed: suite.NotNil(updatedMTOShipment.PickupAddressID, "address ID is not nil")
		suite.Equal(actualPickupAddress.ID, *updatedMTOShipment.PickupAddressID, "hhg actual address id is the same")
		suite.Equal(actualPickupAddress.StreetAddress1, updatedMTOShipment.PickupAddress.StreetAddress1, "hhg actual street address is the same")
		suite.Equal(actualPickupAddress.City, updatedMTOShipment.PickupAddress.City, "hhg actual city is the same")
		suite.Equal(actualPickupAddress.State, updatedMTOShipment.PickupAddress.State, "hhg actual state is the same")
		suite.Equal(actualPickupAddress.PostalCode, updatedMTOShipment.PickupAddress.PostalCode, "hhg actual zip is the same")

		// Verify address on SIT service item
		okResponse := response.(*mtoserviceitemops.CreateMTOServiceItemOK)
		suite.NotZero(okResponse.Payload[0].ID())
		findDOFSIT := false

		for _, sp := range okResponse.Payload {
			var serviceItem models.MTOServiceItem
			id := sp.ID()
			err := suite.DB().Find(&serviceItem, id)
			suite.NoError(err)
			if serviceItem.ReService.Code == models.ReServiceCodeDOFSIT || serviceItem.ReService.Code == models.ReServiceCodeDOPSIT {
				findDOFSIT = true
				// Verify the HHG original pickup address is the original address on the service item
				suite.Equal(originalPickupAddressID, mtoServiceItem.SITOriginHHGOriginalAddressID, "original address ID is the same")
				suite.Equal(originalPickupAddress.StreetAddress1, mtoServiceItem.SITOriginHHGOriginalAddress.StreetAddress1, "original street address is the same")
				suite.Equal(originalPickupAddress.City, mtoServiceItem.SITOriginHHGOriginalAddress.City, "original city is the same")
				suite.Equal(originalPickupAddress.State, mtoServiceItem.SITOriginHHGOriginalAddress.State, "original state is the same")
				suite.Equal(originalPickupAddress.PostalCode, mtoServiceItem.SITOriginHHGOriginalAddress.PostalCode, "original zip is the same")

				// Verify the HHG pickup address is the actual address on the service item
				suite.Equal(updatedMTOShipment.PickupAddressID, mtoServiceItem.SITOriginHHGActualAddressID, "shipment actual address ID is the same")
				suite.Equal(updatedMTOShipment.PickupAddress.StreetAddress1, mtoServiceItem.SITOriginHHGActualAddress.StreetAddress1, "shipment actual street address is the same")
				suite.Equal(updatedMTOShipment.PickupAddress.City, mtoServiceItem.SITOriginHHGActualAddress.City, "shipment actual city is the same")
				suite.Equal(updatedMTOShipment.PickupAddress.State, mtoServiceItem.SITOriginHHGActualAddress.State, "shipment actual state is the same")
				suite.Equal(updatedMTOShipment.PickupAddress.PostalCode, mtoServiceItem.SITOriginHHGActualAddress.PostalCode, "shipment actual zip is the same")
			}
		}
		suite.Equal(true, findDOFSIT, "Found expected ReServiceCodeDOFSIT")
	})

}

func (suite *HandlerSuite) TestCreateMTOServiceItemDestSITHandler() {

	mto := testdatagen.MakeAvailableMove(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: mto,
	})
	testdatagen.MakeDDFSITReService(suite.DB())
	builder := query.NewQueryBuilder(suite.DB())
	mtoChecker := movetaskorder.NewMoveTaskOrderChecker(suite.DB())

	req := httptest.NewRequest("POST", "/mto-service-items", nil)
	sitEntryDate := time.Now()
	mtoServiceItem := models.MTOServiceItem{
		MoveTaskOrderID: mto.ID,
		MTOShipmentID:   &mtoShipment.ID,
		ReService:       models.ReService{Code: models.ReServiceCodeDDFSIT},
		Description:     handlers.FmtString("description"),
		CustomerContacts: models.MTOServiceItemCustomerContacts{
			models.MTOServiceItemCustomerContact{
				Type:                       models.CustomerContactTypeFirst,
				TimeMilitary:               "0400Z",
				FirstAvailableDeliveryDate: time.Now(),
			},
			models.MTOServiceItemCustomerContact{
				Type:                       models.CustomerContactTypeSecond,
				TimeMilitary:               "0400Z",
				FirstAvailableDeliveryDate: time.Now(),
			},
		},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		SITEntryDate: &sitEntryDate,
	}
	params := mtoserviceitemops.CreateMTOServiceItemParams{
		HTTPRequest: req,
		Body:        payloads.MTOServiceItem(&mtoServiceItem),
	}

	suite.T().Run("Successful POST - Integration Test", func(t *testing.T) {
		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemOK{}, response)

		okResponse := response.(*mtoserviceitemops.CreateMTOServiceItemOK)
		suite.NotZero(okResponse.Payload[0].ID())
	})

	suite.T().Run("Successful POST - Create DDASIT standalone", func(t *testing.T) {
		mtoServiceItem.ReService.Code = models.ReServiceCodeDDASIT
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemOK{}, response)

		okResponse := response.(*mtoserviceitemops.CreateMTOServiceItemOK)
		suite.NotZero(okResponse.Payload[0].ID())

	})

	suite.T().Run("POST Failure - Cannot create DDASIT without DDFSIT", func(t *testing.T) {
		mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{})

		mtoServiceItem.ReService.Code = models.ReServiceCodeDDASIT
		mtoServiceItem.MTOShipment = mtoShipment
		mtoServiceItem.MTOShipmentID = &mtoShipment.ID

		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		// CHECK RESULTS
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemNotFound{}, response)

	})

	suite.T().Run("POST failure - 422 Cannot create DDDSIT standalone", func(t *testing.T) {
		// Under test: createMTOServiceItemHandler function
		// Set up:     We hit the endpoint with a DDDSIT MTOServiceItem
		// Expected outcome:
		//             Receive a 422 - Unprocessable Entity
		// SETUP
		// Create the payload
		mtoServiceItem.ReService.Code = models.ReServiceCodeDDDSIT
		creator := mtoserviceitem.NewMTOServiceItemCreator(builder)
		handler := CreateMTOServiceItemHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			creator,
			mtoChecker,
		}

		// CALL FUNCTION UNDER TEST
		req := httptest.NewRequest("POST", "/mto-service-items", nil)
		params := mtoserviceitemops.CreateMTOServiceItemParams{
			HTTPRequest: req,
			Body:        payloads.MTOServiceItem(&mtoServiceItem),
		}

		// CHECK RESULTS
		suite.Error(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)
		suite.IsType(&mtoserviceitemops.CreateMTOServiceItemNotFound{}, response)

	})
}

func (suite *HandlerSuite) TestUpdateMTOServiceItemDDDSIT() {

	// Under test: updateMTOServiceItemHandler.Handle function
	//             MTOServiceItemUpdater.Update service object function
	// SETUP
	// Create the service item in the db for dofsit and dddsit
	timeNow := time.Now()
	dddsit := testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			AvailableToPrimeAt: &timeNow,
		},
		MTOServiceItem: models.MTOServiceItem{
			SITEntryDate: swag.Time(time.Now()),
		},
		ReService: models.ReService{
			Code: "DDDSIT",
		},
	})

	// Create the payload with the desired update
	reqPayload := &primemessages.UpdateMTOServiceItemSIT{
		ReServiceCode:    "DDDSIT",
		SitDepartureDate: *handlers.FmtDate(time.Now().AddDate(0, 0, 5)),
	}
	reqPayload.SetID(strfmt.UUID(dddsit.ID.String()))

	// Create the handler
	queryBuilder := query.NewQueryBuilder(suite.DB())
	handler := UpdateMTOServiceItemHandler{
		handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
		mtoserviceitem.NewMTOServiceItemUpdater(queryBuilder),
	}

	// create the params struct
	req := httptest.NewRequest("PATCH", fmt.Sprintf("/mto-service_items/%s", dddsit.ID), nil)
	eTag := etag.GenerateEtag(dddsit.UpdatedAt)
	params := mtoserviceitemops.UpdateMTOServiceItemParams{
		HTTPRequest: req,
		Body:        reqPayload,
		IfMatch:     eTag,
	}

	suite.T().Run("Successful PATCH - Updated SITDepartureDate on DOPSIT", func(t *testing.T) {
		// Under test: updateMTOServiceItemHandler.Handle function
		//             MTOServiceItemUpdater.Update service object function
		// Set up:     We create an mto service item using DOFSIT (which was created above)
		//             And send an update to the sit entry date
		// Expected outcome:
		//             Receive a success response with the SitDepartureDate updated

		// CALL FUNCTION UNDER TEST
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)

		// CHECK RESULTS
		suite.IsType(&mtoserviceitemops.UpdateMTOServiceItemOK{}, response)
		r := response.(*mtoserviceitemops.UpdateMTOServiceItemOK)
		resp1 := r.Payload

		respPayload := resp1.(*primemessages.MTOServiceItemDestSIT)
		suite.Equal(reqPayload.ID(), respPayload.ID())
		suite.Equal(reqPayload.SitDepartureDate.String(), respPayload.SitDepartureDate.String())

		// Return to good state for next test
		params.IfMatch = respPayload.ETag()

	})

	suite.T().Run("Failed PATCH - No DDDSIT found", func(t *testing.T) {
		// Under test: updateMTOServiceItemHandler.Handle function
		//             MTOServiceItemUpdater.Update service object function
		// Set up:     We use a non existent DDDSIT item
		//             And send an update to DOPSIT to the SitDepartureDate
		// Expected outcome:
		//             Receive a NotFound error response

		// SETUP
		// Replace the request path with a bad id that won't be found
		badUUID := uuid.Must(uuid.NewV4())
		badReq := httptest.NewRequest("PATCH", fmt.Sprintf("/mto-service_items/%s", badUUID), nil)
		params.HTTPRequest = badReq
		reqPayload.SetID(strfmt.UUID(badUUID.String()))

		// CALL FUNCTION UNDER TEST
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)

		// CHECK RESULTS
		suite.IsType(&mtoserviceitemops.UpdateMTOServiceItemNotFound{}, response)

		// return to good state for next test
		params.HTTPRequest = req
		reqPayload.SetID(strfmt.UUID(dddsit.ID.String()))
	})

	suite.T().Run("Failed PATCH - Payment request created", func(t *testing.T) {
		// Under test: updateMTOServiceItemHandler.Handle function
		//             MTOServiceItemUpdater.Update service object function
		// Set up:     We use a DDDSIT that already has a payment request associated
		//             Then try to update the SitDepartureDate on that
		// Expected outcome:
		//             Receive a ConflictError response

		// SETUP
		// Make a payment request and link to the dddsit service item
		paymentRequest := testdatagen.MakeDefaultPaymentRequest(suite.DB())
		cost := unit.Cents(20000)
		testdatagen.MakePaymentServiceItem(suite.DB(), testdatagen.Assertions{
			PaymentServiceItem: models.PaymentServiceItem{
				PriceCents: &cost,
			},
			PaymentRequest: paymentRequest,
			MTOServiceItem: dddsit,
		})

		// CALL FUNCTION UNDER TEST
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)

		// CHECK RESULTS
		suite.IsType(&mtoserviceitemops.UpdateMTOServiceItemConflict{}, response)
	})

}

func (suite *HandlerSuite) TestUpdateMTOServiceItemDOPSIT() {

	// Under test: updateMTOServiceItemHandler.Handle function
	//             MTOServiceItemUpdater.Update service object function
	// SETUP
	// Create the service item in the db for dofsit and DOPSIT
	timeNow := time.Now()
	dopsit := testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			AvailableToPrimeAt: &timeNow,
		},
		MTOServiceItem: models.MTOServiceItem{
			SITEntryDate: swag.Time(time.Now()),
		},
		ReService: models.ReService{
			Code: "DOPSIT",
		},
	})

	// Create the payload with the desired update
	reqPayload := &primemessages.UpdateMTOServiceItemSIT{
		ReServiceCode:    "DOPSIT",
		SitDepartureDate: *handlers.FmtDate(time.Now().AddDate(0, 0, 5)),
	}
	reqPayload.SetID(strfmt.UUID(dopsit.ID.String()))

	// Create the handler
	queryBuilder := query.NewQueryBuilder(suite.DB())
	handler := UpdateMTOServiceItemHandler{
		handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
		mtoserviceitem.NewMTOServiceItemUpdater(queryBuilder),
	}

	// create the params struct
	req := httptest.NewRequest("PATCH", fmt.Sprintf("/mto-service_items/%s", dopsit.ID), nil)
	eTag := etag.GenerateEtag(dopsit.UpdatedAt)
	params := mtoserviceitemops.UpdateMTOServiceItemParams{
		HTTPRequest: req,
		Body:        reqPayload,
		IfMatch:     eTag,
	}

	suite.T().Run("Successful PATCH - Updated SITDepartureDate on DOPSIT", func(t *testing.T) {
		// Under test: updateMTOServiceItemHandler.Handle function
		//             MTOServiceItemUpdater.Update service object function
		// Set up:     We create an mto service item using DOFSIT (which was created above)
		//             And send an update to the sit entry date
		// Expected outcome:
		//             Receive a success response with the SitDepartureDate updated

		// CALL FUNCTION UNDER TEST
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)

		// CHECK RESULTS
		suite.IsType(&mtoserviceitemops.UpdateMTOServiceItemOK{}, response)
		r := response.(*mtoserviceitemops.UpdateMTOServiceItemOK)
		resp1 := r.Payload

		respPayload := resp1.(*primemessages.MTOServiceItemOriginSIT)
		suite.Equal(reqPayload.ID(), respPayload.ID())
		suite.Equal(reqPayload.SitDepartureDate.String(), respPayload.SitDepartureDate.String())

		// Return to good state for next test
		params.IfMatch = respPayload.ETag()

	})

	suite.T().Run("Failed PATCH - No DOPSIT found", func(t *testing.T) {
		// Under test: updateMTOServiceItemHandler.Handle function
		//             MTOServiceItemUpdater.Update service object function
		// Set up:     We use a non existent DOPSIT item
		//             And send an update to DOPSIT to the SitDepartureDate
		// Expected outcome:
		//             Receive a NotFound error response

		// SETUP
		// Replace the request path with a bad id that won't be found
		badUUID := uuid.Must(uuid.NewV4())
		badReq := httptest.NewRequest("PATCH", fmt.Sprintf("/mto-service_items/%s", badUUID), nil)
		params.HTTPRequest = badReq
		reqPayload.SetID(strfmt.UUID(badUUID.String()))

		// CALL FUNCTION UNDER TEST
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)

		// CHECK RESULTS
		suite.IsType(&mtoserviceitemops.UpdateMTOServiceItemNotFound{}, response)

		// return to good state for next test
		params.HTTPRequest = req
		reqPayload.SetID(strfmt.UUID(dopsit.ID.String()))
	})

	suite.T().Run("Failed PATCH - Payment request created", func(t *testing.T) {
		// Under test: updateMTOServiceItemHandler.Handle function
		//             MTOServiceItemUpdater.Update service object function
		// Set up:     We use a DOPSIT that already has a payment request associated
		//             Then try to update the SitDepartureDate on that
		// Expected outcome:
		//             Receive a ConflictError response

		// SETUP
		// Make a payment request and link to the DOPSIT service item
		paymentRequest := testdatagen.MakeDefaultPaymentRequest(suite.DB())
		cost := unit.Cents(20000)
		testdatagen.MakePaymentServiceItem(suite.DB(), testdatagen.Assertions{
			PaymentServiceItem: models.PaymentServiceItem{
				PriceCents: &cost,
			},
			PaymentRequest: paymentRequest,
			MTOServiceItem: dopsit,
		})

		// CALL FUNCTION UNDER TEST
		suite.NoError(params.Body.Validate(strfmt.Default))
		response := handler.Handle(params)

		// CHECK RESULTS
		suite.IsType(&mtoserviceitemops.UpdateMTOServiceItemConflict{}, response)
	})

}
