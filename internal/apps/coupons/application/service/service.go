package service

import (
	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/repo"
	srv "github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/service"
)

type Service struct {
	rpo repo.CouponRepo
}

func NewCouponSrv(r repo.CouponRepo) srv.CouponSrv {
	return &Service{
		rpo: r,
	}
}
