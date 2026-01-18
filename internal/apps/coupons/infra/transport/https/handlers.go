package https

import (
	"encoding/json"
	"log"
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
	mux.HandleFunc("POST /coupons/claim", hdl.ClaimCoupon)
	mux.HandleFunc("GET /coupons/{name}", hdl.CouponInfo)
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

	if reqBody.Name == "" || reqBody.Amount <= 0 {
		w.WriteHeader(400)
		return
	}

	if err := h.srv.CreateCoupon(r.Context(), reqBody.Name, reqBody.Amount); err != nil {
		switch err {
		case coupons.ErrCouponAlreadyExists:
			w.WriteHeader(409)
		default:
			w.WriteHeader(500)
		}
		return
	}

	w.WriteHeader(201)
}

func (h *CouponHandler) ClaimCoupon(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		UserID string `json:"user_id"`
		Coupon string `json:"coupon_name"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	if reqBody.UserID == "" || reqBody.Coupon == "" {
		w.WriteHeader(400)
		return
	}

	if err := h.srv.ClaimCoupon(r.Context(), reqBody.Coupon, reqBody.UserID); err != nil {
		switch err {
		case coupons.ErrCouponInvalid:
			w.WriteHeader(400)
		case coupons.ErrCouponClaimed:
			w.WriteHeader(409)
		case coupons.ErrCouponExpired:
			w.WriteHeader(400)
		default:
			w.WriteHeader(500)
		}
		return
	}

	w.WriteHeader(201)
}

func (h *CouponHandler) CouponInfo(w http.ResponseWriter, r *http.Request) {
	couponName := r.PathValue("name")

	if couponName == "" {
		w.WriteHeader(400)
		return
	}

	data, err := h.srv.CouponInfo(r.Context(), couponName)
	if err != nil {
		log.Println(err)
		switch err {
		case coupons.ErrCouponInvalid:
			w.WriteHeader(400)
		default:
			w.WriteHeader(500)
		}
		return
	}

	body, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
