package main

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/mailru/go-clickhouse"
	_ "github.com/mailru/go-clickhouse"
	"log"
	"math/rand"
	"time"
)
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
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
	if err := sqlEx(); err != nil {
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

	tx, _ := db.Begin()
	stmt, err := tx.Prepare(`
		INSERT INTO example (
		    id,
			sql_lib,
			two_letter,
			array,
			event_day,
			event_dateTime
		) VALUES (?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	for i := 0; i < 100000; i++ {
		if r, err := stmt.Exec(
			i,
			"database/sql",
			randSeq(2),
			clickhouse.Array(rand.Perm(10)),
			clickhouse.Date(time.Now()),
			time.Now(),
		); err != nil {
			_ = r
			return err
		}
	}

	if err = tx.Commit(); err != nil {
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
		    id         Int16,
		    sql_lib        String,
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