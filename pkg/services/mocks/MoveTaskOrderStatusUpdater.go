// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// MoveTaskOrderStatusUpdater is an autogenerated mock type for the MoveTaskOrderStatusUpdater type
type MoveTaskOrderStatusUpdater struct {
	mock.Mock
}

// UpdateMoveTaskOrderStatus provides a mock function with given fields: moveTaskOrderID
func (_m *MoveTaskOrderStatusUpdater) UpdateMoveTaskOrderStatus(moveTaskOrderID uuid.UUID) (*models.MoveTaskOrder, error) {
	ret := _m.Called(moveTaskOrderID)

	var r0 *models.MoveTaskOrder
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.MoveTaskOrder); ok {
		r0 = rf(moveTaskOrderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MoveTaskOrder)
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
