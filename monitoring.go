package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var siteAvailability sync.Map

func monitorSites() {
	for {
		for _, site := range siteList {
			startTime := time.Now()
			resp, err := http.Head(site)
			if err != nil {
				log.Println("error:", err)
				continue
			}
			duration := time.Since(startTime)

			if resp.StatusCode != 200 {
				fmt.Println(site, "Unavailable, status code:", resp.StatusCode)
			} else {
				fmt.Println(site, "Available")
			}

			siteAvailability.Store(site, duration.Seconds())
		}

		time.Sleep(time.Minute * 1)
	}
}
