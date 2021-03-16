// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"

	time "time"

	unit "github.com/transcom/mymove/pkg/unit"
)

// DomesticOriginFirstDaySITPricer is an autogenerated mock type for the DomesticOriginFirstDaySITPricer type
type DomesticOriginFirstDaySITPricer struct {
	mock.Mock
}

// Price provides a mock function with given fields: contractCode, requestedPickupDate, weight, serviceArea
func (_m *DomesticOriginFirstDaySITPricer) Price(contractCode string, requestedPickupDate time.Time, weight unit.Pound, serviceArea string) (unit.Cents, services.PricingParams, error) {
	ret := _m.Called(contractCode, requestedPickupDate, weight, serviceArea)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(string, time.Time, unit.Pound, string) unit.Cents); ok {
		r0 = rf(contractCode, requestedPickupDate, weight, serviceArea)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 services.PricingParams
	if rf, ok := ret.Get(1).(func(string, time.Time, unit.Pound, string) services.PricingParams); ok {
		r1 = rf(contractCode, requestedPickupDate, weight, serviceArea)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(services.PricingParams)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, time.Time, unit.Pound, string) error); ok {
		r2 = rf(contractCode, requestedPickupDate, weight, serviceArea)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PriceUsingParams provides a mock function with given fields: params
func (_m *DomesticOriginFirstDaySITPricer) PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, services.PricingParams, error) {
	ret := _m.Called(params)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(models.PaymentServiceItemParams) unit.Cents); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 services.PricingParams
	if rf, ok := ret.Get(1).(func(models.PaymentServiceItemParams) services.PricingParams); ok {
		r1 = rf(params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(services.PricingParams)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(models.PaymentServiceItemParams) error); ok {
		r2 = rf(params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
