// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"

	uuid "github.com/gofrs/uuid"
)

// PaymentRequestListFetcher is an autogenerated mock type for the PaymentRequestListFetcher type
type PaymentRequestListFetcher struct {
	mock.Mock
}

// FetchPaymentRequestList provides a mock function with given fields: officeUserID, params
func (_m *PaymentRequestListFetcher) FetchPaymentRequestList(officeUserID uuid.UUID, params *services.FetchPaymentRequestListParams) (*models.PaymentRequests, int, error) {
	ret := _m.Called(officeUserID, params)

	var r0 *models.PaymentRequests
	if rf, ok := ret.Get(0).(func(uuid.UUID, *services.FetchPaymentRequestListParams) *models.PaymentRequests); ok {
		r0 = rf(officeUserID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PaymentRequests)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(uuid.UUID, *services.FetchPaymentRequestListParams) int); ok {
		r1 = rf(officeUserID, params)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uuid.UUID, *services.FetchPaymentRequestListParams) error); ok {
		r2 = rf(officeUserID, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
