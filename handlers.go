package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EndpointRequests struct {
	MinCount     int `json:"min_count"`
	MaxCount     int `json:"max_count"`
	RequestCount int `json:"request_count"`
}

var endpointRequests = &EndpointRequests{}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	site := fmt.Sprintf("https://%s", r.URL.Query().Get("url"))
	responseTime := getAccessTime(site)
	endpointRequests.RequestCount++

	fmt.Fprint(w, responseTime, " seconds")
}

func MinHandler(w http.ResponseWriter, r *http.Request) {
	minSite := findMinAccessTimeSite()
	endpointRequests.MinCount++
	fmt.Fprint(w, minSite)
}

func MaxHandler(w http.ResponseWriter, r *http.Request) {
	maxSite := findMaxAccessTimeSite()
	endpointRequests.MaxCount++
	fmt.Fprint(w, maxSite)
}

func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(*endpointRequests)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
