package server

import (
	"os"
	"strconv"
	"time"

	"github.com/soumayg9673/uber-coupon-go/packages/db"
)

type config struct {
	Addr         string        `env:"SERVER_ADDR"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT"`
	PgSQL        db.ConfigDB   `env:"PGSQL_DB_X"`
}

func getConfig() (c config) {
	// Setting port
	c.Addr = os.Getenv("APP_ENV")

	// Setting server read timeout
	if timeOut, err := time.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT")); err != nil {
		return c
	} else {
		c.ReadTimeout = timeOut
	}

	// Setting server write timeout
	if timeOut, err := time.ParseDuration(os.Getenv("SERVER_WRITE_TIMEOUT")); err != nil {
		return c
	} else {
		c.WriteTimeout = timeOut
	}

	// Setting server idle timeout
	if timeOut, err := time.ParseDuration(os.Getenv("SERVER_IDLE_TIMEOUT")); err != nil {
		return c
	} else {
		c.IdleTimeout = timeOut
	}

	// Setting Postgres DB credentials
	c.PgSQL.Name = os.Getenv("PGSQL_DB_NAME")
	c.PgSQL.Host = os.Getenv("PGSQL_DB_HOST")
	c.PgSQL.Port, _ = strconv.Atoi(os.Getenv("PGSQL_DB_PORT"))
	c.PgSQL.User = os.Getenv("PGSQL_DB_USER")
	c.PgSQL.Pwd = os.Getenv("PGSQL_DB_PWD")
	c.PgSQL.Ssl = os.Getenv("PGSQL_DB_SSL")

	return
}
