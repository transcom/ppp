// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	pop "github.com/gobuffalo/pop/v5"

	uuid "github.com/gofrs/uuid"
)

// MTOServiceItemUpdater is an autogenerated mock type for the MTOServiceItemUpdater type
type MTOServiceItemUpdater struct {
	mock.Mock
}

// UpdateMTOServiceItem provides a mock function with given fields: db, serviceItem, eTag, validator
func (_m *MTOServiceItemUpdater) UpdateMTOServiceItem(db *pop.Connection, serviceItem *models.MTOServiceItem, eTag string, validator string) (*models.MTOServiceItem, error) {
	ret := _m.Called(db, serviceItem, eTag, validator)

	var r0 *models.MTOServiceItem
	if rf, ok := ret.Get(0).(func(*pop.Connection, *models.MTOServiceItem, string, string) *models.MTOServiceItem); ok {
		r0 = rf(db, serviceItem, eTag, validator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MTOServiceItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pop.Connection, *models.MTOServiceItem, string, string) error); ok {
		r1 = rf(db, serviceItem, eTag, validator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMTOServiceItemBase provides a mock function with given fields: db, serviceItem, eTag
func (_m *MTOServiceItemUpdater) UpdateMTOServiceItemBase(db *pop.Connection, serviceItem *models.MTOServiceItem, eTag string) (*models.MTOServiceItem, error) {
	ret := _m.Called(db, serviceItem, eTag)

	var r0 *models.MTOServiceItem
	if rf, ok := ret.Get(0).(func(*pop.Connection, *models.MTOServiceItem, string) *models.MTOServiceItem); ok {
		r0 = rf(db, serviceItem, eTag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MTOServiceItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pop.Connection, *models.MTOServiceItem, string) error); ok {
		r1 = rf(db, serviceItem, eTag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMTOServiceItemPrime provides a mock function with given fields: db, serviceItem, eTag
func (_m *MTOServiceItemUpdater) UpdateMTOServiceItemPrime(db *pop.Connection, serviceItem *models.MTOServiceItem, eTag string) (*models.MTOServiceItem, error) {
	ret := _m.Called(db, serviceItem, eTag)

	var r0 *models.MTOServiceItem
	if rf, ok := ret.Get(0).(func(*pop.Connection, *models.MTOServiceItem, string) *models.MTOServiceItem); ok {
		r0 = rf(db, serviceItem, eTag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MTOServiceItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pop.Connection, *models.MTOServiceItem, string) error); ok {
		r1 = rf(db, serviceItem, eTag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMTOServiceItemStatus provides a mock function with given fields: mtoServiceItemID, status, rejectionReason, eTag
func (_m *MTOServiceItemUpdater) UpdateMTOServiceItemStatus(mtoServiceItemID uuid.UUID, status models.MTOServiceItemStatus, rejectionReason *string, eTag string) (*models.MTOServiceItem, error) {
	ret := _m.Called(mtoServiceItemID, status, rejectionReason, eTag)

	var r0 *models.MTOServiceItem
	if rf, ok := ret.Get(0).(func(uuid.UUID, models.MTOServiceItemStatus, *string, string) *models.MTOServiceItem); ok {
		r0 = rf(mtoServiceItemID, status, rejectionReason, eTag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MTOServiceItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, models.MTOServiceItemStatus, *string, string) error); ok {
		r1 = rf(mtoServiceItemID, status, rejectionReason, eTag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
