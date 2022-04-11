package main

import (
	"errors"
	"testing"
)

func TestValidateUser(t *testing.T) {
	var tests = []struct {
		email       string
		password    string
		age         int
		expectedErr error
	}{
		{
			email:       "test@example.com",
			password:    "12345",
			age:         18,
			expectedErr: nil,
		},
		{
			email:       "",
			password:    "12345",
			age:         18,
			expectedErr: errors.New("email cannot be empty"),
		},
		{
			email:       "test@example.com",
			password:    "",
			age:         18,
			expectedErr: errors.New("password cannot be empty"),
		},
		{
			email:       "test@example.com",
			password:    "12345",
			age:         16,
			expectedErr: errors.New("age must be at least 18"),
		},
	}

	for _, tdata := range tests {
		err := validateUser(tdata.email, tdata.password, tdata.age)
		errString := ""
		expectedErrString := ""
		if err != nil {
			errString = err.Error()
		}
		if tdata.expectedErr != nil {
			expectedErrString = tdata.expectedErr.Error()
		}
		if errString != expectedErrString {
			t.Errorf("got \"%s\", expect \"%s\"", errString, expectedErrString)
		}
	}
}
