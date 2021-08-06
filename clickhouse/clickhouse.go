package main

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mailru/go-clickhouse"
	"log"
)

func main() {
	if err := app(); err != nil {
		log.Fatal("some err in app:", err)
	}
}

func app() error {
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)

	go func() {
		errCh <- startApp(cancel)
	}()

	select {
	case <-ctx.Done():
		log.Println("Service shutdown by ctx.Done")
		return nil
	case err := <-errCh:
		return err
	}
}

func startApp(cancel context.CancelFunc) error {
	if err := migrate(); err != nil {
		return err
	}

	return nil
}

func sqlEx() error{
	db, err := sql.Open("clickhouse", "http://127.0.0.1:8123/default")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func sqlxEx() error {
	db, err := sqlx.Open("clickhouse", "http://127.0.0.1:8123/default")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

//migrate migrations
func migrate() error  {
	db, err := sql.Open("clickhouse", "http://127.0.0.1:8123/default")
	if err != nil {
		return err
	}

	defer func() {
		_ = db.Close()
	}()
	//
	if err = db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS example (
		    id        UInt8,
			two_letter FixedString(2),
			array           Array(Int16),
			event_day   	Date,
			event_dateTime  DateTime
		) engine=Memory
	`)

	if err != nil {
		return err
	}

	return nil
}