package service

import "context"

type CouponSrv interface {
	CreateCoupon(ctx context.Context, code string, amount int) error
}
