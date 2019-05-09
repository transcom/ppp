package internalapi

import (
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	"github.com/transcom/mymove/mocks"
	userop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/users"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestUnknownLoggedInUserHandler() {
	unknownUser := testdatagen.MakeDefaultUser(suite.DB())

	req := httptest.NewRequest("GET", "/users/logged_in", nil)
	req = suite.AuthenticateUserRequest(req, unknownUser)

	params := userop.ShowLoggedInUserParams{
		HTTPRequest: req,
	}

	handler := ShowLoggedInUserHandler{handlers.NewHandlerContext(suite.DB(), suite.TestLogger()), &mocks.AccessCodeFetcher{}}

	response := handler.Handle(params)

	okResponse, ok := response.(*userop.ShowLoggedInUserOK)
	suite.True(ok)
	suite.Equal(okResponse.Payload.ID.String(), unknownUser.ID.String())
}

func (suite *HandlerSuite) TestServiceMemberLoggedInUserHandler() {
	firstName := "Joseph"
	sm := testdatagen.MakeExtendedServiceMember(suite.DB(), testdatagen.Assertions{
		ServiceMember: models.ServiceMember{
			FirstName: &firstName,
		},
	})

	req := httptest.NewRequest("GET", "/users/logged_in", nil)
	req = suite.AuthenticateRequest(req, sm)

	params := userop.ShowLoggedInUserParams{
		HTTPRequest: req,
	}

	// creates access code
	code := "TEST1"
	selectedMoveType := models.SelectedMoveTypePPM
	accessCode := models.AccessCode{
		Code:     code,
		MoveType: &selectedMoveType,
	}
	suite.MustSave(&accessCode)

	accessCodeFetcher := &mocks.AccessCodeFetcher{}
	accessCodeFetcher.On("FetchAccessCode",
		mock.AnythingOfType("UUID"),
	).Return(&accessCode, nil)

	handler := ShowLoggedInUserHandler{handlers.NewHandlerContext(suite.DB(), suite.TestLogger()), accessCodeFetcher}

	response := handler.Handle(params)

	okResponse, ok := response.(*userop.ShowLoggedInUserOK)
	suite.True(ok)
	suite.Equal(okResponse.Payload.ID.String(), sm.UserID.String())
	suite.Equal(okResponse.Payload.ServiceMember.AccessCode, code)
	suite.Equal("Joseph", *okResponse.Payload.ServiceMember.FirstName)
}

func (suite *HandlerSuite) TestServiceMemberNoTransportationOfficeLoggedInUserHandler() {
	firstName := "Joseph"
	sm := testdatagen.MakeExtendedServiceMember(suite.DB(), testdatagen.Assertions{
		ServiceMember: models.ServiceMember{
			FirstName: &firstName,
		},
	})

	// Remove transportation office info from current station
	station := sm.DutyStation
	station.TransportationOfficeID = nil
	suite.MustSave(&station)

	req := httptest.NewRequest("GET", "/users/logged_in", nil)
	req = suite.AuthenticateRequest(req, sm)

	params := userop.ShowLoggedInUserParams{
		HTTPRequest: req,
	}

	handler := ShowLoggedInUserHandler{handlers.NewHandlerContext(suite.DB(), suite.TestLogger()), &mocks.AccessCodeFetcher{}}

	response := handler.Handle(params)

	okResponse, ok := response.(*userop.ShowLoggedInUserOK)
	suite.True(ok)
	suite.Equal(okResponse.Payload.ID.String(), sm.UserID.String())
}
