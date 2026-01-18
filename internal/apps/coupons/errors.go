package coupons

import "errors"

var (
	ErrCouponAlreadyExists = errors.New("token already exists.")
	ErrCouponInvalid       = errors.New("invalid token.")
	ErrCouponExpired       = errors.New("coupon expired.")
	ErrCouponClaimed       = errors.New("coupon already claimed.")
)
