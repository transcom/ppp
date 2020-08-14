// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	validate "github.com/gobuffalo/validate"
)

// MoveTaskOrderCreator is an autogenerated mock type for the MoveTaskOrderCreator type
type MoveTaskOrderCreator struct {
	mock.Mock
}

// CreateMoveTaskOrder provides a mock function with given fields: moveTaskOrder
func (_m *MoveTaskOrderCreator) CreateMoveTaskOrder(moveTaskOrder *models.Move) (*models.Move, *validate.Errors, error) {
	ret := _m.Called(moveTaskOrder)

	var r0 *models.Move
	if rf, ok := ret.Get(0).(func(*models.Move) *models.Move); ok {
		r0 = rf(moveTaskOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Move)
		}
	}

	var r1 *validate.Errors
	if rf, ok := ret.Get(1).(func(*models.Move) *validate.Errors); ok {
		r1 = rf(moveTaskOrder)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*validate.Errors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*models.Move) error); ok {
		r2 = rf(moveTaskOrder)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
