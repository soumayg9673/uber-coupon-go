package postgres

import (
	"context"
	"database/sql"

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
		amount
		) VALUES ($1, $2);
		`,
		code,
		amount); err != nil {
		if e, ok := err.(*pq.Error); ok {
			switch e.Code {
			case "23505":
				return coupons.ErrDuplicateCoupon
			}
		}
		return err
	}
	return nil
}
