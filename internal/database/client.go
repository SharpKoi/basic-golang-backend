package database

/* This file defines a Client struct to init/read/write the database
 */

import (
	"encoding/json"
	"os"
)

// Client store the database path.
type Client struct {
	dbPath string
}

// NewClient exactly create a new instance of Client.
func NewClient(dbPath string) Client {
	return Client{
		dbPath: dbPath,
	}
}

func (c Client) createDB() error {
	data, _ := json.Marshal(struct{}{})
	err := os.WriteFile(c.dbPath, data, 0666)

	return err
}

func (c Client) InitDB() error {
	_, err := os.ReadFile(c.dbPath)

	if err != nil {
		// cannot read the database file or not exists
		return c.createDB()
	}

	return err
}

func (c Client) writeDB(db databaseSchema) error {
	// write data into
	data, _ := json.Marshal(db)
	err := os.WriteFile(c.dbPath, data, 0666)

	return err
}

func (c Client) readDB() (databaseSchema, error) {
	data, err := os.ReadFile(c.dbPath)
	schema := databaseSchema{
		Users: map[string]User{},
		Posts: map[string]Post{},
	}
	_ = json.Unmarshal(data, &schema)

	return schema, err
}
