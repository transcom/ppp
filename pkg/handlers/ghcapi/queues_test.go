package ghcapi

import (
	"errors"
	"net/http/httptest"
	"time"

	"github.com/stretchr/testify/mock"

	modelToPayload "github.com/transcom/mymove/pkg/handlers/ghcapi/internal/payloads"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"

	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/queues"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/models/roles"
	"github.com/transcom/mymove/pkg/services/mocks"
	moveorder "github.com/transcom/mymove/pkg/services/move_order"
	paymentrequest "github.com/transcom/mymove/pkg/services/payment_request"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestGetMoveQueuesHandler() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	hhgMoveType := models.SelectedMoveTypeHHG
	// Default Origin Duty Station GBLOC is LKNQ
	hhgMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusSUBMITTED,
		},
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	// Create a move with an origin duty station outside of office user GBLOC
	excludedMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
		},
		TransportationOffice: models.TransportationOffice{
			Gbloc: "AGFM",
		},
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: excludedMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	request := httptest.NewRequest("GET", "/queues/moves", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload := response.(*queues.GetMovesQueueOK).Payload

	order := hhgMove.Orders
	result := payload.QueueMoves[0]

	suite.Len(payload.QueueMoves, 1)
	suite.Equal(order.ServiceMember.ID.String(), result.Customer.ID.String())
	suite.Equal(*order.DepartmentIndicator, string(result.DepartmentIndicator))
	suite.Equal(order.OriginDutyStation.TransportationOffice.Gbloc, string(result.OriginGBLOC))
	suite.Equal(order.NewDutyStation.ID.String(), result.DestinationDutyStation.ID.String())
	suite.Equal(hhgMove.Locator, result.Locator)
	suite.Equal(int64(1), result.ShipmentsCount)

}

func (suite *HandlerSuite) TestGetMoveQueuesBranchFilter() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	hhgMoveType := models.SelectedMoveTypeHHG

	move := models.Move{
		SelectedMoveType: &hhgMoveType,
		Status:           models.MoveStatusSUBMITTED,
	}
	shipment := models.MTOShipment{
		Status: models.MTOShipmentStatusSubmitted,
	}

	// Create an order where the service member has an ARMY affiliation (default)
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move:        move,
		MTOShipment: shipment,
	})

	// Create an order where the service member has an AIR_FORCE affiliation
	airForce := models.AffiliationAIRFORCE
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		MTOShipment: shipment,
		Move:        move,
		ServiceMember: models.ServiceMember{
			Affiliation: &airForce,
		},
	})

	request := httptest.NewRequest("GET", "/queues/moves", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
		Branch:      models.StringPointer("AIR_FORCE"),
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload := response.(*queues.GetMovesQueueOK).Payload

	result := payload.QueueMoves[0]

	suite.Equal(1, len(payload.QueueMoves))
	suite.Equal("AIR_FORCE", result.Customer.Agency)
}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerStatuses() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	hhgMoveType := models.SelectedMoveTypeHHG
	// Default Origin Duty Station GBLOC is LKNQ
	hhgMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusSUBMITTED,
		},
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	// Create a shipment on hhgMove that has Rejected status
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusRejected,
		},
	})

	// Create an order with an origin duty station outside of office user GBLOC
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		TransportationOffice: models.TransportationOffice{
			Name:  "Fort Punxsutawney",
			Gbloc: "AGFM",
		},
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
		},
	})

	request := httptest.NewRequest("GET", "/move-task-orders/{moveTaskOrderID}", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	payload := response.(*queues.GetMovesQueueOK).Payload
	result := payload.QueueMoves[0]

	suite.Equal(ghcmessages.QueueMoveStatus("New move"), result.Status)

	// let's test for the Move approved status
	hhgMove.Status = models.MoveStatusAPPROVED
	_, _ = suite.DB().ValidateAndSave(&hhgMove)

	response = handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload = response.(*queues.GetMovesQueueOK).Payload

	result = payload.QueueMoves[0]

	suite.Equal(ghcmessages.QueueMoveStatus("Move approved"), result.Status)

	// Now let's test Approvals requested
	testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
		MTOServiceItem: models.MTOServiceItem{
			Status: models.MTOServiceItemStatusSubmitted,
		},
		Move: hhgMove,
	})

	response = handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload = response.(*queues.GetMovesQueueOK).Payload

	result = payload.QueueMoves[0]

	suite.Equal(ghcmessages.QueueMoveStatus("Approvals requested"), result.Status)

}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerFilters() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	hhgMoveType := models.SelectedMoveTypeHHG
	submittedMove := models.Move{
		SelectedMoveType: &hhgMoveType,
		Status:           models.MoveStatusSUBMITTED,
	}
	submittedShipment := models.MTOShipment{
		Status: models.MTOShipmentStatusSubmitted,
	}
	airForce := models.AffiliationAIRFORCE

	// New move with AIR_FORCE service member affiliation to test branch filter
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move:        submittedMove,
		MTOShipment: submittedShipment,
		ServiceMember: models.ServiceMember{
			Affiliation: &airForce,
		},
	})

	// Approvals requested
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusAPPROVED,
		},
		MTOShipment: submittedShipment,
	})

	// Move approved
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusAPPROVED,
		},
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusApproved,
		},
	})

	// Move DRAFT and CANCELLED should not be included
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusDRAFT,
		},
		MTOShipment: submittedShipment,
	})
	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusCANCELED,
		},
		MTOShipment: submittedShipment,
	})

	request := httptest.NewRequest("GET", "/queues/moves", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)

	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	suite.Run("loads results with all STATUS selected", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest: request,
			Status: []string{
				modelToPayload.QueueMoveStatusNEWMOVE,
				modelToPayload.QueueMoveStatusAPPROVALSREQUESTED,
				modelToPayload.QueueMoveStatusMOVEAPPROVED,
			},
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload
		suite.EqualValues(3, payload.TotalCount)
		suite.Len(payload.QueueMoves, 3)
	})

	suite.Run("loads results with one STATUS selected", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest: request,
			Status: []string{
				modelToPayload.QueueMoveStatusNEWMOVE,
			},
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload
		suite.EqualValues(1, payload.TotalCount)
		suite.Len(payload.QueueMoves, 1)
		suite.EqualValues(modelToPayload.QueueMoveStatusNEWMOVE, payload.QueueMoves[0].Status)
	})

	suite.Run("1 result with status New Move and branch AIR_FORCE", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest: request,
			Status: []string{
				modelToPayload.QueueMoveStatusNEWMOVE,
			},
			Branch: models.StringPointer("AIR_FORCE"),
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload
		suite.EqualValues(1, payload.TotalCount)
		suite.Len(payload.QueueMoves, 1)
		suite.EqualValues(modelToPayload.QueueMoveStatusNEWMOVE, payload.QueueMoves[0].Status)
		suite.Equal("AIR_FORCE", payload.QueueMoves[0].Customer.Agency)
	})

	suite.Run("No results with status New Move and branch ARMY", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest: request,
			Status: []string{
				modelToPayload.QueueMoveStatusNEWMOVE,
			},
			Branch: models.StringPointer("ARMY"),
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload
		suite.EqualValues(0, payload.TotalCount)
		suite.Len(payload.QueueMoves, 0)
	})
}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerCustomerInfoFilters() {
	dutyStation1 := testdatagen.MakeDutyStation(suite.DB(), testdatagen.Assertions{
		DutyStation: models.DutyStation{
			Name: "This Other Station",
		},
	})

	dutyStation2 := testdatagen.MakeDefaultDutyStation(suite.DB())

	officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{})

	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	hhgMoveType := models.SelectedMoveTypeHHG
	// Default Origin Duty Station GBLOC is LKNQ

	serviceMember1 := testdatagen.MakeServiceMember(suite.DB(), testdatagen.Assertions{
		Stub: true,
		ServiceMember: models.ServiceMember{
			FirstName: models.StringPointer("Zoya"),
			LastName:  models.StringPointer("Darvish"),
			Edipi:     models.StringPointer("11111"),
		},
	})

	serviceMember2 := testdatagen.MakeServiceMember(suite.DB(), testdatagen.Assertions{
		Stub: true,
		ServiceMember: models.ServiceMember{
			FirstName: models.StringPointer("Owen"),
			LastName:  models.StringPointer("Nance"),
			Edipi:     models.StringPointer("22222"),
		},
	})

	move1 := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusSUBMITTED,
		},
		Order: models.Order{
			OriginDutyStation:   &dutyStation1,
			OriginDutyStationID: &dutyStation1.ID,
			NewDutyStation:      dutyStation1,
			NewDutyStationID:    dutyStation1.ID,
		},
		ServiceMember: serviceMember1,
	})

	move2 := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
			Status:           models.MoveStatusSUBMITTED,
		},
		Order: models.Order{
			OriginDutyStation:   &dutyStation2,
			OriginDutyStationID: &dutyStation2.ID,
			NewDutyStation:      dutyStation2,
			NewDutyStationID:    dutyStation2.ID,
		},
		ServiceMember: serviceMember2,
	})

	shipment := models.MTOShipment{
		Status: models.MTOShipmentStatusSubmitted,
	}

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move:        move1,
		MTOShipment: shipment,
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move:        move2,
		MTOShipment: shipment,
	})

	request := httptest.NewRequest("GET", "/queues/moves", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)

	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	suite.Run("returns unfiltered results", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest: request,
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload

		suite.Len(payload.QueueMoves, 2)
	})

	suite.Run("returns results matching last name search term", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest: request,
			LastName:    models.StringPointer("Nan"),
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload
		result := payload.QueueMoves[0]

		suite.Len(payload.QueueMoves, 1)
		suite.Equal("Nance", result.Customer.LastName)
	})

	suite.Run("returns results matching Dod ID search term", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest: request,
			DodID:       serviceMember1.Edipi,
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload
		result := payload.QueueMoves[0]

		suite.Len(payload.QueueMoves, 1)
		suite.Equal("11111", result.Customer.DodID)
	})

	suite.Run("returns results matching Move ID search term", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest: request,
			MoveID:      &move1.Locator,
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload
		result := payload.QueueMoves[0]

		suite.Len(payload.QueueMoves, 1)
		suite.Equal(move1.Locator, result.Locator)

	})

	suite.Run("returns results matching DestinationDutyStation name search term", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest:            request,
			DestinationDutyStation: &dutyStation1.Name,
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload
		result := payload.QueueMoves[0]

		suite.Len(payload.QueueMoves, 1)
		suite.Equal("This Other Station", result.DestinationDutyStation.Name)
	})

	suite.Run("returns results with multiple filters applied", func() {
		params := queues.GetMovesQueueParams{
			HTTPRequest:            request,
			LastName:               models.StringPointer("Dar"),
			DodID:                  serviceMember1.Edipi,
			MoveID:                 &move1.Locator,
			DestinationDutyStation: &dutyStation1.Name,
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		payload := response.(*queues.GetMovesQueueOK).Payload

		suite.Len(payload.QueueMoves, 1)
	})

}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerUnauthorizedRole() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTIO,
	})

	request := httptest.NewRequest("GET", "/queues/moves", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueForbidden{}, response)
}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerUnauthorizedUser() {
	serviceUser := testdatagen.MakeDefaultServiceMember(suite.DB())
	serviceUser.User.Roles = append(serviceUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeCustomer,
	})

	request := httptest.NewRequest("GET", "/queues/moves", nil)
	request = suite.AuthenticateRequest(request, serviceUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueForbidden{}, response)
}

func (suite *HandlerSuite) TestGetMoveQueuesHandlerEmptyResults() {
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	officeUser.User.Roles = append(officeUser.User.Roles, roles.Role{
		RoleType: roles.RoleTypeTOO,
	})

	// Create an order with an origin duty station outside of office user GBLOC
	hhgMoveType := models.SelectedMoveTypeHHG
	excludedMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
		},
		TransportationOffice: models.TransportationOffice{
			Gbloc: "AGFM",
		},
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: excludedMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	request := httptest.NewRequest("GET", "/queues/moves", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetMovesQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetMovesQueueHandler{
		context,
		moveorder.NewMoveOrderFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetMovesQueueOK{}, response)
	payload := response.(*queues.GetMovesQueueOK).Payload

	suite.Len(payload.QueueMoves, 0)
}

func (suite *HandlerSuite) TestGetPaymentRequestsQueueHandler() {
	officeUser := testdatagen.MakeTIOOfficeUser(suite.DB(), testdatagen.Assertions{})

	hhgMoveType := models.SelectedMoveTypeHHG
	// Default Origin Duty Station GBLOC is LKNQ
	hhgMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
		},
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	// Fake this as a day and a half in the past so floating point age values can be tested
	prevCreatedAt := time.Now().Add(time.Duration(time.Hour * -36))

	actualPaymentRequest := testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
		Move: hhgMove,
		PaymentRequest: models.PaymentRequest{
			CreatedAt: prevCreatedAt,
		},
	})

	// Create an order with an origin duty station outside of office user GBLOC
	excludedPaymentRequest := testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
		TransportationOffice: models.TransportationOffice{
			Gbloc: "AGFM",
		},
		Move: models.Move{
			SelectedMoveType: &hhgMoveType,
		},
	})

	testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: excludedPaymentRequest.MoveTaskOrder,
		MTOShipment: models.MTOShipment{
			Status: models.MTOShipmentStatusSubmitted,
		},
	})

	request := httptest.NewRequest("GET", "/queues/payment-requests", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetPaymentRequestsQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetPaymentRequestsQueueHandler{
		context,
		paymentrequest.NewPaymentRequestListFetcher(suite.DB()),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)

	suite.Assertions.IsType(&queues.GetPaymentRequestsQueueOK{}, response)
	payload := response.(*queues.GetPaymentRequestsQueueOK).Payload

	suite.Len(payload.QueuePaymentRequests, 1)

	paymentRequest := *payload.QueuePaymentRequests[0]

	suite.Equal(actualPaymentRequest.ID.String(), paymentRequest.ID.String())
	suite.Equal(actualPaymentRequest.MoveTaskOrderID.String(), paymentRequest.MoveID.String())
	suite.Equal(hhgMove.Orders.ServiceMemberID.String(), paymentRequest.Customer.ID.String())
	suite.Equal(string(paymentRequest.Status), "Payment requested")

	createdAt := actualPaymentRequest.CreatedAt
	age := int64(2)

	suite.Equal(age, paymentRequest.Age)
	suite.Equal(createdAt.Format("2006-01-02T15:04:05.000Z07:00"), paymentRequest.SubmittedAt.String()) // swagger formats to milliseconds
	suite.Equal(hhgMove.Locator, paymentRequest.Locator)

	suite.Equal(*hhgMove.Orders.DepartmentIndicator, string(paymentRequest.DepartmentIndicator))
}

func (suite *HandlerSuite) TestGetPaymentRequestsQueueHandlerUnauthorizedRole() {
	officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})

	request := httptest.NewRequest("GET", "/queues/payment-requests", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetPaymentRequestsQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetPaymentRequestsQueueHandler{
		context,
		paymentrequest.NewPaymentRequestListFetcher(suite.DB()),
	}

	response := handler.Handle(params)

	suite.Assertions.IsType(&queues.GetPaymentRequestsQueueForbidden{}, response)
}

func (suite *HandlerSuite) TestGetPaymentRequestsQueueHandlerServerError() {
	officeUser := testdatagen.MakeTIOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})

	paymentRequestListFetcher := mocks.PaymentRequestListFetcher{}

	paymentRequestListFetcher.On("FetchPaymentRequestList", officeUser.ID,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(nil, errors.New("database query error"))

	request := httptest.NewRequest("GET", "/queues/payment-requests", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetPaymentRequestsQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetPaymentRequestsQueueHandler{
		context,
		&paymentRequestListFetcher,
	}

	response := handler.Handle(params)

	suite.Assertions.IsType(&queues.GetPaymentRequestsQueueInternalServerError{}, response)
}

func (suite *HandlerSuite) TestGetPaymentRequestsQueueHandlerEmptyResults() {
	officeUser := testdatagen.MakeTIOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})

	paymentRequestListFetcher := mocks.PaymentRequestListFetcher{}

	paymentRequestListFetcher.On("FetchPaymentRequestList", officeUser.ID,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(&models.PaymentRequests{}, nil)

	request := httptest.NewRequest("GET", "/queues/payment-requests", nil)
	request = suite.AuthenticateOfficeRequest(request, officeUser)
	params := queues.GetPaymentRequestsQueueParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handler := GetPaymentRequestsQueueHandler{
		context,
		&paymentRequestListFetcher,
	}

	response := handler.Handle(params)

	suite.Assertions.IsType(&queues.GetPaymentRequestsQueueOK{}, response)
	payload := response.(*queues.GetPaymentRequestsQueueOK).Payload

	suite.Len(payload.QueuePaymentRequests, 0)
	suite.Equal(int64(0), payload.TotalCount)
}
