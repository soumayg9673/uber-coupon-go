package postgres

import (
	"database/sql"

	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/repo"
)

type Persist struct {
	db *sql.DB
}

func NewPersistInfra(db *sql.DB) repo.CouponRepo {
	return Persist{
		db: db,
	}
}
