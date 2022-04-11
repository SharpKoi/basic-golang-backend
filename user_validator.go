package main

import (
	"errors"
	"fmt"
)

func validateUser(email, password string, age int) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	if password == "" {
		return errors.New("password cannot be empty")
	}
	if age < 18 {
		return fmt.Errorf("age must be at least %v", 18)
	}

	return nil
}
