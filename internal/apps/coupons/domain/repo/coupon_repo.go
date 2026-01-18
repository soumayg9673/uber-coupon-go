package repo

import "context"

type CouponRepo interface {
	CreateCoupon(ctx context.Context, code string, amount int) error
}
