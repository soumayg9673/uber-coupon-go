package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons"
	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/dto"
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
			remaining_amt
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

func (p *Persist) ClaimCoupon(ctx context.Context, code, userId string) (err error) {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var couponId int
	var balance int
	if err := tx.QueryRowContext(ctx, `
		SELECT id, remaining_amt
		FROM coupons
		WHERE code = $1
		FOR UPDATE
		`,
		code).Scan(
		&couponId,
		&balance,
	); err != nil {
		if err == sql.ErrNoRows {
			return coupons.ErrCouponInvalid
		}
		return err
	}

	if balance == 0 {
		return coupons.ErrCouponExpired
	}

	if _, err := tx.ExecContext(ctx, `
		INSERT INTO claims (
			coupon_id,
			user_id
		) VALUES ($1, $2);
		`,
		couponId,
		userId); err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return coupons.ErrCouponClaimed
		}
		return err
	}

	if _, err = tx.ExecContext(ctx, `
		UPDATE coupons
		SET remaining_amt = remaining_amt - 1
		WHERE id = $1;
	`, couponId); err != nil {
		return err
	}

	return tx.Commit()
}

func (p *Persist) CouponInfo(ctx context.Context, code string) (dto.CouponInfoDB, error) {
	rows, err := p.db.QueryContext(ctx, `
		SELECT coupons.code, coupons.amount, coupons.remaining_amt, claims.user_id
		FROM coupons
		LEFT JOIN claims ON claims.coupon_id = coupons.id
		WHERE coupons.code = $1;
	`, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.CouponInfoDB{}, coupons.ErrCouponInvalid
		}
		return dto.CouponInfoDB{}, err
	}

	defer rows.Close()

	data := dto.CouponInfoDB{}

	for rows.Next() {
		var user sql.NullString
		if err := rows.Scan(
			&data.Name,
			&data.Amount,
			&data.RemAmount,
			&user,
		); err != nil {
			return dto.CouponInfoDB{}, err
		}
		if user.Valid {
			data.User = append(data.User, user.String)
		}
	}

	return data, nil
}
