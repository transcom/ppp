// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"

	time "time"

	unit "github.com/transcom/mymove/pkg/unit"
)

// FuelSurchargePricer is an autogenerated mock type for the FuelSurchargePricer type
type FuelSurchargePricer struct {
	mock.Mock
}

// Price provides a mock function with given fields: contractCode, actualPickupDate, distance, weight, weightBasedDistanceMultiplier, fuelPrice
func (_m *FuelSurchargePricer) Price(contractCode string, actualPickupDate time.Time, distance unit.Miles, weight unit.Pound, weightBasedDistanceMultiplier float64, fuelPrice unit.Millicents) (unit.Cents, services.PricingParams, error) {
	ret := _m.Called(contractCode, actualPickupDate, distance, weight, weightBasedDistanceMultiplier, fuelPrice)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(string, time.Time, unit.Miles, unit.Pound, float64, unit.Millicents) unit.Cents); ok {
		r0 = rf(contractCode, actualPickupDate, distance, weight, weightBasedDistanceMultiplier, fuelPrice)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 services.PricingParams
	if rf, ok := ret.Get(1).(func(string, time.Time, unit.Miles, unit.Pound, float64, unit.Millicents) services.PricingParams); ok {
		r1 = rf(contractCode, actualPickupDate, distance, weight, weightBasedDistanceMultiplier, fuelPrice)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(services.PricingParams)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, time.Time, unit.Miles, unit.Pound, float64, unit.Millicents) error); ok {
		r2 = rf(contractCode, actualPickupDate, distance, weight, weightBasedDistanceMultiplier, fuelPrice)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PriceUsingParams provides a mock function with given fields: params
func (_m *FuelSurchargePricer) PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, services.PricingParams, error) {
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
