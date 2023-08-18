package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

//Final version
func main() {
	router := mux.NewRouter()

	//register route
	router.HandleFunc("/api/time", timeHandler).Methods(http.MethodGet)

	//server
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]string)

	q := r.URL.Query()

	if len(q["tz"]) == 0 {
		loc, _ := time.LoadLocation("UTC")

		result["current_time"] = fmt.Sprint(time.Now().In(loc))
	} else {
		params := strings.Split(q["tz"][0], ",")

		for _, v := range params {
			loc, err := time.LoadLocation(v)
			if err != nil {
				http.Error(w, "invalid timezone", http.StatusNotFound)
				return
			}

			result[v] = fmt.Sprint(time.Now().In(loc))
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

//Notes
//Only register 1 route!
//Other requirements have to do with query params, not a different route altogether
//
//Don't use CurrentTime struct, use a map as result instead.
//So can set json field names (struct only allows 1 fixed struct tag e.g. "current_time")
//
//r.URL.Query()
//Returns map of string : slice of string
//
//if
//No query params given --> url is /api/time
//else
//Query params given --> url is /api/time?tz=...
//
//for _, v := range params
//For each index (ignored), value in the slice
