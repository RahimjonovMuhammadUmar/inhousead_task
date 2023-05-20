package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var SiteList = []string{}

func main() {
	go MonitorSites()

	websitesFile, err := os.Open("websites.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer websitesFile.Close()

	scanner := bufio.NewScanner(websitesFile)
	for scanner.Scan() {
		SiteList = append(SiteList, fmt.Sprintf("https://%s", scanner.Text()))
	}

	router := http.NewServeMux()
	router.HandleFunc("/request", RequestHandler)
	router.HandleFunc("/min", MinHandler)
	router.HandleFunc("/max", MaxHandler)
	router.HandleFunc("/getData", GetDataHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
}
