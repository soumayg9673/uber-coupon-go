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
	PgSQL        db.ConfigDB   `env:"POSTGRES_X"`
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
	c.PgSQL.Name = os.Getenv("POSTGRES_DB")
	c.PgSQL.Host = os.Getenv("POSTGRES_HOST")
	c.PgSQL.Port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	c.PgSQL.User = os.Getenv("POSTGRES_USER")
	c.PgSQL.Pwd = os.Getenv("POSTGRES_PASSWORD")
	c.PgSQL.Ssl = os.Getenv("POSTGRES_SSL")

	return
}
