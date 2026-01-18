package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons"
	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/repo"
)

type Persist struct {
	db *sql.DB
}

func NewPersistInfra(db *sql.DB) repo.CouponRepo {
	return &Persist{
		db: db,
	}
}

func (p *Persist) CreateCoupon(ctx context.Context, code string, amount int) error {
	if _, err := p.db.ExecContext(ctx, `
		INSERT INTO coupons (
		code, 
		amount,
		remaining_amount
		) VALUES ($1, $2, $2);
		`,
		code,
		amount); err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return coupons.ErrCouponAlreadyExists
		}
		return err
	}
	return nil
}
