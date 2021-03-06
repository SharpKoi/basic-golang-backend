package main

import (
	"go-backend-example/internal/database"
	"os"
	"testing"
)

func TestCheckUserExists(t *testing.T) {
	tests := []struct {
		email          string
		expectedResult bool
	}{
		{
			email:          "test@example.com",
			expectedResult: true,
		},
		{
			email:          "boot.dev@example.com",
			expectedResult: true,
		},
		{
			email:          "unknown@example.com",
			expectedResult: false,
		},
	}

	c := database.NewClient(database.DSN{
		Host:     "database",
		Port:     5432,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}, "db/sql")

	for i, test := range tests {
		res, err := c.CheckUserExists(test.email)
		if err != nil {
			t.Errorf("got unexpected error: %v", err)
		}
		if res != test.expectedResult {
			t.Errorf("case %d: got \"%v\", expect \"%v\"", i+1, res, test.expectedResult)
		}
	}
}
