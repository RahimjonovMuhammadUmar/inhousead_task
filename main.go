package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var siteList = []string{}

func main() {
	go monitorSites()

	websitesFile, err := os.Open("websites.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer websitesFile.Close()

	scanner := bufio.NewScanner(websitesFile)
	for scanner.Scan() {
		siteList = append(siteList, fmt.Sprintf("https://%s", scanner.Text()))
	}

	router := http.NewServeMux()
	router.HandleFunc("/request", requestHandler)
	router.HandleFunc("/min", minHandler)
	router.HandleFunc("/max", maxHandler)
	router.HandleFunc("/getData", getDataHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
}
