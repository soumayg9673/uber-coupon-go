package server

import (
	"net/http"

	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/application/service"
	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/infra/persistence/postgres"
	"github.com/soumayg9673/uber-coupon-go/internal/apps/coupons/infra/transport/https"
)

func (app *app) mountApiRoutes() {
	api := http.NewServeMux()
	app.mux.Handle("/api/", http.StripPrefix("/api", api))

	// Register /coupon routes
	couponRepo := postgres.NewPersistInfra(app.pgsql.Get())
	couponSrv := service.NewCouponSrv(couponRepo)
	https.MountCouponHandler(api, couponSrv)
}
