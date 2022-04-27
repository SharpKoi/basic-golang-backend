package main

import (
	"encoding/json"
	"go-backend-example/internal/database"
	"log"
	"net/http"
	"os"
	"time"
)

type apiConfig struct {
	dbClient         database.Client
	baseUrl          string
	usersHandlerPath string
	postsHandlerPath string
}

type errorBody struct {
	Error string `json:"error"`
}

func main() {
	dbClient := database.NewClient(database.DSN{
		Host:     "database",
		Port:     5432,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}, "db/sql")
	err := dbClient.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	apiConf := apiConfig{
		dbClient:         dbClient,
		baseUrl:          "localhost:8080",
		usersHandlerPath: "/api/users",
		postsHandlerPath: "/api/posts",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)
	mux.HandleFunc("/err", testErrHandler)
	mux.HandleFunc(apiConf.usersHandlerPath, apiConf.endpointUsersHandler)
	mux.HandleFunc(apiConf.usersHandlerPath+"/", apiConf.endpointUsersHandler)
	mux.HandleFunc(apiConf.postsHandlerPath, apiConf.endpointPostsHandler)
	mux.HandleFunc(apiConf.postsHandlerPath+"/", apiConf.endpointPostsHandler)

	server := http.Server{
		Addr:         apiConf.baseUrl,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("listening on " + server.Addr + "...")
	err = server.ListenAndServe()
	log.Fatal(err)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if payload != nil {
		response, err := json.Marshal(payload)

		if err != nil {
			log.Println("error occurred while json marshalling", err)
			w.WriteHeader(http.StatusInternalServerError)
			response, _ = json.Marshal(errorBody{
				Error: err.Error(),
			})
			_, _ = w.Write(response)
			return
		}

		_, _ = w.Write(response)
		w.WriteHeader(code)
	}
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("")
	}
	respondWithJson(w, code, errorBody{
		Error: err.Error(),
	})
}
