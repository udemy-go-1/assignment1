package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"strings"
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
	if values := r.URL.Query(); len(values) == 0 || values.Get("tz") == "" {
		getCurrentTime(w)
	} else {
		getTimeZones(w, values)
	}
}

//tz param not provided in the URL or tz param value is ""
func getCurrentTime(w http.ResponseWriter) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		http.Error(w, "failed to get UTC location", 404)
		return
	}
	current := CurrentTime{time.Now().In(loc).String()}

	w.Header().Add("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(current) != nil {
		http.Error(w, "json encoding failed", 500)
	}
}

//example path: /api/time?tz=America/New_York,Asia/Kolkata
func getTimeZones(w http.ResponseWriter, v url.Values) {
	tz := strings.Split(v.Get("tz"), ",") //slice of tz param values

	current := make(map[string]string)

	for i := 0; i < len(tz); i++ {
		loc, err := time.LoadLocation(tz[i])
		if err != nil { //tz is invalid
			http.Error(w, "invalid timezone", 404)
			return
		}
		current[tz[i]] = time.Now().In(loc).String()
	}

	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(current)
	if err != nil {
		http.Error(w, "json encoding failed", 500)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "failed to write response", 500)
	}
}
