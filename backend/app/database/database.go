package database

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Instance *sqlx.DB
	ConnStr  string
}

type Connection func() *sqlx.DB

func NewDB(connStr string) Database {
	d := Database{}
	d.ConnStr = connStr
	return d
}

func (d Database) Connect() (*sqlx.DB, error) {
	var err error
	if d.Instance == nil {
		d.Instance, err = d.getInstance()
	}
	return d.Instance, err
}

func (d Database) getInstance() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", d.ConnStr)
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Nanosecond)
	return db, err
}
