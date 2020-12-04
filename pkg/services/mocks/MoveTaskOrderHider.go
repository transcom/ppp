// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"
)

// MoveTaskOrderHider is an autogenerated mock type for the MoveTaskOrderHider type
type MoveTaskOrderHider struct {
	mock.Mock
}

// Hide provides a mock function with given fields:
func (_m *MoveTaskOrderHider) Hide() (models.Moves, error) {
	ret := _m.Called()

	var r0 models.Moves
	if rf, ok := ret.Get(0).(func() models.Moves); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Moves)
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
