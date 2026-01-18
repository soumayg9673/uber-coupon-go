package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PgSQL struct {
	DB   *sql.DB
	name string
	port int
}

func (p *PgSQL) NewConn(dc ConfigDB) error {
	log.Printf("connecting postgres db %s at port %v\n", dc.Name, dc.Port)
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=%s", dc.User, dc.Pwd, dc.Host, dc.Port, dc.Name, dc.Ssl)

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return err
	}

	p.DB = db
	p.name = dc.Name
	p.port = dc.Port

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Printf("failed connection with postgres db %s at port %v\n", p.name, p.port)
		return err
	}

	log.Printf("successfully connected with postgres db %s at port %v\n", p.name, p.port)
	return nil
}

func (p *PgSQL) Get() *sql.DB {
	if p.DB != nil {
		return p.DB
	}
	return nil
}

func (p *PgSQL) Close() error {
	log.Printf("closing connection to postgres db %s at port %v\n", p.name, p.port)
	return p.DB.Close()
}

func (p *PgSQL) Health() sql.DBStats {
	return p.Get().Stats()
}
