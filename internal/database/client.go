package database

/* This file defines a Client struct to init/read/write the database.
 * It's highly recommended to use pgx directly instead of database/sql if your target is only PostgresQL.
 * But here I demonstrate the usage of the standard library database/sql.
 */

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Client store the database path.
type Client struct {
	sqlPath string
	db      *sql.DB
}

// NewClient exactly create a new instance of Client.
func NewClient(dbURL string, sqlPath string) Client {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	client := Client{
		sqlPath: sqlPath,
		db:      db,
	}

	// ping test
	if err := db.Ping(); err != nil {
		log.Println("Warning: Ping failed.", err)
	}

	return client
}

func (c Client) InitDB() error {
	b, err := ioutil.ReadFile(filepath.Join(c.sqlPath, "init.sql"))
	if err != nil {
		return err
	}

	_sql := string(b)
	scripts := strings.Split(_sql, ";")
	for _, script := range scripts {
		stmt, err := c.db.Prepare(strings.TrimSpace(script))
		if err != nil {
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}
