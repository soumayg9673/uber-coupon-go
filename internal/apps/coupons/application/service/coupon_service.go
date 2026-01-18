package service

import "context"

func (s *Service) CreateCoupon(ctx context.Context, code string, amount int) error {
	return s.rpo.CreateCoupon(ctx, code, amount)
}
