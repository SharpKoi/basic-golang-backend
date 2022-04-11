package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-backend-example/internal/database"
	"log"
	"net/http"
	"strings"
)

func (apiConf apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()

	switch r.Method {
	case http.MethodGet:
		if url == apiConf.usersHandlerPath || url == apiConf.usersHandlerPath+"/" {
			apiConf.handlerGetUsers(w, r)
		} else {
			apiConf.handlerGetUser(w, r)
		}
	case http.MethodPost:
		apiConf.handlerCreateUser(w, r)
	case http.MethodPut:
		apiConf.handlerUpdateUser(w, r)
	case http.MethodDelete:
		apiConf.handlerDeleteUser(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, errors.New("method not supported"))
	}
}

// GET /api/users or /api/users/
func (apiConf apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiConf.dbClient.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJson(w, http.StatusOK, struct {
		Users []database.User `json:"users"`
	}{
		Users: users,
	})
}

// GET /api/users/test@example.com
func (apiConf apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	email := strings.TrimPrefix(url, apiConf.usersHandlerPath)
	if email == "" || email == "/" {
		respondWithError(w, http.StatusNoContent, errors.New("email cannot be omitted"))
		return
	}
	email = strings.TrimPrefix(email, "/")

	user, err := apiConf.dbClient.GetUser(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJson(w, http.StatusOK, user)
}

// POST /api/users or /api/users/
func (apiConf apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	user, err := apiConf.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// the instruction says the user should be marshalled to json
	// is that needed?
	respondWithJson(w, http.StatusCreated, user)
	log.Println("A new user registered!")
}

// PUT /api/users/test@example.com
func (apiConf apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	email := strings.TrimPrefix(url, apiConf.usersHandlerPath)
	if email == "" || email == "/" {
		respondWithError(w, http.StatusNoContent, errors.New("email cannot be omitted"))
		return
	}
	email = strings.TrimPrefix(email, "/")

	params := struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	msg := "Successfully update user's info"
	_, err = apiConf.dbClient.UpdateUser(email, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJson(w, http.StatusOK, struct {
		Message string
	}{
		Message: msg,
	})

	log.Println(msg)
}

// DELETE /api/users/test@example.com
func (apiConf apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	email := strings.TrimPrefix(url, apiConf.usersHandlerPath)
	if email == "" || email == "/" {
		respondWithError(w, http.StatusNoContent, errors.New("email cannot be omitted"))
		return
	}
	email = strings.TrimPrefix(email, "/")

	err := apiConf.dbClient.DeleteUser(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	msg := fmt.Sprintf("Successfully deleted user: %s", email)
	respondWithJson(w, http.StatusOK, struct {
		Message string
	}{
		Message: msg,
	})
	log.Println(msg)
}
