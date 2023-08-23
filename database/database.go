package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func DBconnect() (*sql.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Microsecond)
	defer cancel()
	db, err := sql.Open("sqlite3", "currencies.db")
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS currencies (id INTEGER PRIMARY KEY AUTOINCREMENT, currency TEXT, value REAL, timestamp INTEGER)")
	if err != nil {
		return nil, err
	}
	return db, nil
}
