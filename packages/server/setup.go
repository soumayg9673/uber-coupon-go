package server

import (
	"log"
	"net/http"

	"github.com/soumayg9673/uber-coupon-go/packages/db"
)

type app struct {
	config config
	pgsql  db.Db
	mux    *http.ServeMux
}

func Run() error {

	appCfg := app{
		config: getConfig(),
		pgsql:  &db.PgSQL{},
		mux:    http.NewServeMux(),
	}

	// Connecting to Postgres DB
	if err := db.NewConn(db.PSQL, appCfg.pgsql, appCfg.config.PgSQL); err != nil {
		return err
	}

	httpServer := http.Server{
		Addr:         appCfg.config.Addr,
		Handler:      appCfg.mux,
		ReadTimeout:  appCfg.config.ReadTimeout,
		IdleTimeout:  appCfg.config.IdleTimeout,
		WriteTimeout: appCfg.config.WriteTimeout,
	}

	log.Printf("starting backend server at %s", appCfg.config.Addr)

	return httpServer.ListenAndServe()
}
