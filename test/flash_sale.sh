# Create coupon with 5 stock
curl -X POST http://localhost:8080/api/coupons \
  -H "Content-Type: application/json" \
  -d '{"name":"PROMO_SUPER_6","amount":5}'

URL="http://localhost:8080/api/coupons/claim"

for i in {1..50}; do
  curl -s -o /dev/null -w "%{http_code}\n" \
    -X POST "$URL" \
    -H "Content-Type: application/json" \
    -d "{\"user_id\":\"user_$i\",\"coupon_name\":\"PROMO_SUPER_6\"}" &
done

wait