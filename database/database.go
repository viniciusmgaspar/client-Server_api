package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func DBconnect() (*sql.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	db, error := sql.Open("sqlite3", "currencies.db")
	if error != nil {
		return nil, error
	}
	_, error = db.ExecContext(ctx,fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT, currency TEXT, value REAL, timestamp INTEGER)", "currencies"))
	if error != nil {
		return nil, error
	}
	return db, nil
}
