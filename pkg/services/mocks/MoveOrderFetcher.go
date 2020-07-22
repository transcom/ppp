// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// MoveOrderFetcher is an autogenerated mock type for the MoveOrderFetcher type
type MoveOrderFetcher struct {
	mock.Mock
}

// FetchMoveOrder provides a mock function with given fields: moveTaskOrderID
func (_m *MoveOrderFetcher) FetchMoveOrder(moveTaskOrderID uuid.UUID) (*models.Order, error) {
	ret := _m.Called(moveTaskOrderID)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Order); ok {
		r0 = rf(moveTaskOrderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(moveTaskOrderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMoveOrders provides a mock function with given fields:
func (_m *MoveOrderFetcher) ListMoveOrders() ([]models.Order, error) {
	ret := _m.Called()

	var r0 []models.Order
	if rf, ok := ret.Get(0).(func() []models.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
