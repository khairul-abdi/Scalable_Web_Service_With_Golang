package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/hello", HelloHandler).Methods("GET")
	r.HandleFunc("/v2/devopslearning", Devopslearning).Methods("GET")

	host, port := "0.0.0.0", "8000"
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", host, port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println(fmt.Sprintf("Server run in port %s", port))
	log.Fatal(srv.ListenAndServe())
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Hello world",
	}
	_ = json.NewEncoder(w).Encode(response)
}

func Devopslearning(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Belajar Devops",
	}

	_ = json.NewEncoder(w).Encode(response)
}
