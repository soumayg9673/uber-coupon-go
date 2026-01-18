COUPON="DOUBLE_DIP"
USER="dd_user"

# Create coupon with 5 stock
curl -X POST http://localhost:8080/api/coupons \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"$COUPON\",\"amount\":10}"

URL="http://localhost:8080/api/coupons/claim"

for i in {1..10}; do
  curl -s -o /dev/null -w "%{http_code}\n" \
    -X POST "$URL" \
    -H "Content-Type: application/json" \
    -d "{\"user_id\":\"$USER\",\"coupon_name\":\"$COUPON\"}" &
done

wait

curl http://localhost:8080/api/coupons/$COUPON