package https

import (
	"net/http"

	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/service"
)

type CouponHandler struct {
	srv service.CouponSrv
}

func MountCouponHandler(mux *http.ServeMux, s service.CouponSrv) {
	cp := http.NewServeMux()
	mux.Handle("/coupons/", http.StripPrefix("/coupons", cp))

	hdl := CouponHandler{
		srv: s,
	}

	cp.HandleFunc("GET /hello", hdl.RouteCheck)
}

func (h *CouponHandler) RouteCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
