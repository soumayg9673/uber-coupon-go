package service

import "context"

func (s *Service) CreateCoupon(ctx context.Context, code string, amount int) error {
	return s.rpo.CreateCoupon(ctx, code, amount)
}

func (s *Service) ClaimCoupon(ctx context.Context, code, userId string) error {
	return s.rpo.ClaimCoupon(ctx, code, userId)
}
