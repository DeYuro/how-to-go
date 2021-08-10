package main

import (
	"context"
	"database/sql"
	"fmt"
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
	if err := sqlInsert(); err != nil {
		return err
	}
	if err := sqlxInsert(); err != nil {
		return err
	}

	if err := sqlSelect(); err != nil {
		return err
	}

	if err := sqlxSelect(); err != nil {
		return err
	}

	if err := drop(); err != nil {
		return err
	}
	cancel()

	return nil
}

func sqlInsert() error{
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

	return db.Close()
}

func sqlSelect() error {
	db, err := sqlx.Open("clickhouse", "http://127.0.0.1:8123/default")
	if err != nil {
		return err
	}
	defer db.Close() // ignore error because example

	if err = db.Ping(); err != nil {
		return err
	}

	rows, err := db.Query(`SELECT id, two_letter FROM example WHERE sql_lib = 'database/sql' limit 10`)
	if err != nil {
		return  err
	}
	defer rows.Close() // ignore error because example

	for rows.Next() {
		var (
			id uint32
			twoLetter string
		)

		if err = rows.Scan(&id, &twoLetter); err != nil {
			return err
		}

		fmt.Printf("Id: %d, Code: %s by database/sql rows.Next() \n ", id, twoLetter)
	}

	// Single row
	row := db.QueryRow(`SELECT id, two_letter, sql_lib FROM example WHERE id > 999`)

	if row != nil {
		var (
			id uint32
			twoLetter, sqlLib string
		)
		if err = row.Scan(&id, &twoLetter, &sqlLib); err != nil {
			return err
		}

		fmt.Printf("Id: %d, Code: %s by %s lib QueryRow \n", id, twoLetter, sqlLib)
	}

	return nil
}

func sqlxSelect() error {
	db, err := sqlx.Open("clickhouse", "http://127.0.0.1:8123/default?debug=1")

	if err != nil {
		return err
	}

	defer db.Close() // ignore error because example

	if err = db.Ping(); err != nil {
		return err
	}

	// Upper case is important! Fields must be exportable
	type row struct {
		Id uint32   `db:"id"`
		Code string `db:"two_letter"`
		Lib string  `db:"sql_lib"`
	}

	var item row
	var res []row

	err = db.Select(&res, `SELECT id, two_letter, sql_lib FROM example WHERE sql_lib = 'github.com/jmoiron/sqlx' limit 10`)
	if err != nil {
		return err
	}

	for _, row := range res {
		fmt.Printf("Id: %d, Code: %s by %s lib - *sqlx.Db.Select(...) \n", row.Id, row.Code, row.Lib)
	}

	err = db.Get(&item, `SELECT id, two_letter, sql_lib FROM example WHERE id > 100999`)
	if err != nil {
		return err
	}

	fmt.Printf("Id: %d, Code: %s by %s lib - *sqlx.Db.Get(...) \n", item.Id, item.Code, item.Lib)

	return err
}

func sqlxInsert() error {
	db, err := sqlx.Open("clickhouse", "http://127.0.0.1:8123/default")
	if err != nil {
		return err
	}
	defer db.Close() // ignore error because example

	if err = db.Ping(); err != nil {
		return err
	}

	tx := db.MustBegin()

	stmt, err := tx.Preparex(`
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
		 _ = stmt.MustExec(
			100000 + i,
			"github.com/jmoiron/sqlx",
			randSeq(2),
			clickhouse.Array(rand.Perm(10)),
			clickhouse.Date(time.Now()),
			time.Now(),
		)
	}

	return tx.Commit()
}

//migrate migrations
func migrate() error  {
	db, err := sql.Open("clickhouse", "http://127.0.0.1:8123/default")
	if err != nil {
		return err
	}

	defer db.Close() // ignore error because example

	if err = db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS example (
		    id              UInt32,
		    sql_lib         String,
			two_letter 		FixedString(2),
			array           Array(Int16),
			event_day   	Date,
			event_dateTime  DateTime
		) engine=Memory
	`)

	return err
}

func drop() error {
	db, err := sql.Open("clickhouse", "http://127.0.0.1:8123/default")
	if err != nil {
		return err
	}

	defer db.Close() // ignore error because example

	if err = db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec(`
		DROP TABLE IF EXISTS example
	`)

	return err
}