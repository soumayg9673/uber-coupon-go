package service

import (
	"context"

	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/dto"
)

func (s *Service) CreateCoupon(ctx context.Context, code string, amount int) error {
	return s.rpo.CreateCoupon(ctx, code, amount)
}

func (s *Service) ClaimCoupon(ctx context.Context, code, userId string) error {
	return s.rpo.ClaimCoupon(ctx, code, userId)
}

func (s *Service) CouponInfo(ctx context.Context, code string) (dto.CouponInfoResp, error) {
	data, err := s.rpo.CouponInfo(ctx, code)
	if err != nil {
		return dto.CouponInfoResp{}, err
	}

	body := dto.CouponInfoResp{
		Name:      data.Name,
		Amount:    data.Amount,
		RemAmount: data.Amount - len(data.User),
		Users:     data.User,
	}

	if len(body.Users) == 0 {
		body.Users = []string{}
	}

	return body, nil
}
