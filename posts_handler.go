package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-backend-example/internal/auth"
	"log"
	"net/http"
	"strings"
)

func (apiConf apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiConf.handlerRetrievePosts(w, r)
	case http.MethodPost:
		apiConf.handlerCreatePost(w, r)
	case http.MethodDelete:
		apiConf.handlerDeletePost(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, errors.New("method not supported"))
	}
}

// POST /api/posts or /api/posts/
func (apiConf apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	params := struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	claims, err := auth.VerifyUserToken(r)
	if err != nil {
		respondWithError(w, 401, err)
		return
	}

	// TODO: use scope system to check permission
	if claims.Account != params.UserEmail {
		// user can only post with his/her own email even if user is admin
		respondWithError(w, 403, errors.New("have no permission to access resources"))
		return
	}

	_, err = apiConf.dbClient.CreatePost(params.UserEmail, params.Text)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJson(w, http.StatusOK, "Successfully created a new post!")

	log.Println("A new post created.")
}

// DELETE /api/posts/example-uuid
func (apiConf apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	uuid := strings.TrimPrefix(r.URL.String(), apiConf.postsHandlerPath)
	// we need to check an omitted uuid after /posts/ ?
	if uuid == "" || uuid == "/" {
		respondWithError(w, http.StatusNoContent, errors.New("uuid cannot be omitted"))
		return
	}
	uuid = strings.TrimPrefix(uuid, "/")

	// get the target post
	post, err := apiConf.dbClient.GetPostById(uuid)
	if err != nil {
		respondWithError(w, 500, err)
		return
	}

	// verify user's token and permission
	claims, err := auth.VerifyUserToken(r)
	if err != nil {
		respondWithError(w, 401, err)
		return
	}

	// TODO: use scope system to check permission
	if claims.Role != "admin" && claims.Account != post.UserEmail {
		// only admin and post owner can delete the post
		respondWithError(w, 403, errors.New("have no permission to access resources"))
		return
	}

	// delete post
	_, err = apiConf.dbClient.DeletePost(uuid)
	if err != nil {
		respondWithError(w, http.StatusNoContent, err)
		return
	}

	msg := fmt.Sprintf("Successfully deleted post with uuid: %s", uuid)
	respondWithJson(w, http.StatusOK, msg)
	log.Println(msg)
}

// GET /api/posts/test@example.com
func (apiConf apiConfig) handlerRetrievePosts(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.String(), apiConf.postsHandlerPath)
	if email == "" || email == "/" {
		respondWithError(w, http.StatusNoContent, errors.New("email cannot be omitted"))
		return
	}
	email = strings.TrimPrefix(email, "/")

	posts, err := apiConf.dbClient.GetPosts(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJson(w, http.StatusOK, posts)
}
