// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	move_task_order "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/move_task_order"

	uuid "github.com/gofrs/uuid"
)

// MoveTaskOrderUpdater is an autogenerated mock type for the MoveTaskOrderUpdater type
type MoveTaskOrderUpdater struct {
	mock.Mock
}

// MakeAvailableToPrime provides a mock function with given fields: moveTaskOrderID, eTag, mtoApprovalServiceItems
func (_m *MoveTaskOrderUpdater) MakeAvailableToPrime(moveTaskOrderID uuid.UUID, eTag string, mtoApprovalServiceItems *map[models.ReServiceCode]bool) (*models.Move, error) {
	ret := _m.Called(moveTaskOrderID, eTag, mtoApprovalServiceItems)

	var r0 *models.Move
	if rf, ok := ret.Get(0).(func(uuid.UUID, string, *map[models.ReServiceCode]bool) *models.Move); ok {
		r0 = rf(moveTaskOrderID, eTag, mtoApprovalServiceItems)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Move)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, string, *map[models.ReServiceCode]bool) error); ok {
		r1 = rf(moveTaskOrderID, eTag, mtoApprovalServiceItems)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePostCounselingInfo provides a mock function with given fields: moveTaskOrderID, body, eTag
func (_m *MoveTaskOrderUpdater) UpdatePostCounselingInfo(moveTaskOrderID uuid.UUID, body move_task_order.UpdateMTOPostCounselingInformationBody, eTag string) (*models.Move, error) {
	ret := _m.Called(moveTaskOrderID, body, eTag)

	var r0 *models.Move
	if rf, ok := ret.Get(0).(func(uuid.UUID, move_task_order.UpdateMTOPostCounselingInformationBody, string) *models.Move); ok {
		r0 = rf(moveTaskOrderID, body, eTag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Move)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, move_task_order.UpdateMTOPostCounselingInformationBody, string) error); ok {
		r1 = rf(moveTaskOrderID, body, eTag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
