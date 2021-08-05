package main

import (
	"database/sql"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	if err := app(); err != nil {
		log.Fatal("some err in app:", err)
	}
}

func app() error {
	return nil
}

func sqlEx(errs chan<- error) {
	db, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		errs <- err
		return
	}

	if err = db.Ping(); err != nil {
		errs <- err
		return
	}
}

func sqlxEx(errs chan<- error) {
	db, err := sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		errs <- err
		return
	}

	if err = db.Ping(); err != nil {
		errs <- err
		return
	}
}
