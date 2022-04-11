package main

import (
	"errors"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	// respond with a JSON struct, or you can just pass a string
	respondWithJson(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Just a test! 😜",
	})
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotFound, errors.New("error test"))
}
