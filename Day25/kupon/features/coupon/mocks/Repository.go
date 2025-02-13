// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	coupon "kupon/features/coupon"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddCoupon provides a mock function with given fields: userID, newCoupon
func (_m *Repository) AddCoupon(userID uint, newCoupon coupon.Coupon) (coupon.Coupon, error) {
	ret := _m.Called(userID, newCoupon)

	var r0 coupon.Coupon
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, coupon.Coupon) (coupon.Coupon, error)); ok {
		return rf(userID, newCoupon)
	}
	if rf, ok := ret.Get(0).(func(uint, coupon.Coupon) coupon.Coupon); ok {
		r0 = rf(userID, newCoupon)
	} else {
		r0 = ret.Get(0).(coupon.Coupon)
	}

	if rf, ok := ret.Get(1).(func(uint, coupon.Coupon) error); ok {
		r1 = rf(userID, newCoupon)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCoupons provides a mock function with given fields:
func (_m *Repository) GetCoupons() ([]coupon.Coupon, error) {
	ret := _m.Called()

	var r0 []coupon.Coupon
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]coupon.Coupon, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []coupon.Coupon); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coupon.Coupon)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
