package repo

import (
	"context"

	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/dto"
)

type CouponRepo interface {
	CreateCoupon(ctx context.Context, code string, amount int) error
	ClaimCoupon(ctx context.Context, code, userId string) error
	CouponInfo(ctx context.Context, code string) (dto.CouponInfoDB, error)
}
