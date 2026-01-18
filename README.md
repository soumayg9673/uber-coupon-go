# Uber Coupon Service (Go)

A concurrency-safe coupon service implemented in Go and PostgreSQL, designed to handle **flash sale traffic**, enforce **atomicity**, and prevent **duplicate claims** under high contention.

GitHub Repository:  
👉 https://github.com/soumayg9673/uber-coupon-go

---

## Prerequisites

Ensure the following are installed on your system:

- **Docker Desktop** (includes Docker Engine & Docker Compose)
- **Git**
- (Optional) `curl` for manual API testing

No local Go or PostgreSQL installation is required when running via Docker.

---

## How to Run

Clone the repository:

```bash
git clone https://github.com/soumayg9673/uber-coupon-go.git
cd uber-coupon-go
docker-compose up --build
```

---

## How to Test
Flash Sale
```bash
./test/flash_sale.sh
```

Double Dip
```bash
./test/double_dip.sh
```

---

## Architecture

### Database Design

The system enforces strict separation of concerns.

#### coupons table
Stores current coupon state:
- name (unique)
- amount (initial stock)
- remaining_amount (current availability)

#### coupon_claims table
Stores immutable claim history:
- coupon_id
- user_id
- Unique constraint on (coupon_id, user_id)

Claim history is not embedded inside the coupon record.

### Locking & Atomicity Strategy

PostgreSQL is the source of truth
- All claim operations execute inside a single database transaction
- Row-level locking is enforced using:
```
SELECT id, remaining_amount
FROM coupons
WHERE name = $1
FOR UPDATE;
```


This ensures:

- Serialized access per coupon
- No overselling during flash sales
- Safe read–modify–write semantics

#### Additional guarantees:
- A UNIQUE constraint prevents duplicate claims by the same user
- Default isolation level (READ COMMITTED) is sufficient due to explicit row locking
- No in-memory counters or application-level mutexes are used