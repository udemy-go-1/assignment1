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

	//routes (order matters)
	router.HandleFunc("/api/time", timeHandler)

	//server
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	var current CurrentTime
	tz := r.URL.Query().Get("tz")

	loc, err := time.LoadLocation(tz) //if tz param not provided in the URL or tz param value is "", returns UTC
	if err != nil {                   //tz is invalid
		http.Error(w, "invalid timezone", 404)
		return
	}

	current = CurrentTime{time.Now().In(loc).String()}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(current); err != nil {
		http.Error(w, "json encoding failed", 500)
	}
}
