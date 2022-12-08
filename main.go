package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type CurrentTime struct {
	Current string `json:"current_time"`
}

func main() {
	router := mux.NewRouter()

	//routes
	router.HandleFunc("/api/time", timeHandler)

	//server
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	current := CurrentTime{time.Now().Format("2006-01-02 15:04:05 +0000 UTC")}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(current)
}
