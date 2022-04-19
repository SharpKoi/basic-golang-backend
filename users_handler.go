package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go-backend-example/internal/auth"
	"log"
	"net/http"
	"strings"
	"time"
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
		if url == apiConf.usersHandlerPath || url == apiConf.usersHandlerPath+"/" {
			apiConf.handlerCreateUser(w, r)
		} else if url == apiConf.usersHandlerPath+"/login" || url == apiConf.postsHandlerPath+"login" {
			apiConf.handlerLogin(w, r)
		}
	case http.MethodPut:
		apiConf.handlerUpdateUser(w, r)
	case http.MethodDelete:
		apiConf.handlerDeleteUser(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, errors.New("method not supported"))
	}
}

// POST /api/users/login
func (apiConf apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 401, errors.New("errors occurred while decoding request body"))
		return
	}

	// retrieve user info
	user, err := apiConf.dbClient.GetUser(params.Email)
	if err != nil {
		respondWithError(w, 401, err)
		return
	}

	// password checking
	if params.Password != user.Password {
		respondWithError(w, 401, errors.New("wrong password"))
		return
	}

	// respond with a token
	now := time.Now()
	claim := auth.UserClaims{
		Account: user.Email,
		Role:    user.Role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "go-backend-demo",
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(2 * time.Hour).Unix(),
			Subject:   "Token For User Login",
			Audience:  user.Email,
		},
	}

	// create token
	token, err := auth.CreateUserToken(claim)
	if err != nil {
		respondWithError(w, 500, err)
		return
	}

	// respond with token
	respondWithJson(w, 200, struct {
		Token string
	}{
		Token: token,
	})
}

// GET /api/users or /api/users/
func (apiConf apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.VerifyUserToken(r)
	if err != nil {
		respondWithError(w, 401, err)
		return
	}

	// TODO: use scope system to check permission
	if claims.Role != "admin" {
		respondWithError(w, 403, errors.New("have no permission to access resources"))
		return
	}

	users, err := apiConf.dbClient.GetUsers(10)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJson(w, http.StatusOK, users)
}

// GET /api/users/test@example.com
func (apiConf apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	email := strings.TrimPrefix(url, apiConf.usersHandlerPath)

	if email == "" || email == "/" {
		respondWithError(w, http.StatusNoContent, errors.New("email cannot be omitted"))
		return
	}
	// do not trim prefix with "/" first because the url may not contain "/" so that it would trim nothing.
	// trim the prefix "/" after check
	email = strings.TrimPrefix(email, "/")

	claims, err := auth.VerifyUserToken(r)
	if err != nil {
		respondWithError(w, 401, err)
		return
	}

	// TODO: use scope system to check permission
	log.Println(claims.Account)
	if claims.Role != "admin" && claims.Account != email {
		respondWithError(w, 403, errors.New("have no permission to access resources"))
		return
	}

	user, err := apiConf.dbClient.GetUser(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJson(w, http.StatusOK, user)
}

// POST /api/users or /api/users/
func (apiConf apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.VerifyUserToken(r)
	if err != nil {
		respondWithError(w, 401, err)
		return
	}

	// TODO: use scope system to check permission
	if claims.Role != "admin" {
		respondWithError(w, 403, errors.New("have no permission to access resources"))
		return
	}

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Role     string `json:"role"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	_, err = apiConf.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age, params.Role)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// the instruction says the user should be marshalled to json
	// is that needed?
	respondWithJson(w, http.StatusCreated, "Successfully created a new user!")
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

	claims, err := auth.VerifyUserToken(r)
	if err != nil {
		respondWithError(w, 401, err)
		return
	}

	// TODO: use scope system to check permission
	if claims.Role != "admin" && claims.Account != email {
		respondWithError(w, 403, errors.New("have no permission to access resources"))
		return
	}

	params := struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Role     string `json:"role"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	msg := "Successfully update user's info"
	_, err = apiConf.dbClient.UpdateUser(email, params.Password, params.Name, params.Age, params.Role)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJson(w, http.StatusOK, msg)
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

	claims, err := auth.VerifyUserToken(r)
	if err != nil {
		respondWithError(w, 401, err)
		return
	}

	// TODO: use scope system to check permission
	if claims.Role != "admin" && claims.Account != email {
		respondWithError(w, 403, errors.New("have no permission to access resources"))
		return
	}

	_, err = apiConf.dbClient.DeleteUser(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	msg := fmt.Sprintf("Successfully deleted user: %s", email)
	respondWithJson(w, http.StatusOK, msg)
	log.Println(msg)
}
