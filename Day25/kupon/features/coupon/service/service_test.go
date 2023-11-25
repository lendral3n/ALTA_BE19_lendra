package service_test

import (
	"errors"
	"kupon/features/coupon"
	"kupon/features/coupon/mocks"
	"kupon/features/coupon/service"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestAddCoupon(t *testing.T) {
	repo := mocks.NewRepository(t)
	s := service.New(repo)
	var inputData = coupon.Coupon{Title: "Diskon 50%", LinkCoupon: "https://diskon.com", CodeCoupon: "DISKON50", Images: "diskon.jpg"}
	var successReturnData = coupon.Coupon{ID: uint(1), Title: "Diskon 50%", LinkCoupon: "https://diskon.com", CodeCoupon: "DISKON50", Images: "diskon.jpg"}
	var falseData = coupon.Coupon{}
	t.Run("Success Case", func(t *testing.T) {
		repo.On("AddCoupon", uint(1), inputData).Return(successReturnData, nil).Once()
		res, err := s.AddCoupon(&jwt.Token{}, inputData)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		repo.On("AddCoupon", uint(1), falseData).Return(coupon.Coupon{}, errors.New("something happened")).Once()
		res, err := s.AddCoupon(&jwt.Token{}, falseData)
		assert.Error(t, err)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestGetCoupons(t *testing.T) {
	repo := mocks.NewRepository(t)
	s := service.New(repo)
	t.Run("Success Case", func(t *testing.T) {
		repo.On("GetCoupons").Return([]coupon.Coupon{}, nil).Once()
		res, err := s.GetCoupons(&jwt.Token{})
		assert.Nil(t, err)
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		repo.On("GetCoupons").Return(nil, errors.New("something happened")).Once()
		res, err := s.GetCoupons(&jwt.Token{})
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}
