package https

import (
	"encoding/json"
	"net/http"

	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons"
	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/domain/service"
)

type CouponHandler struct {
	srv service.CouponSrv
}

func MountCouponHandler(mux *http.ServeMux, s service.CouponSrv) {
	hdl := CouponHandler{
		srv: s,
	}

	mux.HandleFunc("POST /coupons", hdl.CreateCoupon)
}

func (h *CouponHandler) CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Name   string `json:"name"`
		Amount int    `json:"amount"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	if err := h.srv.CreateCoupon(r.Context(), reqBody.Name, reqBody.Amount); err != nil {
		switch err {
		case coupons.ErrDuplicateCoupon:
			w.WriteHeader(409)
			return
		default:
			w.WriteHeader(500)
		}
		return
	}

	w.WriteHeader(201)
}
