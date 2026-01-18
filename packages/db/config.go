package db

import (
	"database/sql"
	"errors"
)

type ConfigDB struct {
	Name string
	Host string
	Port int
	User string
	Pwd  string
	Ssl  string
}

type dbType = int

const (
	PSQL dbType = iota
)

var (
	ErrNoDbTypeExists = errors.New("invalid db type")
)

type Db interface {
	NewConn(ConfigDB) error
	Get() *sql.DB
	Close() error
	Health() sql.DBStats
}

func NewConn(dt dbType, d Db, dc ConfigDB) error {
	switch dt {
	case PSQL:
		d.NewConn(dc)
	}
	return ErrNoDbTypeExists
}

func Get(d Db) *sql.DB {
	return d.Get()
}

func Close(d Db) error {
	return d.Close()
}

func Health(d Db) sql.DBStats {
	return d.Health()
}
