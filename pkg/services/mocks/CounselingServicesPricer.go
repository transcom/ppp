// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"

	time "time"

	unit "github.com/transcom/mymove/pkg/unit"
)

// CounselingServicesPricer is an autogenerated mock type for the CounselingServicesPricer type
type CounselingServicesPricer struct {
	mock.Mock
}

// Price provides a mock function with given fields: contractCode, mtoAvailableToPrimeAt
func (_m *CounselingServicesPricer) Price(contractCode string, mtoAvailableToPrimeAt time.Time) (unit.Cents, []services.PricingParam, error) {
	ret := _m.Called(contractCode, mtoAvailableToPrimeAt)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(string, time.Time) unit.Cents); ok {
		r0 = rf(contractCode, mtoAvailableToPrimeAt)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 []services.PricingParam
	if rf, ok := ret.Get(1).(func(string, time.Time) []services.PricingParam); ok {
		r1 = rf(contractCode, mtoAvailableToPrimeAt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]services.PricingParam)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, time.Time) error); ok {
		r2 = rf(contractCode, mtoAvailableToPrimeAt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PriceUsingParams provides a mock function with given fields: params
func (_m *CounselingServicesPricer) PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, []services.PricingParam, error) {
	ret := _m.Called(params)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(models.PaymentServiceItemParams) unit.Cents); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 []services.PricingParam
	if rf, ok := ret.Get(1).(func(models.PaymentServiceItemParams) []services.PricingParam); ok {
		r1 = rf(params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]services.PricingParam)
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
