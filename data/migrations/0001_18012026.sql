CREATE TABLE coupons (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(90) NOT NULL UNIQUE,
    amount INT NOT NULL,
    remaining_amt INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE claims (
    id SERIAL PRIMARY KEY,
    coupon_id BIGINT NOT NULL,
    user_id VARCHAR(90) NOT NULL,
    claimed_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_coupon
        FOREIGN KEY (coupon_id) REFERENCES coupons(id),

    CONSTRAINT uq_coupon_user
        UNIQUE (coupon_id, user_id)
);